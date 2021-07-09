package main

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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

// Each condition is a string of <type> :
//TODO :: WIP on AST for crazy conditions
func (k *karmen) evaluateConditions(conditions map[ConditionName]ConditionValue, uuid uuid.UUID) (bool, bool) {
	b := true
	fail := false
	for key, val := range conditions {
		if key != "if" {
			log.Println("evaluateConditions() runtime error. Key should be 'if:'. Failing the job")
			b = false
			fail = true
			break
		}
		tokens := strings.Fields(string(val))
		for _, token := range tokens {
			switch token {
			case "(":
			case ")":
			case "&&":
			case "||":
			case ">":
			case "<":
			case "<=":
			case ">=":
			case "==":
			default:
				//determine if variable or primitive
				if isInt(token) {
					// log.Println("Found int:", token)
				} else if isFloat(token) {
					// log.Println("Found float:", token)
				} else {
					// make sure the variable doesn't have special characters besides periods
					if isVariable(token) {
						debugPrintln("Found variable:", token)
						// until I get the AST figured out for this...assume it is a passString
						if val, ok := k.State.Events[UUID(uuid.String())][Variable(token)]; ok {
							//this key exists - does it pass?
							debugPrintln("Found a", val)
							if val == "true" {
								b = true
								fail = false
							} else if val == "false" {
								b = false
								fail = false
							} else {
								log.Println("passString was not true or false. It was " + val + ". Something is really wrong")
							}
						} else {
							log.Println("Invalid token.", token, "is not variable but was not defined at the time it was called")
							b = false
							fail = true
							break
						}
					} else {
						log.Println("Invalid token.", token, "is not a valid variable")
						b = false
						fail = true
						break
					}
				}
			}
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
			k.eventPrint(uuid, "Skipping", block.Type, " because condition failed")
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
				k.eventPrint(uuid, "Skipping", block.Type, " because condition failed")
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
	// Build a grpc action
	a := &pb.Action{ActionName: string(action.ActionName), Timestamp: time.Now().Unix(), Parameters: k.dumbDownParamMap(action.Parameters)}
	// Form that into a request
	request := &pb.ActionRequest{Action: a, RequesterName: requesterName}
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
	var pass bool
	if isPass(response.Result.Code) {
		passString = "true"
		pass = true
	} else {
		passString = "false"
		pass = false
	}

	// Store the state of this action
	k.State.Events[UUID(uuid.String())] = make(map[Variable]VariableValue)
	for key, val := range response.Result.Parameters {
		k.State.Events[UUID(uuid.String())][Variable(action.HostName)+Variable(key)] = VariableValue(val)
	}
	// Store the overall result
	k.State.Events[UUID(uuid.String())][Variable(action.HostName)+".pass"] = VariableValue(passString)

	return pass
}

func (k *karmen) eventPrint(uuid uuid.UUID, s ...interface{}) {
	var a []interface{}
	a = append(a, uuid.String())
	s = append(a, s...)
	log.Println(s...)
}
