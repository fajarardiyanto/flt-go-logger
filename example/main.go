package main

import (
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

type Message struct {
	msg string
}

func main() {
	l := lib.NewLib().Init()
	l.SetFormat("text").SetReportCaller(true)
	l.Info(Message{
		msg: "Also working with map struct",
	})
	l.Info("Simple testing logger")
}
