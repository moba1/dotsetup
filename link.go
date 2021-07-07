package dotsetup

type Link struct {
	Source    string
	Destination string
	Force     bool
}

func (l *Link) RawCommands() []RawCommand {
	option := "-s"
	if l.Force {
		option = "-sf"
	}
	command := []string{
		"ln", option, l.Source, l.Destination,
	}
	return []RawCommand{
		command,
	}
}
