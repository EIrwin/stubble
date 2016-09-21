package test;

import (
    "testing"
    "log"

    "github.com/eirwin/stubble/config"
)

const (
    testConfigPath = "test.yml"
    testEndpoint = "GET /test test_response.json"
    testHost = "localhost"
    testPort = "8282"
)

func TestRead(t *testing.T){
    c,err := config.Read(testConfigPath)
    if err != nil {
        t.Errorf(err.Error())
    }

    if c.Host != testHost {
        log.Println("'invalid host read from config yml")
        t.Fail()
    }

    if c.Port != testPort {
        log.Println("invalid port read from config yml")
        t.Fail()
    }

    if c.Endpoints[0] != testEndpoint {
        log.Println("invalid endpoint read from config yml")
        t.Fail()
    }
}