package dotsetup

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Script struct {
	tasks []Task
	Debug bool
}

func NewScript(tasks []Task) Script {
	return Script{
		tasks: tasks,
	}
}

func (s *Script) Execute(sudoPass string) error {
	for _, task := range s.tasks {
		for _, command := range task.Commands() {
			rawCommand := command.rawCommand
			if command.doRoot {
				rawCommand = append(RawCommand{"sudo", "-S"}, rawCommand...)
			}
			cmd := exec.Command(rawCommand[0], rawCommand[1:]...)
			var errOut bytes.Buffer
			cmd.Stderr = &errOut
			if command.doRoot {
				cmd.Stdin = bytes.NewBuffer([]byte(sudoPass))
			}

			if s.Debug {
				prompt := fmt.Sprintf(
					"[%s]",
					time.Now().Format(time.RFC3339),
				)
				commandStr := fmt.Sprintf("%v", command)
				println(prompt, commandStr)
				cmd.Stdout = os.Stdout
			}
			if err := cmd.Run(); err != nil {
				return errors.New(errOut.String())
			}
		}
	}

	return nil
}

func (s *Script) Flat() []RawCommand {
	script := []RawCommand{}
	for _, task := range s.tasks {
		for _, command := range task.Commands() {
			script = append(script, command.rawCommand)
		}
	}
	return script
}
