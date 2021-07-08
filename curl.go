package dotsetup

type Curl struct {
	Args []string
}

func (c *Curl) Commands() []Command {
	return []Command{
		{
			rawCommand: append([]string{"curl"}, c.Args...),
			doRoot:     false,
		},
	}
}
