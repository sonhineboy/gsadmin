package initialize

import (
	"github.com/sonhineboy/gsadmin/service/app/event"
	"github.com/sonhineboy/gsadmin/service/app/listener"
	event2 "github.com/sonhineboy/gsadmin/service/pkg/event"
)

func EventInit() event2.EventDispatcher {

	EventDispatcher := event2.NewDispatcher()
	EventDispatcher.Register(event.TestEvent{}.GetEventName(), listener.NewTestListener())
	EventDispatcher.Register(event.LoginEvent{}.GetEventName(), listener.NewTestListener())
	return EventDispatcher

}
