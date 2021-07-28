package dotsetup

type ExecuteCommand struct {
	RawCommand RawCommand
	DoRoot     bool
}

type Execute struct {
	RawCommands []ExecuteCommand
}

func (e *Execute) Commands() []Command {
	cs := []Command{}
	for _, c := range e.RawCommands {
		cs = append(cs, Command{
			rawCommand: c.RawCommand,
			doRoot:     c.DoRoot,
		})
	}
	return cs
}
