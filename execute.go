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
		rc := RawCommand{}
		if c.DoRoot {
			rc = RawCommand{"sudo", "-S"}
		}
		cs = append(cs, Command{
			rawCommand: append(rc, c.RawCommand...),
			doRoot:     c.DoRoot,
		})
	}
	return cs
}
