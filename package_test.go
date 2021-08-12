package dotsetup

import "testing"

func TestPackage_Command(t *testing.T) {
	pack := "sample-package"
	p := Package{
		Name: pack,
	}
	var commands RawCommand
	switch Os {
	case "debian", "ubuntu":
		commands = RawCommand{
			"sudo", "-S", "apt-get", "install", "-y", pack,
		}
	case "fedora":
		commands = RawCommand{
			"sudo", "-S", "dnf", "install", "-y", pack,
		}
	case "opensuse-leap", "opensuse-tumbleweed":
		commands = RawCommand{
			"sudo", "-S", "zypper", "-n", "install", pack,
		}
	default:
		t.Error("unsupported OS")
	}
	expected := []Command{
		NewCommand(commands, true),
	}
	test_CommandArray(t, expected, p.Commands())
}
