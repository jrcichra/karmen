package message

import (
	"time"

	"github.com/jrcichra/karmen/karmen-go-client/common"
)

//Message - handling message functions
type Message struct {
	Type          string      `json:"type"`           //Type of message being sent
	Timestamp     int64       `json:"timestamp"`      //What time this message was created
	ContainerName string      `json:"container_name"` //Container Name we want to address
	Name          string      `json:"name"`           //Name of the event/action/container based on type
	ResponseCode  int         `json:"response_code"`  //Response code (might be nil based on type)
	Params        interface{} `json:"params"`         //Params attached to the event
	ID            string      `json:"id"`             //Message IDs for the clients to keep track of their messages (passed thru)
}

// MakeRegisterContainer -
func (m *Message) MakeRegisterContainer(containerName string) {
	m.Name = containerName
	m.Timestamp = time.Now().Unix()
	m.ResponseCode = common.OK
	m.Params = nil
	m.Type = common.REGISTERCONTAINER
	m.ContainerName = containerName
	m.ID = ""
}

//MakeRegisterEvent -
func (m *Message) MakeRegisterEvent(containerName string, eventName string) {
	m.Name = eventName
	m.Timestamp = time.Now().Unix()
	m.ResponseCode = common.OK
	m.Params = nil
	m.Type = common.REGISTEREVENT
	m.ContainerName = containerName
	m.ID = ""
}

//MakeRegisterAction -
func (m *Message) MakeRegisterAction(containerName string, actionName string) {
	m.Name = actionName
	m.Timestamp = time.Now().Unix()
	m.ResponseCode = common.OK
	m.Params = nil
	m.Type = common.REGISTERACTION
	m.ContainerName = containerName
	m.ID = ""
}

//MakeEmitEvent -
func (m *Message) MakeEmitEvent(containerName string, eventName string, params interface{}) {
	m.Name = eventName
	m.Timestamp = time.Now().Unix()
	m.ResponseCode = common.OK
	m.Params = params
	m.Type = common.EMITEVENT
	m.ContainerName = containerName
	m.ID = ""
}

//MakeActionResponse -
func (m *Message) MakeActionResponse(containerName string, actionName string, rc int) {
	m.Name = containerName
	m.Timestamp = time.Now().Unix()
	m.ResponseCode = rc
	m.Params = nil
	m.Type = common.TRIGGERACTIONRESPONSE
	m.ContainerName = containerName
	m.ID = ""
}
