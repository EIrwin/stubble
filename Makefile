test:
	cd endpoints && go test
	cd config && go test

install-deps:
	go get github.com/jwaldrip/odin/cli
	go get gopkg.in/yaml.v2

install:
	go install
