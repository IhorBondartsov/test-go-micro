FROM golang:1.12-alpine3.9 as builder

WORKDIR /go/src/github.com/IhorBondartsov/test-k8s/auth-ms

RUN apk add git dep

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .

RUN  CGO_ENABLED=0 go build -o /build/auth-ms cmd/main.go

FROM alpine:3.9

ENV AUTH_MS_GRPCADDRESS "127.0.0.1:50003"

COPY --from=builder /build/auth-ms /app/

ENTRYPOINT ["/app/auth-ms"]