package docsetup

import "strings"

type RawCommand []string

func (s *RawCommand) String() string {
	return strings.Join([]string(*s), " ")
}

type Command interface {
	RawCommands() []RawCommand
}
