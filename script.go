package docsetup

import (
	"bytes"
	"errors"
	"os/exec"
)

type Script struct {
	commands []Command
}

func NewScript(commands []Command) Script {
	return Script{
		commands: commands,
	}
}

func (s *Script) Execute() error {
	for _, command := range s.flat() {
		cmd := exec.Command(command[0], command[1:]...)
		var errOut bytes.Buffer
		cmd.Stderr = &errOut
		if err := cmd.Run(); err != nil {
			return errors.New(errOut.String())
		}
	}

	return nil
}

func (s *Script) flat() []RawCommand {
	script := []RawCommand{}
	for _, command := range s.commands {
		script = append(script, command.RawCommand()...)
	}
	return script
}
