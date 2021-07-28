package dotsetup

import "testing"

func TestExecute_Command(t *testing.T) {
	rawCommand := RawCommand{
		"col", "-b", "-x",
	}
	e := Execute{
		RawCommands: []ExecuteCommand{
			{
				RawCommand: rawCommand,
				DoRoot:     false,
			},
		},
	}
	expected := []Command{
		NewCommand(
			rawCommand,
			false,
		),
	}
	test_CommandArray(t, expected, e.Commands())
}
