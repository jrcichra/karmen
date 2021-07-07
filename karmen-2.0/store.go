package main

type Config struct {
	Events map[string]*Event
}

type Event struct {
	HostName  string
	EventName string
	Blocks    []*Block
}

type Block struct {
	Type    string
	Actions []*Action
}

type Action struct {
	HostName   string
	ActionName string
	Parameters map[string]string
	Conditions map[string]string
}
