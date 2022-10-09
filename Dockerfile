FROM golang:1.19-bullseye

WORKDIR /src
COPY src .

RUN go build -o app main.go

EXPOSE 8080

ENTRYPOINT ["./app"]