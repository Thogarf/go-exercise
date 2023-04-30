# syntax=docker/dockerfile:1

# Default BASE_IMAGE
ARG BASE_IMAGE=gcr.io/distroless/static

# multi-stage builder
FROM golang:1.19-alpine as builder

WORKDIR $GOPATH/src/go-exercise

COPY go.mod .
COPY src/ .

RUN go mod download
RUN go mod verify

RUN go get go-exercise

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-exercise .

# finalized container
FROM $BASE_IMAGE

COPY --from=builder /go-exercise .

CMD ["./go-exercise"]