# API Gateway

This project is an API Gateway service built using the Echo framework in Go. The API Gateway acts as an entry point for various API requests, providing routing, middleware, and other functionalities.

## Features

-   Routing using Echo framework
-   Middleware support
-   Environment-based configuration
-   Logging
-   Simple default route

## Getting Started

### Prerequisites

-   Go 1.16+
-   Git
-   Make sure to set up the required environment variables as described below.

### Rsys Logging setup

Setting up the rsys logging.

#### Step #1

Create api-gateway rsys conf file.

```bash
sudo /etc/rsyslog.d/api-gateway.conf
```

#### Step #2

Add Below content to the `api-gateway.conf`.

```bash
if $programname == 'api-gateway' then /var/log/api-gateway.log
& stop
```

#### Step #3

Restart rsys service.

```bash
sudo service rsyslog restart
```

#### Step #4

Run the following command to create log file and give required permissions.

```bash
sudo touch /var/log/api-gateway.log
sudo chown syslog:adm /var/log/api-gateway.log
sudo chmod 664 /var/log/api-gateway.log
```

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

-   `DEBUG=true`
-   `PORT=3002`
-   `GO_CLOCK_SERVER=http://localhost:3001`
-   `CONFIG_PATH=/home/vikas/projects/api-gateway/config.toml`
-   `LOG_FILENAME="api-gateway.log"`
-   `ENVIRONMENT="development"`
-   `QUEUE_CONCURRENT_WORKERS=10`
-   `QUEUE_CRITICAL_WORKERS=6`
-   `QUEUE_DEFAULT_WORKERS=3`
-   `QUEUE_LOW_WORKERS=1`
-   `QUEUE_REDIS_ADDRESS=""`
-   `RATE_LIMIT_REQ_PER_SEC=10`
-   `RATE_LIMIT_BURST_REQUEST=30`
-   `RATE_LIMIT_EXPIRES_IN=3`

### Running the Application

1. Set the environment variables:

    ```sh
    export PORT=3002
    export DEBUG=true
    export GO_CLOCK_SERVER=http://localhost:3001
    export CONFIG_PATH=/home/vikas/projects/api-gateway/config.toml
    export LOG_FILENAME="api-gateway.log"
    export ENVIRONMENT="development"
    export QUEUE_CONCURRENT_WORKERS=10
    export QUEUE_CRITICAL_WORKERS=6
    export QUEUE_DEFAULT_WORKERS=3
    export QUEUE_LOW_WORKERS=1
    export QUEUE_REDIS_ADDRESS=""
    export RATE_LIMIT_REQ_PER_SEC=10
    export RATE_LIMIT_BURST_REQUEST=30
    export RATE_LIMIT_EXPIRES_IN=3
    ```

2. Run the application:

    ```sh
    go run cmd/main/main.go
    ```

3. The API Gateway will be accessible at `http://localhost:3002`.

### Project Structure

```plaintext
.
├── cmd
│   └── main
│       ├── main.go
│       └── main_test.go
├── config.toml
├── go.mod
├── go.sum
├── helpers
│   ├── config-loader.go
│   ├── constant
│   │   └── constants.go
│   ├── converters.go
│   ├── file-checker.go
│   ├── logger.go
│   ├── path-utils.go
│   ├── required-checks.go
│   ├── response-provider.go
│   └── url-extractor.go
├── internal
│   ├── handler
│   │   ├── exclude-routes.go
│   │   ├── route-handler.go
│   │   └── test
│   │       ├── cuncurrent_request_test.go
│   │       └── exclud_routes_test.go
│   ├── middleware
│   │   ├── middleware-auth.go
│   │   ├── middleware-handler.go
│   │   └── middleware-rate-limiter.go
│   └── requests
│       ├── post.go
│       ├── request-handler.go
│       └── request-manager.go
├── isdelve
│   └── nodelve.go
├── project-structure.txt
├── README.md
├── request_load_test.go
├── test
│   └── test_env_setter.go
├── third_party
│   └── asyncq.go
└── tmp
    └── main
```

### Detailed Description of Some Key Files and Directories

-   **cmd/main**: Contains the entry point of the application.

    -   `main.go`: The main entry point of the API Gateway.
    -   `main_test.go`: Tests for the main application.

-   **helpers**: Contains utility functions and helpers used across the application.

    -   `config-loader.go`: Loads the configuration from the specified path.
    -   `file-checker.go`: Utility for checking file existence and permissions.
    -   `logger.go`: Logging utility.
    -   `path-utils.go`: Utilities for handling file paths.
    -   `required-checks.go`: Checks for required environment variables and configurations.
    -   `response-provider.go`: Provides standard responses for the API.
    -   `url-extractor.go`: Utility for extracting URL parameters and paths.

-   **internal**: Contains the core logic of the application.

    -   **handler**: Handles the routing and request handling.
        -   `exclude-routes.go`: Manages routes that should be excluded from certain middleware.
        -   `route-handler.go`: Main route handler for the API requests.
        -   **test**: Contains tests for the handlers.
            -   `concurrent_request_test.go`: Tests for concurrent request handling.
            -   `exclude_routes_test.go`: Tests for route exclusion functionality.
    -   **middleware**: Contains middleware functions.
        -   `middleware-auth.go`: Middleware for handling authentication.
        -   `middleware-handler.go`: General middleware handler.
    -   **requests**: Manages the request processing.
        -   `post.go`: Handles POST requests.
        -   `request-handler.go`: Main request handler.
        -   `request-manager.go`: Manages the lifecycle of requests.

-   **isdelve**: Contains utilities for debugging with Delve.

    -   `nodelve.go`: Utility to check if running with Delve.

-   **test**: Contains testing utilities and setup.

    -   `test_env_setter.go`: Sets up the environment for tests.

-   **third_party**: Contains third-party integrations and utilities.

    -   `goclockauth.go`: Integration for Go Clock authentication.
    -   `utils.go`: General utilities for third-party integrations.

-   **config.toml**: Configuration file for the application.

-   **go.mod**: Go module file.

-   **go.sum**: Go dependencies file.

-   **README.md**: This file.

-   **request_load_test.go**: Load testing for the request handling.
