package dotsetup_test

import (
	"reflect"
	"testing"

	"github.com/moba1/dotsetup"
)

func TestScript_Flat(t *testing.T) {
	rawCommands := [...]dotsetup.RawCommand{
		{"echo", "hello"},
		{"echo", "world"},
	}
	execs := []dotsetup.Task{
		&dotsetup.Execute{
			RawCommands: []dotsetup.ExecuteCommand{
				{
					RawCommand: rawCommands[0],
					DoRool: false,
				},
			},
		},
		&dotsetup.Execute{
			RawCommands: []dotsetup.ExecuteCommand{
				{
					RawCommand: rawCommands[1],
					DoRool: false,
				},
			},
		},
	}
	s := dotsetup.NewScript(execs)
	expected := []dotsetup.RawCommand{}
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
