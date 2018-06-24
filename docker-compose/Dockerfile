FROM golang:1.10 as builder
WORKDIR /go/src/github.com/brightsparc/docker-compose-golang-redis/
COPY app.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/brightsparc/docker-compose-golang-redis/app .
ENTRYPOINT /app
EXPOSE 5000
HEALTHCHECK --interval=5m --timeout=3s CMD /app -ping
