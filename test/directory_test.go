package dotsetup_test

import (
	"path"
	"testing"

	"github.com/moba1/dotsetup"
)

func TestDirectory_Command(t *testing.T) {
	path := path.Join("sample", "path")
	mode := "755"
	d := dotsetup.Directory{
		Path: path,
		Mode: mode,
	}
	expected := []dotsetup.RawCommand{
		{
			"mkdir", "-p", "-m", mode,
			path,
		},
	}
	test_Command(t, expected, &d)

	d = dotsetup.Directory{
		Path: path,
	}
	expected = []dotsetup.RawCommand{
		{
			"mkdir", "-p",
			path,
		},
	}
	test_Command(t, expected, &d)
}
