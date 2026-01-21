# ShakerQueue

A distributed Go backend for managing cocktail orders using PostgreSQL,
RabbitMQ, JWT authentication, and Docker.

------------------------------------------------------------------------

## Overview

ShakerQueue is a backend system designed to manage cocktail orders in a
bar environment.\
It supports multiple user roles (admin, employee, customer), secure
authentication, and asynchronous order processing using a message queue.

The system is built as a distributed application with:

-   An HTTP API server\
-   A background consumer worker\
-   A PostgreSQL database\
-   A RabbitMQ message queue

This architecture allows the system to scale, remain responsive under
load, and cleanly separate responsibilities between services.

------------------------------------------------------------------------

## Architecture

    Client → API Server → RabbitMQ → Consumer → PostgreSQL

-   **API Server**\
    Handles authentication, order creation, recipe management, and user
    management.

-   **Consumer Service**\
    Processes orders asynchronously from RabbitMQ.

-   **RabbitMQ**\
    Decouples order submission from processing.

-   **PostgreSQL**\
    Stores users, recipes, and orders.

The services are fully containerized using Docker and orchestrated with
Docker Compose.

------------------------------------------------------------------------

## Features

-   JWT-based authentication\
-   Role-based access control (admin, employee, customer)\
-   Secure password hashing\
-   Asynchronous order processing\
-   REST API\
-   Type-safe SQL queries via SQLC\
-   Dockerized services\
-   RabbitMQ message queue\
-   PostgreSQL persistence

------------------------------------------------------------------------

## Tech Stack

-   **Language:** Go (Golang)\
-   **Database:** PostgreSQL\
-   **Message Queue:** RabbitMQ\
-   **Auth:** JWT\
-   **SQL Layer:** SQLC\
-   **Containerization:** Docker & Docker Compose

------------------------------------------------------------------------

## Getting Started

### Prerequisites

-   Docker\
-   Docker Compose

### Run the project

From the project root:

``` bash
docker compose up --build
```

### Services

-   API Server: http://localhost:8080\
-   RabbitMQ UI: http://localhost:15672
    -   Username: `guest`\
    -   Password: `guest`

------------------------------------------------------------------------

## API Examples

### Login

**POST** `/login`

Request:

``` json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:

``` json
{
  "user": {
    "id": "uuid",
    "role": "customer"
  },
  "token": "jwt-token"
}
```

------------------------------------------------------------------------

### Create Order

**POST** `/orders`

``` json
{
  "recipe": "Margarita"
}
```

The order is queued in RabbitMQ and processed asynchronously by the
consumer service.

------------------------------------------------------------------------

## Design Decisions

### Why RabbitMQ?

Orders are processed asynchronously to avoid blocking API requests and
to allow horizontal scaling of workers.\
This improves responsiveness and reliability under load.

### Why SQLC?

SQLC provides compile-time safety for SQL queries while keeping full
control over the database schema and SQL performance.

### Why JWT?

JWT enables stateless authentication, making it suitable for distributed
systems with multiple services.

### Why Docker?

Docker ensures consistent development and deployment environments and
simplifies service orchestration.

------------------------------------------------------------------------

## Project Structure

    cmd/
      server/     # API server entry point
      consumer/   # Worker service
      seed/       # Database seeding

    internal/
      app/        # Dependency wiring
      auth/       # Authentication logic
      database/   # SQLC-generated queries
      handlers/   # HTTP handlers
      queue/      # RabbitMQ logic

    sql/
      schema/     # Database migrations
      queries/    # SQL queries

------------------------------------------------------------------------

## Future Improvements

-   Add automated tests\
-   Implement graceful shutdown\
-   Add observability (metrics + tracing)\
-   Improve retry logic for failed jobs\
-   Add CI/CD pipeline\
-   Add rate limiting

------------------------------------------------------------------------

## Why This Project Exists

ShakerQueue was built to demonstrate real-world backend engineering
concepts:

-   Service separation\
-   Asynchronous processing\
-   Secure authentication\
-   Database-driven design\
-   Containerized deployment

It goes beyond simple CRUD apps and reflects patterns used in production
systems.

------------------------------------------------------------------------

## License

MIT
