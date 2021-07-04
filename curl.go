package docsetup

type Curl struct {
	Args []string
}

func (c *Curl) Command() []RawCommand {
	return []RawCommand{
		append([]string{"curl"}, c.Args...),
	}
}
