package event

type TestEvent struct {
	Name string
}

func NewTestEvent(name string) *TestEvent {
	return &TestEvent{Name: name}
}

func (t TestEvent) GetEventName() string {
	return "testEvent"
}
