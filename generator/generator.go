package generator

import (
	"bytes"
	"html/template"
	"os"

	"github.com/eirwin/stubby/config"
	"github.com/eirwin/stubby/endpoints"
)

const (
	templatePath = "generator.tmpl"
	templateName = "Stubby"
)

type Generator struct {
	Config config.Config
}

type generatorContext struct {
	Host      string
	Port      string
	Endpoints []endpoints.Endpoint
}

func (g Generator) Run() error {
	endpoints, err := parseEndpoints(g.Config)
	if err != nil {
		panic(err)
	}

	context := generatorContext{
		Host:      g.Config.Host,
		Port:      g.Config.Port,
		Endpoints: endpoints,
	}

	t := template.New(templateName)
	t, err = t.Parse(generator())
	if err != nil {
		panic(err)
	}
	if err == nil {
		buff := bytes.NewBufferString("")
		t.Execute(buff, context)
		writeFile(buff)
	}

	return nil
}

func parseEndpoints(c config.Config) ([]endpoints.Endpoint, error) {
	var parsed []endpoints.Endpoint
	for _, endpoint := range c.Endpoints {
		e, err := endpoints.Parse(endpoint)
		if err != nil {
			return parsed, err
		}
		parsed = append(parsed, e)
	}
	return parsed, nil
}

func writeFile(buffer *bytes.Buffer) error {
	f, err := os.Create("stubby.go")
	if err != nil {
		return err
	}
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func New(c config.Config) Generator {
	return Generator{
		Config: c,
	}
}

func generator() string {
	return `package main 

            import (
                "net/http"
                "io/ioutil"
                "encoding/json"
                "log"
            )

            func main(){
                
                {{with .Endpoints}}
                    {{range .}}
                        http.HandleFunc("{{.Path}}", func(w http.ResponseWriter, r *http.Request) {
                            file,_ := ioutil.ReadFile("{{.FilePath}}")
                            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                            if err := json.NewEncoder(w).Encode(file); err != nil {
                                panic(err)
                            }
                        })
                    {{end}}
                {{end}}

                log.Fatal(http.ListenAndServe("{{.Host}}:{{.Port}}", nil))
            }`
}
