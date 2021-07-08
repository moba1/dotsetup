package dotsetup

type Link struct {
	Source      string
	Destination string
	Force       bool
}

func (l *Link) Commands() []Command {
	option := "-s"
	if l.Force {
		option = "-sf"
	}
	rawCommand := []string{
		"ln", option, l.Source, l.Destination,
	}
	return []Command{
		{
			rawCommand: rawCommand,
			doRoot:     false,
		},
	}
}
