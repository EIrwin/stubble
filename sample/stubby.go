
package main 

import (
    "net/http"
    "io/ioutil"
    "log"
)

func main(){
log.Println("Running stubby server on localhost:8282")

    
        
            http.HandleFunc("/api/v1/groups", func(w http.ResponseWriter, r *http.Request) {
                
                    if "GET" == r.Method {
                        file,_ := ioutil.ReadFile("responses/groups_get.json")
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.Write(file)
                    }
                
            })
        
    
        
            http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
                
                    if "POST" == r.Method {
                        file,_ := ioutil.ReadFile("responses/users_post.json")
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.Write(file)
                    }
                
                    if "GET" == r.Method {
                        file,_ := ioutil.ReadFile("responses/users_get.json")
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.Write(file)
                    }
                
                    if "PUT" == r.Method {
                        file,_ := ioutil.ReadFile("responses/users_put.json")
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.Write(file)
                    }
                
            })
        
    

log.Fatal(http.ListenAndServe("localhost:8282", nil))
}