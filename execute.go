package docsetup

type Execute struct {
	RawCommand []string
}

func (e *Execute) RawCommands() []RawCommand {
	return []RawCommand{
		e.RawCommand,
	}
}
