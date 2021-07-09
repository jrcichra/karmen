package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

func debugPrintln(args ...interface{}) {
	if debug {
		log.Println(args...)
	}
}

func debugSpew(args ...interface{}) {
	if debug {
		spew.Dump(args...)
	}
}

func isPass(code int64) bool {
	pass := false
	if code == 200 {
		pass = true
	}
	return pass
}
