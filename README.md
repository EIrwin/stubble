#stubble - Mock JSON API Generator [![GitHub release](https://img.shields.io/github/release/eirwin/stubble.svg)](https://github.com/eirwin/stubble/releases) [![Build Status](https://travis-ci.org/EIrwin/stubble.svg?branch=master)](https://travis-ci.org/eirwin/stubble)

<img src="stubble-gopher.png" height="300" width="210" />

## What is stubble?
Stubble is a mock JSON API generator that uses a YAML specification to define mock API endpoints and responses.

## Why stubble?
Current API response mocking solutions bloat client and/or server side code. Stubble can be ran 100% from your client and server leaving it clean and free of unecessary bloat.

## Example
  
Stubble expects a simple `YAML` configuration to generate a mock JSON API. 
 
 ```yaml
host: "localhost"

port: "8282"

endpoints:
  - "GET /api/v1/users responses/users_get.json"
  - "POST /api/v1/users responses/users_post.json"
  - "PUT /api/v1/users responses/users_put.json"
  - "GET /api/v1/groups responses/groups_get.json"
  - "POST /api/v1/groups responses/groups_post.json"
  ```
  
Assuming binary is available, running the following will result in a stubble server being generated. The command assumes we have a file named `sample.yaml` that contains the configuration above.

`./stubble -p=sample.yml`

This results in the following `stubble.go` file to be generated.

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Running stubble server on localhost:8282")

	http.HandleFunc("/api/v1/groups", func(w http.ResponseWriter, r *http.Request) {

		if "GET" == r.Method {
			file, _ := ioutil.ReadFile("responses/groups_get.json")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(file)
		}

		if "POST" == r.Method {
			file, _ := ioutil.ReadFile("responses/groups_post.json")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(file)
		}

	})

	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {

		if "GET" == r.Method {
			file, _ := ioutil.ReadFile("responses/users_get.json")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(file)
		}

		if "POST" == r.Method {
			file, _ := ioutil.ReadFile("responses/users_post.json")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(file)
		}

		if "PUT" == r.Method {
			file, _ := ioutil.ReadFile("responses/users_put.json")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(file)
		}

	})

	log.Fatal(http.ListenAndServe("localhost:8282", nil))
}
```

Finally, we can start our stubble server by simply running the command

`go run stubble.go`

## Samples
To run the samples in `/samples`, perform the following steps

Try it out (Currently only can be done manually)

1.Clone the repository
```
git clone git@github.com:eirwin/stubble
```
2.Build and/or install
```
go build
go install
```
3.Navigate to `sample` directory
```
cd sample
```
4.Run stubble generator against YAML configuration
```
$GOPATH/bin/stubble -p=sample.yml
```
5.Start stubble server
```
go run stubble.go
```
6.Test out stubble
```
curl localhost:8282/api/v1/users
```

If everything was done correctly, you should see the following

```
curl localhost:8282/api/v1/users
{
    "users" : [
        {
            "name" : "User A"
        },
        {
            "group": "User B"
        },
        {
            "group": "User C"
        }
    ]
}
```


## Development

Stubble uses go templates to build http handlers in Go.

## Contributing

1. Fork it ( https://github.com/eirwin/stubble/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Tasks
1. Add Makefile
2. Add additional content-type support.
3. Implement HTTP response code support.
4. Dynamic data in responses using templating.



