package client

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/jrcichra/karmen/karmen-go-client/result"

	"github.com/jrcichra/karmen/karmen-go-client/common"

	"github.com/jrcichra/karmen/karmen-go-client/message"
)

//Client - Karmen Client struct
type Client struct {
	Host                          string
	Port                          int
	conn                          net.Conn
	Hostname                      string
	functions                     map[string]func(params map[string]interface{}, result *result.Result)
	events                        map[string]chan message.Message
	out                           chan message.Message
	registerContainerResponseChan chan message.Message
	registerEventResponseChan     chan message.Message
	registerActionResponseChan    chan message.Message
	dispatchedEventChan           chan message.Message
}

//Start - Client Constructor of sorts
func (c *Client) Start(host string, port int) {
	s, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	c.conn = s
	c.Host = host
	c.Hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	c.Port = port
	go c.handleOutput()
}

func (c *Client) handleOutput() {
	for {
		m := <-c.out
		b, err := json.Marshal(&m)
		if err != nil {
			panic(err)
		}
		b = append(b, '\n')
		c.conn.Write(b)
	}
}

func (c *Client) handleAction(actionName string, params map[string]interface{}, result *result.Result) {
	//run the action we are supposed to run
	c.functions[actionName](params, result)
	//whatever result is (because it's a pointer) should be the pass/fail of the function
	log.Println("Result for action", actionName, "is:", result.GetResult())
	var m message.Message
	if result.GetResult() {
		m.MakeActionResponse(c.Hostname, actionName, common.OK)
	} else {
		m.MakeActionResponse(c.Hostname, actionName, common.ERROR)
	}
	c.out <- m
}

func (c *Client) handleEvent(msg message.Message) {
	//Just wait for the event to be done
	resp := <-c.events[msg.ID]
	log.Println("Event", resp.Name, "finished with a return code of", resp.ResponseCode)
	//Delete this channel from the map, it's done
	delete(c.events, msg.ID)
}

func (c *Client) processParams(params interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	switch p := params.(type) {
	case map[string]map[string]interface{}:
		for k := range p {
			res[k] = p[k]["value"]
		}
	default:
		log.Println("processParams does not understand this paramset")
	}
	return res
}

func (c *Client) handleMessages() {
	for {
		// read directly from the socket, expecting each json message to be newline separated
		d := json.NewDecoder(c.conn)
		var msg message.Message
		err := d.Decode(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		//Read the type and send it to the proper function for further processing
		switch msg.Type {
		case common.REGISTERCONTAINERRESPONSE:
			c.registerContainerResponseChan <- msg
		case common.REGISTERACTIONRESPONSE:
			c.registerActionResponseChan <- msg
		case common.REGISTEREVENTRESPONSE:
			c.registerEventResponseChan <- msg
		case common.DISPATCHEDEVENT:
			c.dispatchedEventChan <- msg
			//Make a new chan for this event id in the map - our thread will listen on this channel for it to finish
			c.events[msg.ID] = make(chan message.Message)
			//Also set up a thread to tell us when this event is complete
			go c.handleEvent(msg)
		case common.TRIGGERACTION:
			var r result.Result
			log.Println("Params before processing:", msg.Params)
			p := c.processParams(msg.Params)
			log.Println("Params after processing:", p)
			//Now run the function they registered with a callback that will fetch the result and send an action response
			go c.handleAction(msg.Name, p, &r)
		default:
			log.Println("Unknown Type:", msg.Type)
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

//RegisterContainer -
func (c *Client) RegisterContainer() {
	var m message.Message
	m.MakeRegisterContainer(c.Hostname)
	c.out <- m
	resp := <-c.registerContainerResponseChan
	log.Println("response=", resp)
	if resp.ResponseCode != common.OK {
		log.Println("While registering the container, we got a bad return code:", resp.ResponseCode)
	} else {
		log.Println("Succesfully registered container")
	}
}

//RegisterEvent -
func (c *Client) RegisterEvent(eventName string) {
	var m message.Message
	m.MakeRegisterEvent(c.Hostname, eventName)
	c.out <- m
	resp := <-c.registerEventResponseChan
	log.Println("response=", resp)
	if resp.ResponseCode != common.OK {
		log.Println("While registering the event", resp.Name, ", we got a bad return code:", resp.ResponseCode)
	} else {
		log.Println("Succesfully registered event", resp.Name)
	}
}

//RegisterAction -
func (c *Client) RegisterAction(actionName string, actionFunction func(params map[string]interface{}, result *result.Result)) {
	c.functions[actionName] = actionFunction
	var m message.Message
	m.MakeRegisterAction(c.Hostname, actionName)
	c.out <- m
	resp := <-c.registerActionResponseChan
	log.Println("response=", resp)
	if resp.ResponseCode != common.OK {
		log.Println("While registering the action", resp.Name, ", we got a bad return code:", resp.ResponseCode)
	} else {
		log.Println("Succesfully registered action", resp.Name)
	}
}

//EmitEvent -
func (c *Client) EmitEvent(eventName string, params interface{}) {
	var m message.Message
	m.MakeEmitEvent(c.Hostname, eventName, params)
	c.out <- m
	resp := <-c.dispatchedEventChan
	log.Println("response=", resp)
	if resp.ResponseCode != common.OK {
		log.Println("While emitting the event", resp.Name, ", we got a bad return code:", resp.ResponseCode)
	} else {
		log.Println("Succesfully emitted event", resp.Name)
	}
}