# NATS Listener Service

This project is a simple Go service that listens to and processes messages received via NATS. The service subscribes to a specific NATS subject and applies a processing function to every received message.

## Project Structure

```bash
nats-listener-service/
├── cmd/
│   └── main.go                 # Main entry point for the service
├── internal/
│   └── listener/
│       └── listener.go         # Service that listens to and processes NATS messages
├── pkg/
│   └── nats/
│       └── nats.go             # Helper function for NATS connection
├── Dockerfile                  # Dockerfile for packaging the service
└── go.mod                      # Go module and dependency file
```

## Requirements

Go 1.20+
Docker (optional, for running with Docker)
NATS Server

## Installation

Clone the repository and download the necessary dependencies:

```bash
git clone https://github.com/yourusername/nats-listener-service.git
cd nats-listener-service
go mod tidy
```

## Configuration

The service is configured to connect to a NATS server and listen to a specific subject. You can modify the subject or other configurations directly in the main.go or listener.go files.

## Running the Service

Locally
To run the service locally, make sure your NATS server is running and then execute:

```bash
go run cmd/main.go
```

## Docker

To build and run the service using Docker:

1. Build the Docker image:

```bash
docker build -t nats-listener-service .
```

2. Run the Docker container:

```bash
docker run --rm -e NATS_URL=nats://localhost:4222 nats-listener-service
```

Ensure that the NATS server is accessible at the specified URL.

## Usage

Once the service is running, it will listen to the order.created subject (or any other subject you configure) and process incoming messages.

### Sending Test Messages

You can test the service by sending a message to the subscribed subject. Here is an example using Go:

```go
package main

import (
    "log"
    "github.com/nats-io/nats.go"
)

func main() {
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatalf("Error connecting to NATS: %v", err)
    }
    defer nc.Close()

    err = nc.Publish("order.created", []byte("Order123"))
    if err != nil {
        log.Fatalf("Error publishing message: %v", err)
    }

    log.Println("Published message to order.created")
}
```

## Logs

The service logs every received message and its processing status. Check the logs to verify the processing of incoming messages.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
