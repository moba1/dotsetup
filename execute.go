package docsetup

type Execute struct {
	Command string
}

func (e *Execute) Script() []string {
	return []string{
		e.Command,
	}
}
