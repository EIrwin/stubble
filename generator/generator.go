package generator

import (
	"bytes"
	"html/template"
	"os"

	"github.com/eirwin/stubby/config"
	"github.com/eirwin/stubby/endpoints"
)

const (
	templateName = "Stubby"
)

type Generator struct {
	Config config.Config
}

type generatorContext struct {
	Host        string
	Port        string
	EndpointMap map[string][]endpoints.Endpoint
}

func (g Generator) Run() error {
	endpoints, err := parseEndpoints(g.Config)
	if err != nil {
		panic(err)
	}

	endpointMap := generateEndpointMap(endpoints)

	context := generatorContext{
		Host:        g.Config.Host,
		Port:        g.Config.Port,
		EndpointMap: endpointMap,
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

func generateEndpointMap(definitions []endpoints.Endpoint) map[string][]endpoints.Endpoint {
	endpointMap := make(map[string][]endpoints.Endpoint, 0)
	for _, def := range definitions {
		_, ok := endpointMap[def.Path]
		var defs []endpoints.Endpoint
		if ok {
			defs = endpointMap[def.Path]
		} else {
			defs = make([]endpoints.Endpoint, 0)
		}
		defs = append(defs, def)
		endpointMap[def.Path] = defs
	}
	return endpointMap
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
	return `
package main 

import (
    "net/http"
    "io/ioutil"
    "log"
)

func main(){
log.Println("Running stubby server on {{.Host}}:{{.Port}}")
{{with .EndpointMap}}
    {{ range $key, $value := . }}
        {{with $value}}
            http.HandleFunc("{{ $key }}", func(w http.ResponseWriter, r *http.Request) {
                {{ range .}}
                    if "{{ .Method }}" == r.Method {
                        file,_ := ioutil.ReadFile("{{.FilePath}}")
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.Write(file)
                    }
                {{end}}
            })
        {{end}}
    {{end}}
{{end}}
log.Fatal(http.ListenAndServe("{{.Host}}:{{.Port}}", nil))
}`
}
