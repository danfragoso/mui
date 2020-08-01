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
	fmt.Println("\u001b[41;1m FAIL \u001b[0m " + test.name)
	fmt.Printf("\u001b[31mError:\u001b[0m " + err.Error() + "\n")
	fmt.Println(test.description)

	fmt.Print("\n")
	t.Error(err)
}

func (test *UnitTest) Pass() {
	fmt.Println("\u001b[42;1m PASS \u001b[0m " + test.name)
	fmt.Println(test.description)
	fmt.Println("")
}

func TestParser(t *testing.T) {
	fmt.Print("Testing MUI parser...\n\n")
	var parserTests []*UnitTest

	parserTests = append(parserTests, &UnitTest{
		name:            "Simple node",
		filename:        "tests/node.mui",
		description:     "Simplest node example, with content",
		testingFunction: parseNode,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Simple node space",
		filename:        "tests/node_space.mui",
		description:     "Simplest node example, with content, whitespace, tab and newline",
		testingFunction: parseNodeSpace,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Props",
		filename:        "tests/prop.mui",
		description:     "Node with content, whitespace and props",
		testingFunction: parseProp,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Multiple props",
		filename:        "tests/multiprop.mui",
		description:     "Node with content, whitespace and props",
		testingFunction: parseMultiProp,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Multiple props and spaces",
		filename:        "tests/multiprop_space.mui",
		description:     "Node with content, whitespace and props separated by space",
		testingFunction: parseMultiPropSpace,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Single child",
		filename:        "tests/children.mui",
		description:     "Node with no content, whitespace and a single child",
		testingFunction: parseChildren,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Same level children",
		filename:        "tests/same_level_children.mui",
		description:     "Node with no content, whitespace and children",
		testingFunction: parseSameLevelChildren,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Missing paranthesis",
		filename:        "tests/missing_parenthesis.mui",
		description:     "The parser should return an error",
		testingFunction: parseMissingParenthesis,
	})

	parserTests = append(parserTests, &UnitTest{
		name:            "Missing paranthesis after content",
		filename:        "tests/missing_parenthesis_content.mui",
		description:     "The parser should return an error",
		testingFunction: parseMissingParenthesisContent,
	})

	for _, test := range parserTests {
		if err := test.Run(); err != nil {
			test.Fail(err, t)
		} else {
			test.Pass()
		}
	}
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
