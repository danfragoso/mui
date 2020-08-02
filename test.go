package mui

import (
	"fmt"
	"io/ioutil"
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

func LoadTest(fileName string) string {
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic("⚠️ Can't read file " + fileName)
	}

	fileString := string(fileData)

	// lines := strings.Split(fileString, "\n")
	// for idx, line := range lines {
	// 	fmt.Println("\033[36m  " + strconv.Itoa(idx+1) + "|  " + line + "\033[0m")
	// }
	// fmt.Println(" ")
	return fileString
}
