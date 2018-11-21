.PHONY: deps
deps:
	go get -d -v ./...
	go get golang.org/x/oauth2
	go get github.com/google/go-github/github

.PHONY: build
build:
	go install

