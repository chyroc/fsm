package fsm

type FSM struct {
	Events    []Event
	Callbacks []Callback
}

func New(events []Event, callbacks []Callback) FSM {
	return FSM{
		Events:    events,
		Callbacks: callbacks,
	}
}
