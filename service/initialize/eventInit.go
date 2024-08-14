package initialize

import (
	"github.com/sonhineboy/gsadmin/service/app/event"
	"github.com/sonhineboy/gsadmin/service/app/listener"
	"github.com/sonhineboy/gsadmin/service/src"
)

func EventInit() src.EventDispatcher {

	EventDispatcher := src.NewDispatcher()
	EventDispatcher.Register(event.TestEvent{}.GetEventName(), listener.NewTestListener())
	EventDispatcher.Register(event.LoginEvent{}.GetEventName(), listener.NewTestListener())
	return EventDispatcher

}
