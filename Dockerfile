FROM golang:1.16.3 AS golang-builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

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

EXPOSE 8080

ENTRYPOINT ["/main"]