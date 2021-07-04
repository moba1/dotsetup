package docsetup

type Execute struct {
	Command string
}

func (e *Execute) ToCommand() []string {
	return []string{
		e.Command,
	}
}
