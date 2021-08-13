FROM golang:1.16.0-alpine as builder

RUN apk update && \
    apk add git

ENV GOPATH=/go \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/Mshivam2409/hls-streamer/src

COPY . .

RUN go mod download && \
    go build -o gostreamer


FROM ubuntu:latest

RUN apt-get update && \
    apt-get install --yes --no-install-recommends ffmpeg && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/*

COPY --from=builder /go/src/github.com/Mshivam2409/hls-streamer/src/gostreamer /usr/local/bin/
COPY hls.yaml /root/.hls.yaml

WORKDIR /root/

CMD [ "/usr/local/bin/gostreamer", "-config", "hls.yaml" ]