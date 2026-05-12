# Build stage
FROM golang:1.23-alpine AS build
WORKDIR /app
# swag generates swagger docs from annotations
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Cache dependencies (go mod download creates go.sum if absent)
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init && CGO_ENABLED=0 go build -o server .

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]
