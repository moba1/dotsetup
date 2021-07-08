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

func test_Command(t *testing.T, expectedVal dotsetup.Command, realVal dotsetup.Command) {
	if !reflect.DeepEqual(expectedVal, realVal) {
		t.Error(
			"expected value: ", expectedVal,
			", ",
			"real value: ", realVal,
		)
	}
}

func test_CommandArray(t *testing.T, expectedVals []dotsetup.Command, realVals []dotsetup.Command) {
	for i, expectedVal := range expectedVals {
		test_Command(t, expectedVal, realVals[i])
	}
}
