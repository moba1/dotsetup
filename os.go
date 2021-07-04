package dotsetup

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

var Os string

func init() {
	var err error
	switch runtime.GOOS {
	case "linux":
		Os, err = linux()
		if len(Os) == 0 {
			err = errors.New("unsupported linux distro")
		}
	default:
		log.Fatal("unsupported OS")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func linux() (string, error) {
	cmd := exec.Command("/bin/sh", "-c", ". /etc/os-release; echo $NAME")
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s", errOut.String())
	}
	return out.String(), nil
}
