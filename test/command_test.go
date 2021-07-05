package dotsetup_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/moba1/dotsetup"
)

func TestRawCommand_String(t *testing.T) {
	rc := dotsetup.RawCommand{
		"ls", "-l",
	}
	expected := strings.Join(rc, " ")
	retVal := rc.String()
	if retVal != expected {
		t.Error(
			"expected: ", fmt.Sprintf("'%s'", expected),
			", ",
			"return value: ", fmt.Sprintf("'%s'", retVal),
		)
	}
}

func test_Command(t *testing.T, expected []dotsetup.RawCommand, c dotsetup.Command) {
	retVal := c.RawCommands()
	if !reflect.DeepEqual(expected, retVal) {
		t.Error(
			"expected: ", expected,
			", ",
			"return value: ", retVal,
		)
	}
}
