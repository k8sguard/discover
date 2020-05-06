FROM golang:1.14.2-stretch as BUILD_STAGE
COPY . /src
WORKDIR /src
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o discover ./cmd/discover/
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=BUILD_STAGE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILD_STAGE /src/discover .
CMD ["./discover"]