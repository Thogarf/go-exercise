FROM golang:1.19-alpine

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app
COPY go.mod ./

COPY . .



RUN go build -o goExcercise src/main.go 

CMD ./goExcercise

