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

#### Step3 
Restart rsys service.
```bash
sudo service rsyslog restart
```

#### Step4 
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
├── helpers
│   ├── config-loader.go
│   ├── file-checker.go
│   ├── logger.go
│   ├── required-checks.go
│   ├── response-provider.go
│   ├── test
│   │   └── configuration_loader_test.go
│   └── url-extractor.go
├── internal
│   ├── handler
│   │   ├── exclude-routes.go
│   │   ├── route-handler.go
│   │   └── test
│   │       └── exclud_routes_test.go
│   ├── middleware
│   │   ├── middleware-auth.go
│   │   └── middleware-handler.go
│   └── requests
│       ├── post.go
│       ├── request-handler.go
│       └── request-manager.go
├── isdelve
│   └── nodelve.go
├── README.md
├── test
│   └── test_env_setter.go
├── third_party
│   ├── goclockauth.go
│   └── utils.go
├── tmp
└── vendor
```