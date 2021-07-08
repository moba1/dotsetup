package dotsetup

import "strings"

type RawCommand []string

func (s *RawCommand) String() string {
	return strings.Join([]string(*s), " ")
}

type Command struct {
	rawCommand RawCommand
	doRoot     bool
}

func NewCommand(rawCommand RawCommand, doRoot bool) Command {
	return Command{
		rawCommand: rawCommand,
		doRoot: doRoot,
	}
}

func (c *Command) RawCommand() RawCommand {
	return c.rawCommand
}

func (c *Command) DoRoot() bool {
	return c.doRoot
}

type Task interface {
	Commands() []Command
}
