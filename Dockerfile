FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go env -w GOTOOLCHAIN=auto GOPROXY=https://goproxy.cn,direct && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/forest-blog

FROM scratch

WORKDIR /app

COPY --from=builder /app/forest-blog .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY config.json .
COPY views ./views
COPY public ./public

EXPOSE 80

CMD ["./forest-blog"]
