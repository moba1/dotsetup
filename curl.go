package docsetup

type Curl struct {
	Args []string
}

func (c *Curl) RawCommands() []RawCommand {
	return []RawCommand{
		append([]string{"curl"}, c.Args...),
	}
}
