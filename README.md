# KOEC, KOng Error Customization (Status: Under Development)

## Status

We're currently facing an issue with the [PDK (Plugin Development Kit)](https://pkg.go.dev/github.com/Kong/go-pdk), specifically around locating the method to retrieve the response code for upstream handling. We're working to resolve this as quickly as possible and will resume development shortly. We apologize for the delay and appreciate your patience.

As the famous quote from Thomas Edison goes, "Genius is one percent inspiration, ninety-nine percent perspiration." We're in the hard-working phase but will get things back on track soon!

**And** if you have any suggestions on how to retrieve the response code from upstream within the `Access(kong *pdk.PDK)` method, feel free to share them. We're open to discussing potential solutions and setting a clear direction for improving this aspect of the project. Your input could help accelerate the progress and ensure we're heading in the right direction.

## Description

KOEC (Kong Error Customization) is a Go-based module designed to streamline the process of customizing error pages within the Kong API Gateway. This lightweight, efficient tool provides developers with a flexible way to create, manage, and display tailored error responses, ensuring smooth integration with Kong’s infrastructure. Built with simplicity in mind, KOEC enhances user experience by allowing structured, consistent error handling across services.

## How to use KOEC

KOEC is designed to be included in a Kong Gateway container as a binary. To use it, you need to build the KOEC application and incorporate it into Kong's Docker container. Below is an example of how to modify Kong's `Dockerfile` to include KOEC.

```Dockerfile
# Build the KOEC binary
FROM golang:1.22.0 as plugin-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o koec main.go

# Add KOEC to the Kong container
FROM kong:3.7.1
USER root
COPY --from=plugin-builder /app/koec /usr/local/bin/koec
...
```

Once the image with KOEC is built, you can configure Kong to use the KOEC plugin. Follow the configuration steps outlined in the [Kong documentation](https://docs.konghq.com/gateway/latest/plugin-development/pluginserver/go/#example-configuration). An example configuration is provided below:

```yaml
pluginserver_names = koec

pluginserver_koec_socket = /usr/local/kong/koec.socket
pluginserver_koec_start_cmd = /usr/local/bin/koec
pluginserver_koec_query_cmd = /usr/local/bin/koec -dump
```