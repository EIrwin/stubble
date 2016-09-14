package endpoints

import (
	"strconv"
	"strings"
)

type Endpoint struct {
	Method   string
	Path     string
	FilePath string
	Code     int
}

func Parse(definition string) (Endpoint, error) {
	var endpoint Endpoint
	parts := strings.Split(definition, " ")
	length := len(parts)

	//parse Method
	if length >= 1 {
		endpoint.Method = parts[0]
	}

	//parse Path
	if length >= 2 {
		endpoint.Path = parts[1]
	}

	//parse file Path
	if length >= 3 {
		endpoint.FilePath = parts[2]
	}

	//parse code
	if length >= 4 {
		endpoint.Code, _ = strconv.Atoi(parts[3])
	}
	return endpoint, nil
}
