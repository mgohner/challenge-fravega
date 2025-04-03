# Challenge Fravega API

This repository contains the API service for the Fravega challenge.

## Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose

## Quick Start

### Running Development Environment

1. Run the application with docker:

```bash
make docker-run
```

2. Run the application on local with go:

```bash
make run
```

To start all dependencies:

```bash
make docker-compose-up
```

To stop all services:

```bash
make docker-compose-down
```


### Additional Commands

- Build the application:
```bash
make build
```

- Run tests:
```bash
make test
```

- Format code:
```bash
make fmt
```

- Run linter:
```bash
make lint
```

## Docker Operations

- Build Docker image:
```bash
make docker-build
```

- Run application in Docker:
```bash
make docker-run
```

## Purchase Order Mock Operations

- Start purchase order mock server:
```bash
make po-mock-up
```

- Stop purchase order mock server:
```bash
make po-mock-down
```