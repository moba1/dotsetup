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
		expected := []dotsetup.Command{
			dotsetup.NewCommand(
				dotsetup.RawCommand{
					"sudo", "-S", "apt-get", "install", "-y", "sample-package",
				},
				true,
			),
		}
		test_CommandArray(t, expected, p.Commands())
	default:
		t.Error("unsupported OS")
	}
}
