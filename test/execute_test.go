package dotsetup_test

import (
	"testing"

	"github.com/moba1/dotsetup"
)

func TestExecute_Command(t *testing.T) {
	command := []string{
		"col", "-b", "-x",
	}
	e := dotsetup.Execute{
		RawCommand: command,
	}
	expected := []dotsetup.RawCommand{
		command,
	}
	test_Command(t, expected, &e)
}
