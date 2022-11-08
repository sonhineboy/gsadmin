package listener

import (
	"fmt"
	event2 "ginedu2/service/app/event"
	"ginedu2/service/src"
)

type TestListener struct {
}

func NewTestListener() *TestListener {
	return &TestListener{}
}

func (t *TestListener) Process(event src.Event) {
	switch ev := event.(type) {
	case *event2.LoginEvent:
		fmt.Println(ev.Name)
	}
}
