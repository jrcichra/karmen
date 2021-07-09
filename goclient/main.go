package main

import (
	"log"
	"strconv"
	"time"

	"github.com/jrcichra/karmen/goclient/karmen"
)

func sleep(parameters map[string]string) *karmen.Result {
	result := &karmen.Result{}
	seconds, err := strconv.Atoi(parameters["seconds"])
	if err != nil {
		log.Println(err)
		result.Code = 500
	} else {
		log.Println("Sleeping for", seconds, "seconds")
		time.Sleep(time.Duration(seconds) * time.Second)
		log.Println("Done sleeping")
		result.Code = 200
	}
	return result
}

func main() {
	k := karmen.KarmenClient{}
	k.Init("bob")
	k.AddAction(sleep, "sleep")
	k.Register("127.0.0.1", 8080)
	result, err := k.RunEvent("pleaseSleep")
	if err != nil {
		panic(err)
	}
	log.Println("Result of pleaseSleep is:", result.Result.Code)
}
