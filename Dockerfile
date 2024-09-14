# Build the KOEC binary
FROM golang:1.22.0 as plugin-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o koec main.go