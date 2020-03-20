.PHONY: build tidy test package watch devsetup

build:
	go build ./...

test: build
	go test ./...

tidy: build
	go mod tidy
	go fmt ./...

watch: build
	modd

package:
	go-bindata -o bot/templates.go -pkg bot data/
	docker build . -t covidzero/slackbot:latest

run: package
	docker-compose up

devsetup:
	mv go.mod go.mod
	GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
	go install github.com/cortesi/modd/cmd/modd
