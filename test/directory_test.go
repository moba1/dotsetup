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
	expected := []dotsetup.Command{
		dotsetup.NewCommand(
			dotsetup.RawCommand{"mkdir", "-p", "-m", mode, path},
			false,
		),
	}
	test_CommandArray(t, expected, d.Commands())

	d = dotsetup.Directory{
		Path: path,
	}
	expected = []dotsetup.Command{
		dotsetup.NewCommand(
			dotsetup.RawCommand{"mkdir", "-p", path},
			false,
		),
	}
	test_CommandArray(t, expected, d.Commands())
}
