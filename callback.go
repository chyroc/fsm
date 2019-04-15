package fsm

// before_<event_name>
// after_<event_name>
// before_all
// after_all

type Callback struct {
	Name     string
	Function func(e *Event) error
}
