FROM golang:1.21 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify

COPY . . 
RUN go build -v -o /app/bin/main

EXPOSE 8080

CMD ["./bin/main"]
