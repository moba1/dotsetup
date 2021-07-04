package docsetup

import "strings"

type Curl struct {
	Args []string
}

func (c *Curl) ToCommand() []string {
	command := append([]string{"curl"}, c.Args...)
	return []string{
		strings.Join(command, " "),
	}
}
