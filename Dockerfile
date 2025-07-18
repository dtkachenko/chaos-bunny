# Dockerfile
FROM golang:1.24.1 AS builder
WORKDIR /app
COPY . .
RUN ls -al && cd ./cmd/chaos-bunny && go build -o /chaos-bunny

FROM debian:bullseye-slim
COPY --from=builder /chaos-bunny /usr/bin/chaos-bunny
ENTRYPOINT ["/usr/bin/chaos-bunny"]
