package docsetup

type Link struct {
	Source    string
	Destition string
	Force     bool
}

func (l *Link) Command() []RawCommand {
	option := "-s"
	if l.Force {
		option = "-sf"
	}
	command := []string{
		"ln", option, l.Source, l.Destition,
	}
	return []RawCommand{
		command,
	}
}
