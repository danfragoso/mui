package mui

import (
	"io/ioutil"
	"testing"
)

func TestParseNode(t *testing.T) {
	testString := LoadTest("node")

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
	testString := LoadTest("node_space")

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
	testString := LoadTest("prop")

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

func LoadTest(testName string) string {
	fileName := "tests/" + testName + ".mui"
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic("⚠️ Can't load test " + testName)
	}

	return string(fileData)
}
