package src

import "errors"

type Event interface {
	GetEventName() string
}

type Listener interface {
	Process(event Event)
}

type EventDispatcher struct {
	Event map[string][]Listener
}

func NewDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Event: make(map[string][]Listener),
	}
}

func (e *EventDispatcher) Register(eventName string, listener Listener) {
	e.Event[eventName] = append(e.Event[eventName], listener)
}

func (e *EventDispatcher) Dispatch(eventName Event) error {
	listener, ok := e.Event[eventName.GetEventName()]
	if ok {
		for _, v := range listener {
			v.Process(eventName)
		}
		return nil
	}
	return errors.New("未知事件")
}
