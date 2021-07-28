package dotsetup

import (
	"reflect"
	"testing"
)

func TestScript_Flat(t *testing.T) {
	rawCommands := [...]RawCommand{
		{"echo", "hello"},
		{"echo", "world"},
	}
	execs := []Task{
		&Execute{
			RawCommands: []ExecuteCommand{
				{
					RawCommand: rawCommands[0],
					DoRoot:     false,
				},
			},
		},
		&Execute{
			RawCommands: []ExecuteCommand{
				{
					RawCommand: rawCommands[1],
					DoRoot:     false,
				},
			},
		},
	}
	s := NewScript(execs)
	expected := []RawCommand{}
	for _, exec := range execs {
		for _, command := range exec.Commands() {
			expected = append(expected, command.RawCommand())
		}
	}
	retVal := s.Flat()
	if !reflect.DeepEqual(expected, retVal) {
		t.Error(
			"expected: ", expected,
			", ",
			"return value: ", retVal,
		)
	}
}
