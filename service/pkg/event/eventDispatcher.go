package event

import "errors"

type Event interface {
	GetEventName() string
}

type Listener interface {
	Process(event Event)
}

type DispatcherEvent struct {
	Event map[string][]Listener
}

func NewDispatcher() *DispatcherEvent {
	return &DispatcherEvent{
		Event: make(map[string][]Listener),
	}
}

func (e *DispatcherEvent) Register(eventName string, listener Listener) {
	e.Event[eventName] = append(e.Event[eventName], listener)
}

func (e *DispatcherEvent) Dispatch(eventName Event) error {
	listener, ok := e.Event[eventName.GetEventName()]
	if ok {
		for _, v := range listener {
			v.Process(eventName)
		}
		return nil
	}
	return errors.New("未知事件")
}
