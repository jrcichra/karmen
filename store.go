package main

import pb "github.com/jrcichra/karmen/grpc"

type Config struct {
	Events map[string]*Event
}

type Event struct {
	HostName  HostName
	EventName EventName
	Blocks    []*Block
}

type Block struct {
	Type    BlockType
	Actions []*Action
}

type Action struct {
	HostName   HostName
	ActionName ActionName
	Parameters map[ParameterName]ParameterValue
	Conditions map[ConditionName]ConditionValue
}

type Host struct {
	Online     bool
	Events     map[string]string
	Actions    map[string]string
	Dispatcher pb.Karmen_ActionDispatcherServer
	Dispatched chan struct{}
	Deallocate chan struct{}
}

type BlockType string
type ParameterName string
type ParameterValue string
type ConditionName string
type ConditionValue string
type EventName string
type ActionName string
type HostName string
type Variable string
type VariableValue string
type UUID string

type Results map[Variable]VariableValue
type State struct {
	Hosts  map[HostName]*Host
	Events map[UUID]Results
}
