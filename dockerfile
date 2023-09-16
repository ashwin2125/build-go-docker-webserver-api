# Build Stage
FROM golang:1.19 AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY simple_go_api.go .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin .

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/bin /app/
EXPOSE 8080
ENTRYPOINT [ "/app/bin" ]
