package dotsetup_test

import (
	"testing"

	"github.com/moba1/dotsetup"
)

func TestPackage_Command(t *testing.T) {
	pack := "sample-package"
	p := dotsetup.Package{
		Name: pack,
	}
	switch dotsetup.Os {
	case "Debian GNU/Linux":
		expected := []dotsetup.RawCommand{
			{
				"sudo", "apt", "install", "-y", "sample-package",
			},
		}
		test_Command(t, expected, &p)
	default:
		t.Error("unsupported OS")
	}
}
