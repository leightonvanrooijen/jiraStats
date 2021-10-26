FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go mod download


# ENTRYPOINT go build && ./server

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT go build && CompileDaemon -command="./server"