package docsetup

type Execute struct {
	RawCommand []string
}

func (e *Execute) Command() []RawCommand {
	return []RawCommand{
		e.RawCommand,
	}
}
