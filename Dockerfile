FROM golang:1.21 as builder

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

RUN mkdir /app
COPY .. /app
WORKDIR /app

ENV DATABASE_URL ${DATABASE_URL}

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
