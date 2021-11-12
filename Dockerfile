FROM golang:1.16.3 AS golang-builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# [x509: certificate signed by unknown authority를 해결하기 위해서 CA certificates 추가](https://velog.io/@byron1st/x.509-certificate-signed-by-unknown-authority)
RUN apk update && apk upgrade && apk add --no-cache ca-certificates

WORKDIR /build
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o main ./

WORKDIR /dist
RUN cp /build/main ./
RUN cp /build/configs/config.prod.json ./

FROM scratch

LABEL author="parkgang[Kyungeun Park]<ruddms936@naver.com>"
LABEL version="0.1.0"

ENV GO_ENV=production

COPY --from=golang-builder /dist/main ./
COPY --from=golang-builder /dist/config.prod.json ./configs/config.prod.json
COPY --from=golang-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/main"]