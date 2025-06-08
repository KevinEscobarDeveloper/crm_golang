# Etapa 1: build
FROM golang:1.23 as builder

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o orchard-service ./cmd/app/main.go

# Etapa 2: imagen liviana
FROM alpine:latest

RUN adduser -D appuser
WORKDIR /home/appuser

COPY --from=builder /app/orchard-service .

USER appuser
EXPOSE 8080
EXPOSE 2222

ENTRYPOINT ["./orchard-service"]
