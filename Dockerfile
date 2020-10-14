FROM golang:1.14.1-alpine3.11 as builder

ENV GO111MODULE=on

WORKDIR /go/src/wikigo

COPY go.sum go.mod /go/src/wikigo/

RUN go mod download

COPY . /go/src/wikigo/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/wikigo .

FROM alpine

COPY --from=builder /go/src/wikigo/bin/wikigo /

EXPOSE 8080

ENTRYPOINT ["/wikigo"]
