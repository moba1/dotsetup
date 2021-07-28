package dotsetup

import (
	"path"
	"testing"
)

func TestDirectory_Command(t *testing.T) {
	path := path.Join("sample", "path")
	mode := "755"
	d := Directory{
		Path: path,
		Mode: mode,
	}
	expected := []Command{
		NewCommand(
			RawCommand{"mkdir", "-p", "-m", mode, path},
			false,
		),
	}
	test_CommandArray(t, expected, d.Commands())

	d = Directory{
		Path: path,
	}
	expected = []Command{
		NewCommand(
			RawCommand{"mkdir", "-p", path},
			false,
		),
	}
	test_CommandArray(t, expected, d.Commands())
}
