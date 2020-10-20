FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN GOOS=linux go build -o todo ./todo-service/cmd/todo/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/todo .

CMD ["./todo"]