package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

func (c *Config) readConfig(filename string) yaml.Node {
	yfile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var data yaml.Node
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (c *Config) dumpConfig() error {
	for eventfullname, event := range c.Events {
		// log.Println("Event:", eventfullname)
		for _, block := range event.Blocks {
			log.Println("Event:", eventfullname)
			log.Println("\tType", block.Type)
			for _, action := range block.Actions {
				log.Println("\t\tAction:", action.ActionName)
				for k, v := range action.Parameters {
					log.Println("\t\t\tParameter:", k, "=", v)
				}
				for k, v := range action.Conditions {
					log.Println("\t\t\tCondition:", k, v)
				}
			}
		}
	}
	return nil
}

func (c *Config) LoadConfig(filename string) error {
	// read
	data := c.readConfig(filename)
	// allocate memory
	c.Events = make(map[string]*Event)
	// parse
	return c.parseConfig(data)
}

func (c *Config) parseConfig(data yaml.Node) error {
	// for some reason you have to go into the content at the top, nothing to do with events yet
	if len(data.Content) > 1 {
		log.Fatal("Invalid YAML. There should only be one top level key")
	}
	// log.Println("parseConfig()")

	for _, events := range data.Content {
		c.parseEvents(events)
	}
	return nil
}

func (c *Config) parseEvents(events *yaml.Node) {
	// 0 = string::event, 1 = map of container.event keys
	// log.Println("parseEvents()")
	if len(events.Content) == 2 && events.Content[0].Value == "events" {
		// parse individual event
		i := 0
		// 0 = Value, 1 = Map
		eventMaps := events.Content[1].Content
		for i < len(eventMaps) {
			c.parseEvent(eventMaps[i].Value, eventMaps[i+1])
			i += 2
		}

	} else {
		log.Fatal("Invalid YAML. There should only be one top level key, 'events:'")
	}
}

func (c *Config) parseEvent(fullname string, event *yaml.Node) {
	// log.Println("parseEvent() -", fullname)
	split := strings.Split(fullname, ".")
	if len(split) != 2 {
		log.Fatal("Invalid YAML. Events should be named hostname.eventname. Found '" + fullname + "'.")
	}
	hostname := split[0]
	eventname := split[1]

	// log.Println("hostname  =", hostname)
	// log.Println("eventname =", eventname)

	blocks := make([]*Block, 0)

	typeMaps := event.Content
	for i := 0; i < len(typeMaps); i++ {
		block := &Block{}
		blocks = append(blocks, block)
		c.parseType(typeMaps[i].Content[0].Value, typeMaps[i].Content[1], block)

	}

	//add the block to the event
	c.Events[fullname] = &Event{HostName: hostname, EventName: eventname, Blocks: blocks}

}

func (c *Config) parseType(typ string, m *yaml.Node, block *Block) {
	// log.Println("parseType() -", typ)
	// log.Println("type is", typ)

	block.Type = typ

	actions := make([]*Action, 0)

	for _, v := range m.Content {
		a := &Action{}
		c.parseAction(v, a)
		actions = append(actions, a)
	}

	block.Actions = actions
}

func (c *Config) parseAction(action *yaml.Node, a *Action) {
	fullname := action.Value
	if fullname == "" {
		fullname = action.Content[0].Value
	}

	// log.Println("parseAction() -", fullname)

	split := strings.Split(fullname, ".")
	if len(split) != 2 {
		log.Fatal("Invalid YAML. Actions should be named hostname.actionname. Found '" + fullname + "'.")
	}
	hostname := split[0]
	actioname := split[1]

	a.HostName = hostname
	a.ActionName = actioname

	// log.Println("hostname  =", hostname)
	// log.Println("actioname =", actioname)

	parameters := make(map[string]string)
	conditions := make(map[string]string)

	if len(action.Content) == 2 {
		i := 0
		for i < len(action.Content[1].Content[0].Content) {
			c.parseParameter(action.Content[1].Content[0].Content[i].Value, action.Content[1].Content[0].Content[i+1], parameters, conditions)
			i += 2
		}
	}

	a.Parameters = parameters
	a.Conditions = conditions

}

func (c *Config) parseParameter(name string, parameters *yaml.Node, pMap map[string]string, cMap map[string]string) {
	value := parameters.Value
	// log.Println("parseParameter() -", name+":", value)

	if name == "if" {
		c.parseCondition(name, value, cMap)
	} else {
		pMap[name] = value
	}

}

func (c *Config) parseCondition(name string, value string, cMap map[string]string) {
	// is the condition valid syntactically?
	//split it by whitespace
	tokens := strings.Fields(value)
	if len(tokens) <= 0 {
		log.Fatal("Invalid YAML. Condition:", value, " was empty")
	}
	for _, token := range tokens {
		switch token {
		case "&&", "||":
			// log.Println("Found join:", token)
		case ">", "<", "<=", ">=", "==":
			// log.Println("Found comparison:", token)
		default:
			//determine if variable or primitive
			if isInt(token) {
				// log.Println("Found int:", token)
			} else if isFloat(token) {
				// log.Println("Found float:", token)
			} else {
				// make sure the variable doesn't have special characters
				if isVariable(token) {
					// log.Println("Found variable:", token)
				} else {
					log.Fatal("Invalid YAML. ", token, " is not a valid variable")
				}
			}
		}
	}
	// if we made it here, it's at least valid
	cMap[name] = value
}

func isFloat(s string) bool {
	ret := false
	_, err := strconv.ParseFloat(s, 32)
	if err == nil {
		ret = true
	}
	return ret
}

func isInt(s string) bool {
	ret := false
	_, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		ret = true
	}
	return ret
}

func isVariable(s string) bool {
	res := true
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.' {
			res = false
			break
		}
	}
	return res
}
