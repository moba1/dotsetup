package dotsetup

import (
	"log"
)

type Package struct {
	Name string
}

func (p *Package) RawCommands() []RawCommand {
	var pm string
	var pm_opts []string
	switch Os {
	case "Fedora":
		pm = "dnf"
		pm_opts = []string{"install", "-y"}
	case "Debian GNU/Linux":
		pm = "apt"
		pm_opts = []string{"install", "-y"}
	default:
		log.Fatal("unsupported OS")
	}

	command := append([]string{"sudo", pm}, pm_opts...)
	command = append(command, p.Name)
	return []RawCommand{
		command,
	}
}
