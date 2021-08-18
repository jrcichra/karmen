package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jrcichra/conditions"
	pb "github.com/jrcichra/karmen/grpc"
)

func (k *karmen) dumbDownParamMap(m map[ParameterName]ParameterValue) map[string]string {
	// This is annoying - Ultimate fix is have the strict types defined in gRPC?
	newMap := make(map[string]string)
	for key, val := range m {
		newMap[string(key)] = string(val)
	}
	return newMap
}

func dumbDownStateMap(m map[Variable]VariableValue) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, val := range m {
		switch val {
		case "true", "false":
			b, _ := strconv.ParseBool(string(val))
			newMap[string(key)] = b
		default:
			newMap[string(key)] = string(val)
		}
	}
	return newMap
}

// m2 takes precendence over m1
func combineParamMaps(m1, m2 map[ParameterName]ParameterValue) map[ParameterName]ParameterValue {
	newMap := make(map[ParameterName]ParameterValue)
	for k, v := range m1 {
		newMap[k] = v
	}
	for k, v := range m2 {
		newMap[k] = v
	}
	return newMap
}

func (k *karmen) evaluateConditions(cMap map[ConditionName]ConditionValue, uuid uuid.UUID) (bool, bool) {
	b := true
	fail := false
	for key, val := range cMap {
		if key != "if" {
			k.eventPrint(uuid, "evaluateConditions() runtime error 1. Key should be 'if:'. Failing the job")
			b = false
			fail = true
			break
		}
		p := conditions.NewParser(strings.NewReader(string(val)))
		expr, err := p.Parse()
		if err != nil {
			k.eventPrint(uuid, "evaluateConditions() runtime error 2.")
			k.eventPrint(uuid, err)
			b = false
			fail = true
			break
		}

		b, err = conditions.Evaluate(expr, dumbDownStateMap(k.State.Events[UUID(uuid.String())]))
		if err != nil {
			k.eventPrint(uuid, "evaluateConditions() runtime error 3.")
			k.eventPrint(uuid, err)
			b = false
			fail = true
			break
		}
		k.eventPrint(uuid, "evaluateConditions() got a", b, "on condition", val)
		// break on a false
		if !b {
			break
		}
	}
	return b, fail
}

func (k *karmen) runBlock(block *Block, requesterName string, uuid uuid.UUID) bool {
	var b bool
	if block.Type == "serial" {
		b = k.runSerialBlock(block, requesterName, uuid)
	} else if block.Type == "parallel" {
		b = k.runParallelBlock(block, requesterName, uuid)
	} else {
		k.eventPrint(uuid, "Unknown Block type: "+string(block.Type))
		b = false
	}
	return b
}

// Execute each action one at a time
func (k *karmen) runSerialBlock(block *Block, requesterName string, uuid uuid.UUID) bool {
	overallResult := true
	for _, action := range block.Actions {
		//Check the conditons
		run, fail := k.evaluateConditions(action.Conditions, uuid)
		if fail {
			k.eventPrint(uuid, "Failing block because of configuration error")
			break
		}
		if run {
			result := k.runAction(uuid, action, requesterName)
			if overallResult && !result {
				overallResult = false
			}
		} else {
			k.eventPrint(uuid, "Skipping action", action.ActionName, "because condition failed")
		}
	}
	return overallResult
}

func (k *karmen) runParallelBlock(block *Block, requesterName string, uuid uuid.UUID) bool {

	type Data struct {
		Action  *Action
		Result  bool
		Skipped bool
	}

	c := make(chan Data)
	for _, action := range block.Actions {
		//Run each check in parallel
		go func(a *Action, c chan Data) {
			//Check the conditons
			run, fail := k.evaluateConditions(a.Conditions, uuid)
			if fail {
				k.eventPrint(uuid, "Failing block because of configuration error")
				c <- Data{Action: a, Result: false, Skipped: false}
			}
			if run {
				res := k.runAction(uuid, a, requesterName)
				c <- Data{Action: a, Result: res, Skipped: false}
			} else {
				k.eventPrint(uuid, "Skipping", block.Type, "because a condition failed")
				c <- Data{Action: a, Result: true, Skipped: true}
			}
		}(action, c)
	}

	overallResult := true
	// Gather up the results
	for range block.Actions {
		res := <-c
		if !res.Result && !res.Skipped {
			overallResult = false
		}
	}

	return overallResult
}

func (k *karmen) runAction(uuid uuid.UUID, action *Action, requesterName string) bool {
	// Build a grpc action - building in the actions from the static yaml and runtime actions
	// runtime parameters take precedence over static parameters
	a := &pb.Action{ActionName: string(action.ActionName), Timestamp: time.Now().Unix(), Parameters: k.dumbDownParamMap(combineParamMaps(action.Parameters, k.State.EventStates[UUID(uuid.String())]))}
	// Form that into a request
	request := &pb.ActionRequest{Action: a, RequesterName: requesterName}
	// wait for us to have a dispatcher
	for {
		if _, ok := k.State.Hosts[action.HostName]; ok {
			if k.State.Hosts[action.HostName].Dispatcher == nil {
				log.Println("Waiting for dispatcher on host " + action.HostName + "...")
				time.Sleep(1 * time.Second)
			} else {
				// We have a dispatcher, so we can break out of the loop
				break
			}
		} else {
			log.Println("Waiting for host " + action.HostName + "...")
			time.Sleep(1 * time.Second)
		}
	}
	// Send the request
	k.eventPrint(uuid, "Dispatching action:", action.ActionName, "on", action.HostName)
	err := k.State.Hosts[action.HostName].Dispatcher.Send(request)
	if err != nil {
		k.eventPrint(uuid, err)
	}
	// Wait for a response (because it's serial)
	response, err := k.State.Hosts[action.HostName].Dispatcher.Recv()
	if err != nil {
		k.eventPrint(uuid, err)
	}
	k.eventPrint(uuid, "Action response:")
	k.eventPrint(uuid, response)
	// Parse the return code - may be expanded later

	var passString string
	pass := isPass(response.Result.Code)
	if pass {
		passString = "true"
	} else {
		passString = "false"
	}

	// TODO consider passing through response params to the event state

	// Store the state of this action
	k.State.Events[UUID(uuid.String())] = make(map[Variable]VariableValue)
	for key, val := range response.Result.Parameters {
		k.State.Events[UUID(uuid.String())][Variable(action.HostName)+"-"+Variable(key)] = VariableValue(val)
	}
	// Store the overall result
	k.State.Events[UUID(uuid.String())][Variable(action.HostName)+"-"+Variable(action.ActionName)+"-pass"] = VariableValue(passString)

	return pass
}

func (k *karmen) eventPrint(uuid uuid.UUID, s ...interface{}) {
	var a []interface{}
	a = append(a, uuid.String())
	s = append(a, s...)
	log.Println(s...)
}
