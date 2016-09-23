package endpoints

import (
	"log"
	"testing"
)

const (
	testDefinition = "GET /test test_response.json"
	testDefinitionMultipleSpaces = "GET  /test test_response.json"
	testMethod     = "GET"
	testPath       = "/test"
	testFilePath   = "test_response.json"
)

func TestParse(t *testing.T) {
	assertParseResult(testDefinition,t)
}

func TestParseWithMultipleSpaces(t *testing.T){
	assertParseResult(testDefinitionMultipleSpaces,t)
}

func assertParseResult(definition string,t *testing.T) {
	e, err := Parse(definition)
	if err != nil {
		t.Error(err.Error())
	}

	if e.Method != testMethod {
		log.Println("'invalid parsed method")
		t.Fail()
	}

	if e.Path != testPath {
		log.Println("'invalid parsed path")
		t.Fail()
	}

	if e.FilePath != testFilePath {
		log.Println("invalid parsed file path")
		t.Fail()
	}
}
