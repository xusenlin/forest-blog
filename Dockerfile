FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/forest-blog

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache git && addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/forest-blog .
COPY config.json .
COPY views ./views
COPY public ./public

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 80

CMD ["./forest-blog"]
