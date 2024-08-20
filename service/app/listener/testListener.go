package listener

import (
	"fmt"
	event2 "github.com/sonhineboy/gsadmin/service/app/event"
	"github.com/sonhineboy/gsadmin/service/pkg/event"
)

type TestListener struct {
}

func NewTestListener() *TestListener {
	return &TestListener{}
}

func (t *TestListener) Process(event event.Event) {
	switch ev := event.(type) {
	case *event2.LoginEvent:
		fmt.Println(ev.Name)
	}
}
