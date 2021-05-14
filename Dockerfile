# [해당 문서를 참고하였습니다](https://thesorauniverse.com/posts/kr/golang/making-golang-docker-img-best-practices/)
FROM golang:1.16.3 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]