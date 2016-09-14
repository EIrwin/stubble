package main

import (
	"log"

	"github.com/jwaldrip/odin/cli"

	"github.com/eirwin/stubby/config"
	"github.com/eirwin/stubby/generator"
)

const (
	version  = "0.0.1"
	title    = "Stubby Mock JSON API Generator"
	pathFlag = "path"
)

var app = cli.New(version, title, func(c cli.Command) {
	if c.Flag(pathFlag) != nil {
		path := c.Flag("path").Get().(string)
		config, err := config.Read(path)
		if err != nil {
			log.Println(err.Error())
		}

		generator := generator.New(config)
		err = generator.Run()
		if err != nil {
			log.Println(err.Error())
		}
	}
})

func init() {
	app.DefineStringFlag("path", "", "Path to .yaml file defining Stubby configuration")
	app.AliasFlag('p', "path")
}

func main() {
	app.Start()
}
