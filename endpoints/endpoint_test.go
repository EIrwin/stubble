package endpoints

import (
	"log"
	"testing"
)

const (
	testDefinition = "GET /test test_response.json"
	testMethod     = "GET"
	testPath       = "/test"
	testFilePath   = "test_response.json"
)

func TestParse(t *testing.T) {
	e, err := Parse(testDefinition)
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
