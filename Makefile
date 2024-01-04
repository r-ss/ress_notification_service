build:
	go get ./...
	go mod vendor
	env GOARCH=amd64 GOOS=linux GIN_MODE=release go build -ldflags="-s -w" -o bin/telegram telegram/main.go

clean:
	rm -rf ./bin ./vendor

dockerbuild:
	docker build -t ress/notificationservice .

docker_run:
	docker run -p 5008:8080 ress/notificationservice

docker_daemon:
	docker run -d -p 5008:8080 ress/notificationservice