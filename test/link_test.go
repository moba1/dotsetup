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
	expected := []dotsetup.RawCommand{
		{
			"ln", "-s",
			src,
			dest,
		},
	}
	l := dotsetup.Link{
		Source: src,
		Destition: dest,
		Force: force,
	}
	test_Command(t, expected, &l)

	l.Force = true
	expected = []dotsetup.RawCommand{
		{
			"ln", "-sf",
			src,
			dest,
		},
	}
	test_Command(t, expected, &l)
}
