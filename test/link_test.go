package dotsetup_test

import (
	"path"
	"testing"

	"github.com/moba1/dotsetup"
)

func TestLink_Command(t *testing.T) {
	src := path.Join("sample", "source")
	dest := path.Join("sample", "destination")
	force := false
	expected := []dotsetup.Command{
		dotsetup.NewCommand(
			dotsetup.RawCommand{
				"ln", "-s", src, dest,
			},
			false,
		),
	}
	l := dotsetup.Link{
		Source:      src,
		Destination: dest,
		Force:       force,
	}
	test_CommandArray(t, expected, l.Commands())

	l.Force = true
	expected = []dotsetup.Command{
		dotsetup.NewCommand(
			dotsetup.RawCommand{
			"ln", "-sf", src, dest,
			},
			false,
		),
	}
	test_CommandArray(t, expected, l.Commands())
}
