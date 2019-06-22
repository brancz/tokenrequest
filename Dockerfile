FROM golang:1.12-alpine AS build
RUN apk add --update make && apk add --no-cache git
WORKDIR /go/src/github.com/brancz/tokenrequest
COPY . .
RUN go build && cp /go/src/github.com/brancz/tokenrequest/tokenrequest /usr/local/bin

FROM alpine:3.9
RUN apk add -U --no-cache ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build /usr/local/bin/tokenrequest .
ENTRYPOINT ["./tokenrequest"]
