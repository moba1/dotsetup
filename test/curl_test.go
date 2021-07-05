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
	expected := []dotsetup.RawCommand{
		append(dotsetup.RawCommand{"curl"}, args...),
	}
	test_Command(t, expected, &c)
}
