# builder image
FROM golang:1.18.1-alpine as base

WORKDIR /builder
RUN apk add upx
COPY go.mod go.sum /builder/
RUN go mod download
COPY . .
RUN go build \
  -ldflags "-s -w" \
  -o /builder/main /builder/cmd/app/main.go
RUN upx -9 /builder/main

# runner image
FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/main main
CMD ["/app/main"]