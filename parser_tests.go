package mui

import (
	"errors"
	"fmt"
)

func parseNode(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if err != nil {
		return err
	}

	if tree.GetName() != "frame" {
		return errors.New("❌ Error, expected name frame got: " + tree.GetName())
	}

	if tree.GetContent() != "Hello World" {
		return errors.New("❌ Error, expected Hello World got: " + tree.GetContent())
	}

	return nil
}

func parseNodeSpace(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "frame" {
		return errors.New("❌ Error, expected name frame got: " + tree.GetName())
	}

	if tree.GetContent() != "Hello World" {
		return errors.New("❌ Error, expected Hello World got: " + tree.GetContent())
	}

	return nil
}

func parseProp(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "screen" {
		return errors.New("❌ Error, expected name screen got: " + tree.GetName())
	}

	if tree.GetContent() != "Hi!" {
		return errors.New("❌ Error, expected Hi! got: " + tree.GetContent())
	}

	if tree.GetProp("theme") != "dark" {
		return errors.New("❌ Error, expected 'dark' got: " + tree.GetProp("theme"))
	}

	return nil
}

func parseMultiProp(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "test" {
		return errors.New("❌ Error, expected name test got: " + tree.GetName())
	}

	if tree.GetContent() != "test!" {
		return errors.New("❌ Error, expected test! got: " + tree.GetContent())
	}

	if tree.GetProp("keyOne") != "valueOne" {
		return errors.New("❌ Error, expected 'valueOne' got: " + tree.GetProp("keyOne"))
	}

	if tree.GetProp("keyTwo") != "valueTwo" {
		return errors.New("❌ Error, expected 'valueTwo' got: " + tree.GetProp("keyTwo"))
	}

	return nil
}

func parseMultiPropSpace(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "test" {
		return errors.New("❌ Error, expected name test got: " + tree.GetName())
	}

	if tree.GetContent() != "test!" {
		return errors.New("❌ Error, expected test! got: " + tree.GetContent())
	}

	if tree.GetProp("keyOne") != "valueOne" {
		return errors.New("❌ Error, expected 'valueOne' got: " + tree.GetProp("keyOne"))
	}

	if tree.GetProp("keyTwo") != "valueTwo" {
		return errors.New("❌ Error, expected 'valueTwo' got: " + tree.GetProp("keyTwo"))
	}

	return nil
}

func parseChildren(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "test" {
		return errors.New("❌ Error, expected name test got: " + tree.GetName())
	}

	firstChild := tree.Children[0]

	if firstChild.GetName() != "test" {
		return errors.New("❌ Error, expected name test got: " + firstChild.GetName())
	}

	if firstChild.GetContent() != "test" {
		return errors.New("❌ Error, expected test got: " + firstChild.GetContent())
	}

	return nil
}

func parseSameLevelChildren(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree == nil {
		return errors.New("🤔 Parsed node is nil for some reason")
	}

	if err != nil {
		return err
	}

	if tree.GetName() != "root" {
		return errors.New("❌ Error, expected name root got: " + tree.GetName())
	}

	firstChild := tree.Children[0]

	if firstChild.GetName() != "childOne" {
		return errors.New("❌ Error, expected name childOne got: " + firstChild.GetName())
	}

	if firstChild.GetContent() != "childContent_1" {
		return errors.New("❌ Error, expected chilContent_1 got: " + firstChild.GetContent())
	}

	secondChild := tree.Children[1]

	if secondChild.GetName() != "childTwo" {
		return errors.New("❌ Error, expected name childTwo got: " + secondChild.GetName())
	}

	if secondChild.GetContent() != "childContent_2" {
		return errors.New("❌ Error, expected chilContent_2 got: " + secondChild.GetContent())
	}

	return nil
}

func parseMissingParenthesis(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree != nil {
		return errors.New("🤔 Parsed node should be nil but is not, for some reason")
	}

	if err == nil {
		return errors.New("❌ The parser should throw an missing parenthesis error")
	}

	return nil
}

func parseMissingParenthesisContent(test *UnitTest) error {
	testString := LoadTest(test.filename)

	tree, err := Parse(testString)
	if tree != nil {
		return errors.New("🤔 Parsed node should be nil but is not, for some reason")
	}

	fmt.Println(err)
	if err == nil {
		return errors.New("❌ The parser should throw an missing parenthesis error")
	}

	return nil
}
