package docsetup

import "strings"

type Link struct {
	Source    string
	Destition string
	Force     bool
}

func (l *Link) ToCommand() []string {
	option := "-s"
	if l.Force {
		option = "-sf"
	}
	command := []string{
		"ln", option, l.Source, l.Destition,
	}
	return []string{
		strings.Join(command, " "),
	}
}
