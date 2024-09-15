# Build the KOEC binary
FROM golang:1.22.0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o koec main.go