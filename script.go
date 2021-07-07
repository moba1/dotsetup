package dotsetup

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type Script struct {
	commands []Command
	Debug bool
}

func NewScript(commands []Command) Script {
	return Script{
		commands: commands,
	}
}

func (s *Script) Execute() error {
	for _, command := range s.Flat() {
		if s.Debug {
			prompt := fmt.Sprintf(
				"[%s]",
				time.Now().Format(time.RFC3339),
			)
			commandStr := fmt.Sprintf("%v", command)
			println(prompt, commandStr)
		}

		cmd := exec.Command(command[0], command[1:]...)
		var errOut bytes.Buffer
		cmd.Stderr = &errOut
		if err := cmd.Run(); err != nil {
			return errors.New(errOut.String())
		}
	}

	return nil
}

func (s *Script) Flat() []RawCommand {
	script := []RawCommand{}
	for _, command := range s.commands {
		script = append(script, command.RawCommands()...)
	}
	return script
}
