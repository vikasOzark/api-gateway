# API Gateway

This project is an API Gateway service built using the Echo framework in Go. The API Gateway acts as an entry point for various API requests, providing routing, middleware, and other functionalities to manage and control the flow of requests to different backend services.

## Features

- Routing using Echo framework
- Middleware support
- Environment-based configuration
- Logging
- Simple default route

## Getting Started

### Prerequisites

- Go 1.16+
- Git
- Make sure to set up the required environment variables as described below.

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/vikasOzark/api-gateway.git
    cd api-gateway
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Environment Variables

The application relies on the following environment variables:

- `PORT`: The port on which the API Gateway will run.
- `DEBUG`: Set to `true` to enable debug mode.

### Running the Application

1. Set the environment variables:

    ```sh
    export PORT=8080
    export DEBUG=true
    ```

2. Run the application:

    ```sh
    go run cmd/main/main.go
    ```

3. The API Gateway will be accessible at `http://localhost:8080`.

### Project Structure

```plaintext
├── cmd
│   └── main
│       └── main.go
├── config.toml
├── go.mod
├── go.sum
├── helpers
│   ├── config-loader.go
│   ├── file-checker.go
│   ├── logger.go
│   ├── required-checks.go
│   ├── response-provider.go
│   └── url-extractor.go
├── internal
│   ├── handler
│   │   ├── exclude-routes.go
│   │   └── route-handler.go
│   ├── middleware
│   │   ├── middleware-auth.go
│   │   └── middleware-handler.go
│   └── requests
│       ├── post.go
│       ├── request-handler.go
│       └── request-manager.go
├── isdelve
│   └── nodelve.go                  # This file contains the [Enable] variable to work with IDE debuggers. 
├── README.md
├── third_party
│   ├── goclockauth.go
│   └── utils.go
└── tmp
    └── main
```