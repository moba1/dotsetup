package dotsetup

import (
	"path"
	"testing"
)

func TestLink_Command(t *testing.T) {
	src := path.Join("sample", "source")
	dest := path.Join("sample", "destination")
	force := false
	expected := []Command{
		NewCommand(
			RawCommand{
				"ln", "-s", src, dest,
			},
			false,
		),
	}
	l := Link{
		Source:      src,
		Destination: dest,
		Force:       force,
	}
	test_CommandArray(t, expected, l.Commands())

	l.Force = true
	expected = []Command{
		NewCommand(
			RawCommand{
				"ln", "-sf", src, dest,
			},
			false,
		),
	}
	test_CommandArray(t, expected, l.Commands())
}
