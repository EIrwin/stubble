test:
	cd endpoints/test && go test
	cd config/test && go test

install-deps:
	go get github.com/jwaldrip/odin/cli

install:
	go install