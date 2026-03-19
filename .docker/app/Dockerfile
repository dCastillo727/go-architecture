# ---- Build Stage ----
FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main-bin ./cmd/main

# ---- Final Stage ----
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /main-bin .

EXPOSE 8080

CMD ["./main-bin"]
