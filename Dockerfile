FROM golang:1.17-alpine3.16 as builder


RUN mkdir -p /go/src/com/github/cristian-garzon/discord-bot-go
ADD . /go/src/com/github/cristian-garzon/discord-bot-go
WORKDIR /go/src/com/github/cristian-garzon/discord-bot-go

COPY go.mod .
COPY go.sum .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.16

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=0 /go/src/com/github/cristian-garzon/discord-bot-go .

ENTRYPOINT ["/root/app"]

