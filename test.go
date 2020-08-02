package mui

import (
	"fmt"
	"testing"
)

type UnitTest struct {
	name            string
	filename        string
	description     string
	testingFunction func(*UnitTest) error
}

func (test *UnitTest) Run() error {
	return test.testingFunction(test)
}

func (test *UnitTest) Fail(err error, t *testing.T) {
	fmt.Println("\u001b[30m\u001b[41;1m FAIL \u001b[0m " + test.name)
	fmt.Printf("\u001b[31mError:\u001b[0m " + err.Error() + "\n")
	fmt.Println("[" + test.description + "]")

	fmt.Print("\n")
	t.Error(err)
}

func (test *UnitTest) Pass() {
	fmt.Println("\u001b[30m\u001b[42;1m PASS \u001b[0m " + test.name)
	fmt.Println("[" + test.description + "]")
	fmt.Println("")
}
