package dotsetup_test

import (
	"reflect"
	"testing"

	"github.com/moba1/dotsetup"
)

func TestScript_Flat(t *testing.T) {
	execs := []dotsetup.Command{
		&dotsetup.Execute{
			RawCommand: []string{"echo", "hello"},
		},
		&dotsetup.Execute{
			RawCommand: []string{"echo", "world"},
		},
	}
	s := dotsetup.NewScript(execs)
	expected := []dotsetup.RawCommand{}
	for _, exec := range execs {
		expected = append(expected, exec.RawCommands()...)
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
