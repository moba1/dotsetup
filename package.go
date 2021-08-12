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
	case "opensuse-leap", "opensuse-tumbleweed":
		pm = append(sudo, "zypper")
		pm_opts = []string{"-n", "install"}
	case "fedora":
		pm = append(sudo, "dnf")
		pm_opts = []string{"install", "-y"}
	case "debian", "ubuntu":
		pm = append(sudo, "apt-get")
		pm_opts = []string{"install", "-y"}
	case "darwin":
		pm = []string{"brew"}
		pm_opts = []string{"install", "-f"}
		doRoot = false
	case "arch":
		pm = append(sudo, "pacman")
		pm_opts = []string{"--noconfirm", "-S"}
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
