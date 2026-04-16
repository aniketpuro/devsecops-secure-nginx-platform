# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Final secure stage
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /
EXPOSE 8080
USER 1001
ENTRYPOINT ["/main"]