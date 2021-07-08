package dotsetup_test

import (
	"testing"

	"github.com/moba1/dotsetup"
)

func TestCurl_Command(t *testing.T) {
	args := []string{
		"-X", "POST",
		"https://example.com",
	}
	c := dotsetup.Curl{
		Args: args,
	}
	expectedVals := []dotsetup.Command{
		dotsetup.NewCommand(
			append(dotsetup.RawCommand{"curl"}, args...),
			false,
		),
	}
	test_CommandArray(t, expectedVals, c.Commands())
}
