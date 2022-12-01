package main

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/src"
)

type TestEvent struct {
	Name string
}

func (t TestEvent) GetEventName() string {
	return "test"
}

type TestListener struct {
}

func (l TestListener) Process(event src.Event) {
	_, ok := event.(TestEvent)
	fmt.Println(event, ok)
	fmt.Println(event.GetEventName())
}

func main() {

	event := TestEvent{
		Name: "ssss",
	}
	var listener = TestListener{}
	dispatcher := src.NewDispatcher()
	dispatcher.Register(event.GetEventName(), &listener)
	_ = dispatcher.Dispatch(event)
}
