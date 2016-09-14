#stubby - Mock JSON API Generator

# What is stubby?
Stubby is a mock JSON API generator that uses a YAML specification to define mock API endpoints and responses.

# Why stubby?
Current API response mocking solutions bloat client and/or server side code. Stubby can be ran 100% from your client and server leaving it clean and free of unecessary bloat.

# Example
  
Stubby expects a simple `YAML` configuration to generate a mock JSON API. 
 
 ``` 
host: "localhost"

port: "8282"

endpoints:
  - "GET /api/v1/users responses/users_get.json"
  - "POST /api/v1/users responses/users_post.json"
  - "PUT /api/v1/users responses/users_put.json"
  - "GET /api/v1/groups responses/groups_get.json"
  - "POST /api/v1/groups responses/groups_post.json"
  ```
  
Assuming binary is available, running the following will result in a stubby server being generated. The command assumes we have a file named `sample.yaml` that contains the configuration above.

`./stubby -p=sample.yml`

This results in the following `stubby.go` file to be generated.

```
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Running stubby server on localhost:8282")

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

Finally, we can start our stubby server by simply running the command

`go run stubby.go`

