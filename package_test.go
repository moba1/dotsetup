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
	case "fedora", "centos":
		commands = RawCommand{
			"sudo", "-S", "dnf", "install", "-y", pack,
		}
	case "opensuse-leap", "opensuse-tumbleweed":
		commands = RawCommand{
			"sudo", "-S", "zypper", "-n", "install", pack,
		}
	case "arch":
		commands = RawCommand{
			"sudo", "-S", "pacman", "--noconfirm", "-S", pack,
		}
	case "gentoo":
		commands = RawCommand{
			"emrege", pack,
		}
	default:
		t.Error("unsupported OS")
	}
	expected := []Command{
		NewCommand(commands, true),
	}
	test_CommandArray(t, expected, p.Commands())
}
