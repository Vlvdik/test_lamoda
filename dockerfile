FROM golang:1.21.1-alpine AS builder

WORKDIR /app

COPY . .
COPY .env .  
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .  
RUN chmod +x main

CMD ["./main"]