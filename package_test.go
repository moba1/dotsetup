package dotsetup

import "testing"

func TestPackage_Command(t *testing.T) {
	pack := "sample-package"
	p := Package{
		Name: pack,
	}
	switch Os {
	case "Debian GNU/Linux":
		expected := []Command{
			NewCommand(
				RawCommand{
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
