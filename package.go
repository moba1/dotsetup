package dotsetup

import "log"

type Package struct {
	Name string
}

func (p *Package) Commands() []Command {
	var pm []string
	var pm_opts []string
	doRoot := true
	sudo := []string{"sudo", "-S"}
	switch Os {
	case "Fedora", "CentOS Linux":
		pm = append(sudo, "dnf")
		pm_opts = []string{"install", "-y"}
	case "Debian GNU/Linux", "Ubuntu":
		pm = append(sudo, "apt-get")
		pm_opts = []string{"install", "-y"}
	case "darwin":
		pm = []string{"brew"}
		pm_opts = []string{"install", "-f"}
		doRoot = false
	default:
		log.Fatal("unsupported OS")
	}

	rawCommand := append(pm, pm_opts...)
	rawCommand = append(rawCommand, p.Name)
	return []Command{
		{
			rawCommand: rawCommand,
			doRoot:     doRoot,
		},
	}
}
