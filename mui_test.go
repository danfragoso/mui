package mui

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestParseNode(t *testing.T) {
	testName := "node"
	testString := LoadTest(testName)

	tree, err := Parse(testString)
	if err != nil {
		t.Error(err)
	}

	if tree.GetName() != "frame" {
		t.Error("❌ Error, expected name frame got: ", tree.GetName())
	}

	if tree.GetContent() != "Hello World" {
		t.Error("❌ Error, expected Hello World got: ", tree.GetContent())
	}

}

func TestParseNodeSpace(t *testing.T) {
	testName := "node_space"
	testString := LoadTest(testName)

	tree, err := Parse(testString)
	if tree == nil {
		t.Error("🤔 Parsed node is nil for some reason...")
	}

	if err != nil {
		t.Error("❌ ", err)
	}

	if tree.GetName() != "frame" {
		t.Error("❌ Error, expected name frame got: ", tree.GetName())
	}

	if tree.GetContent() != "Hello World" {
		t.Error("❌ Error, expected Hello World got: ", tree.GetContent())
	}

}

func TestParseProp(t *testing.T) {
	testName := "prop"
	testString := LoadTest(testName)

	tree, err := Parse(testString)
	if tree == nil {
		t.Error("🤔 Parsed node is nil for some reason...")
	}

	if err != nil {
		t.Error("❌ ", err)
	}

	if tree.GetName() != "screen" {
		t.Error("❌ Error, expected name screen got: ", tree.GetName())
	}

	if tree.GetContent() != "Hi!" {
		t.Error("❌ Error, expected Hi! got: ", tree.GetContent())
	}

	if tree.GetProp("theme") != "dark" {
		t.Error("❌ Error, expected 'dark' got: ", tree.GetProp("theme"))
	}

}

func TestParseMultiprop(t *testing.T) {
	testName := "multiprop"
	testString := LoadTest(testName)

	tree, err := Parse(testString)
	if tree == nil {
		t.Error("🤔 Parsed node is nil for some reason...")
	}

	if err != nil {
		t.Error("❌ ", err)
	}

	if tree.GetName() != "test" {
		t.Error("❌ Error, expected name test got: ", tree.GetName())
	}

	if tree.GetContent() != "test!" {
		t.Error("❌ Error, expected test! got: ", tree.GetContent())
	}

	if tree.GetProp("keyOne") != "valueOne" {
		t.Error("❌ Error, expected 'valueOne' got: ", tree.GetProp("keyOne"))
	}

	if tree.GetProp("keyTwo") != "valueTwo" {
		t.Error("❌ Error, expected 'valueTwo' got: ", tree.GetProp("keyTwo"))
	}

}

func TestParseMultipropSpace(t *testing.T) {
	testName := "multiprop_space"
	testString := LoadTest(testName)

	tree, err := Parse(testString)
	if tree == nil {
		t.Error("🤔 Parsed node is nil for some reason...")
	}

	if err != nil {
		t.Error("❌ ", err)
	}

	if tree.GetName() != "test" {
		t.Error("❌ Error, expected name test got: ", tree.GetName())
	}

	if tree.GetContent() != "test!" {
		t.Error("❌ Error, expected test! got: ", tree.GetContent())
	}

	if tree.GetProp("keyOne") != "valueOne" {
		t.Error("❌ Error, expected 'valueOne' got: ", tree.GetProp("keyOne"))
	}

	if tree.GetProp("keyTwo") != "valueTwo" {
		t.Error("❌ Error, expected 'valueTwo' got: ", tree.GetProp("keyTwo"))
	}

}

func LoadTest(testName string) string {
	fileName := "tests/" + testName + ".mui"
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic("⚠️ Can't load test " + testName)
	}

	fileString := string(fileData)

	fmt.Println("🔥 \033[31;1;4m" + testName + ".mui\033[0m")

	lines := strings.Split(fileString, "\n")
	for idx, line := range lines {
		fmt.Println("\033[36m  " + strconv.Itoa(idx+1) + "|  " + line + "\033[0m")
	}
	fmt.Println(" ")
	return fileString
}
