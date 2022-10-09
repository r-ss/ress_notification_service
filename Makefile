.PHONY: build clean deploy

build:
	go get ./...
	go mod vendor
	env GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-s -w" -o bin/telegram telegram/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose