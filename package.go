package dotsetup

import "log"

type Package struct {
	Name string
}

func (p *Package) RawCommands() []RawCommand {
	var pm []string
	var pm_opts []string
	switch Os {
	case "Fedora":
		pm = []string{"sudo", "dnf"}
		pm_opts = []string{"install", "-y"}
	case "Debian GNU/Linux":
		pm = []string{"sudo", "apt-get"}
		pm_opts = []string{"install", "-y"}
	case "darwin":
		pm = []string{"brew"}
		pm_opts = []string{"install", "-f"}
	default:
		log.Fatal("unsupported OS")
	}

	command := append(pm, pm_opts...)
	command = append(command, p.Name)
	return []RawCommand{
		command,
	}
}
