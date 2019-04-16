package fsm

import "fmt"

type FSM struct {
	current   string
	Events    map[string]Event
	Callbacks map[string]Callback
}

const (
	afterALL  = "after_all"
	beforeALL = "before_all"
)

func New(initEvent string, events map[string]Event, callbacks map[string]Callback) *FSM {
	return &FSM{current: initEvent, Events: events, Callbacks: callbacks}
}

func (f *FSM) Trigger(event string) error {
	e, ok := f.Events[event]
	if !ok {
		return fmt.Errorf("%s is invalid event", event)
	}

	if f.current != e.From {
		return fmt.Errorf("can't trigger event %s from %s", event, f.current)
	}

	if fun, ok := f.Callbacks[beforeALL]; ok {
		if err := fun(e); err != nil {
			return nil
		}
	}

	if fun, ok := f.Callbacks["before_"+event]; ok {
		if err := fun(e); err != nil {
			return nil
		}
	}

	if fun, ok := f.Callbacks["after_"+event]; ok {
		if err := fun(e); err != nil {
			return nil
		}
	}

	if fun, ok := f.Callbacks[afterALL]; ok {
		if err := fun(e); err != nil {
			return nil
		}
	}

	f.current = e.To

	return nil
}

func (f *FSM) CurrentEvent() string {
	return f.current
}
