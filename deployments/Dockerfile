#build stage
FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache build-base
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /go/bin/app /src/cmd/fibsrv


#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=fibsrv Version=0.0.1
EXPOSE 3000
