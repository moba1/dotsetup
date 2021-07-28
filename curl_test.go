package dotsetup

import "testing"

func TestCurl_Command(t *testing.T) {
	args := []string{
		"-X", "POST",
		"https://example.com",
	}
	c := Curl{
		Args: args,
	}
	expectedVals := []Command{
		NewCommand(
			append(RawCommand{"curl"}, args...),
			false,
		),
	}
	test_CommandArray(t, expectedVals, c.Commands())
}
