package mui

import (
	"fmt"
	"testing"
)

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
