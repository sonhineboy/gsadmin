package event

import (
	"errors"
	"sync"
)

type Event interface {
	GetEventName() string
}

type Listener interface {
	Process(event Event)
}

type Dispatcher struct {
	Event map[string][]Listener
	sync.RWMutex
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Event: make(map[string][]Listener),
	}
}

func (e *Dispatcher) Register(eventName string, listener Listener) {
	defer e.Unlock()
	e.Lock()
	e.Event[eventName] = append(e.Event[eventName], listener)
}

func (e *Dispatcher) Dispatch(eventName Event) error {
	defer e.RUnlock()
	e.RLock()
	listener, ok := e.Event[eventName.GetEventName()]
	if ok {
		var copyListener []Listener
		//避免冲突所以复制
		copy(copyListener, listener)
		for _, v := range copyListener {
			v.Process(eventName)
		}
		return nil
	}
	return errors.New("未知事件")
}
