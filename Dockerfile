FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o devex ./cmd/api

FROM scratch
COPY --from=builder /app/devex /devex
ENV GIN_MODE=release
EXPOSE 10000
CMD ["./devex"]