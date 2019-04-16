package fsm_test

import (
	"github.com/Chyroc/fsm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	as := assert.New(t)

	var test []string
	f := fsm.New(
		"init",
		map[string]fsm.Event{
			"a": {
				From: "init",
				To:   "b",
			},
			"c": {
				From: "init",
				To:   "b",
			},
		},
		map[string]fsm.Callback{
			"before_all": func(e fsm.Event) error {
				test = append(test, "1")
				return nil
			},
			"before_a": func(e fsm.Event) error {
				test = append(test, "2")
				return nil
			},
			"after_a": func(e fsm.Event) error {
				test = append(test, "3")
				return nil
			},
			"after_all": func(e fsm.Event) error {
				test = append(test, "4")
				return nil
			},
		},
	)

	as.Equal("init", f.CurrentEvent())

	{
		err := f.Trigger("invalid_event")
		as.NotNil(err)
		as.Equal("invalid_event is invalid event", err.Error())
	}

	{
		as.Nil(f.Trigger("a"))
		as.Equal([]string{"1", "2", "3", "4"}, test)
		as.Equal("b", f.CurrentEvent())
	}

	{
		err := f.Trigger("c")
		as.NotNil(err)
		as.Equal("can't trigger event c from b", err.Error())
	}

	{
		err := f.Trigger("a")
		as.NotNil(err)
		as.Equal("can't trigger event a from b", err.Error())
	}
}
