package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrcichra/karmen/karmen-go-client/karmen"
	"github.com/jrcichra/karmen/karmen-go-client/result"
)

func a1(params map[string]interface{}, result *result.Result) {
	log.Println("Params a1 got:")
	spew.Dump(params)
	result.Pass()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	k := &karmen.Karmen{}
	k.Start()
	k.RegisterContainer()
	k.RegisterEvent("e1")
	k.RegisterAction("a1", a1)
	m := make(map[string]interface{})
	m["hi"] = "bye"
	k.EmitEvent("e1", m)
}
