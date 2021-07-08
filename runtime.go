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
		log.Println("Unknown Block type: " + string(block.Type))
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
			log.Println("Failing block because of configuration error")
			break
		}
		if run {
			// Build a grpc action
			a := &pb.Action{ActionName: string(action.ActionName), Timestamp: time.Now().Unix(), Parameters: k.dumbDownParamMap(action.Parameters)}
			// Form that into a request
			request := &pb.ActionRequest{Action: a, RequesterName: requesterName}
			// Send the request
			log.Println("Dispatching action:", action.ActionName, "on", action.HostName)
			err := k.State.Hosts[action.HostName].Dispatcher.Send(request)
			if err != nil {
				log.Println(err)
			}
			// Wait for a response (because it's serial)
			response, err := k.State.Hosts[action.HostName].Dispatcher.Recv()
			if err != nil {
				log.Println(err)
			}
			log.Println("Action response:")
			log.Println(response)
			// Parse the return code - may be expanded later

			var passString string
			if isPass(response.Result.Code) {
				passString = "true"
			} else {
				passString = "false"
				overallResult = false
			}

			// Store the state of this action
			k.State.Events[UUID(uuid.String())] = make(map[Variable]VariableValue)
			for key, val := range response.Result.Parameters {
				k.State.Events[UUID(uuid.String())][Variable(action.HostName)+Variable(key)] = VariableValue(val)
			}
			// Store the overall result
			k.State.Events[UUID(uuid.String())][Variable(action.HostName)+".pass"] = VariableValue(passString)
		} else {
			log.Println("Skipping", block.Type, " because condition failed")
		}
	}
	return overallResult
}

func (k *karmen) runParallelBlock(block *Block, requesterName string, uuid uuid.UUID) bool {
	return true
}
