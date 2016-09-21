test:
	cd endpoints && go test
	cd config && go test

install-deps:
	go get github.com/jwaldrip/odin/cli

install:
	go install