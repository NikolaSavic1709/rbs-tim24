# Mini-Zanzibar: Global Authorization System

Mini-Zanzibar is a simplified version of Google's Zanzibar, designed to provide secure, consistent, and scalable access control. This project implements a flexible configuration language for access control policies, stores and evaluates ACLs using LevelDB, and leverages ConsulDB for namespace configuration. With low latency and high availability, Mini-Zanzibar ensures efficient authorization decisions. The system includes Docker setup for easy deployment and integration with services like Nginx, Consul, and PostgreSQL.

## Getting Started

You can start system with 
```bash
docker-compose up --build
```
Before starting it is necessary to configure .env file and add your values.
Docker-compose up will start Postgres DB, Consul DB, Redis DB, three instances of go application on different ports and Nginx reverse proxy which is also load balancer for those three instances.

You can start our apps separately with make file:

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
