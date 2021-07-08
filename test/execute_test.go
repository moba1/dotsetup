package dotsetup_test

import (
	"testing"

	"github.com/moba1/dotsetup"
)

func TestExecute_Command(t *testing.T) {
	rawCommand := dotsetup.RawCommand{
		"col", "-b", "-x",
	}
	e := dotsetup.Execute{
		RawCommands: []dotsetup.ExecuteCommand{
			{
				RawCommand: rawCommand,
				DoRool: false,
			},
		},
	}
	expected := []dotsetup.Command{
		dotsetup.NewCommand(
			rawCommand,
			false,
		),
	}
	test_CommandArray(t, expected, e.Commands())
}
