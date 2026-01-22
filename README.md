# ShakerQueue

A distributed Go backend for managing cocktail orders in a bar environment,
built with PostgreSQL, RabbitMQ, JWT authentication, and Docker.

---

## Motivation

ShakerQueue was inspired by real-world experience working in the hospitality industry.
Bars are high-throughput, time-sensitive environments where efficiency, clarity, and
resilience matter. Drink orders often pile up during rushes, and poorly designed systems
can slow service, frustrate staff, and impact customer experience.

This project simulates a realistic bar order workflow by modeling:

- Asynchronous drink preparation
- Clear separation between order intake and fulfillment
- Role-based access for admins, employees, and customers
- A system that stays responsive even under load

The goal is to reflect how a production-grade backend might support real-world drink
operations behind the bar.

---

## Overview

ShakerQueue is a distributed backend system designed to manage cocktail orders.
It supports multiple user roles (admin, employee, customer), secure authentication,
and asynchronous order processing using a message queue.

The system is composed of:

- An HTTP API server
- A background consumer worker
- A PostgreSQL database
- A RabbitMQ message queue

This architecture allows the system to scale horizontally, remain responsive under load,
and cleanly separate responsibilities between services.

---

## Architecture

```
Client
  ↓
API Server
  ↓
RabbitMQ
  ↓
Consumer Worker
  ↓
PostgreSQL
```

### Components

- **API Server**  
  Handles authentication, order creation, recipe management, and user management.

- **Consumer Service**  
  Processes orders asynchronously from RabbitMQ.

- **RabbitMQ**  
  Decouples order submission from order processing.

- **PostgreSQL**  
  Stores users, recipes, and orders.

All services are fully containerized using Docker and orchestrated with Docker Compose.

---

## Features

- JWT-based authentication
- Role-based access control (admin, employee, customer)
- Secure password hashing
- Asynchronous order processing
- RESTful API
- Type-safe SQL queries via SQLC
- Dockerized services
- RabbitMQ message queue
- PostgreSQL persistence

---

## Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Message Queue:** RabbitMQ
- **Authentication:** JWT
- **SQL Layer:** SQLC
- **Containerization:** Docker & Docker Compose

---

## Quick Start

### Prerequisites

- Docker
- Docker Compose

### Run the Project

From the project root:

```bash
docker compose up --build
```

### Available Services

- API Server: http://localhost:8080
- RabbitMQ Management UI: http://localhost:15672  
  - Username: `guest`  
  - Password: `guest`

---

## Usage

### Order Flow (Design Overview)

```
Customer places order
        ↓
API validates request & auth
        ↓
Order published to RabbitMQ
        ↓
Consumer picks up order
        ↓
Order persisted in PostgreSQL
        ↓
Employee can view/process order
```

This flow ensures the API remains responsive while drink preparation is handled
asynchronously in the background.

---

## API Examples

### Login

**POST** `/login`

Request:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "user": {
    "id": "uuid",
    "role": "customer"
  },
  "token": "jwt-token"
}
```

---

### Create Order

**POST** `/orders`

Request:

```json
{
  "recipe": "Margarita"
}
```

The order is queued in RabbitMQ and processed asynchronously by the consumer service.

---

## Design Decisions

### Why RabbitMQ?

Orders are processed asynchronously to avoid blocking API requests and to allow
horizontal scaling of workers. This improves responsiveness and reliability during
peak usage.

### Why SQLC?

SQLC provides compile-time safety for SQL queries while maintaining full control
over schema design and query performance.

### Why JWT?

JWT enables stateless authentication, which works well in distributed systems
with multiple services.

### Why Docker?

Docker ensures consistent development and deployment environments and simplifies
local orchestration of multiple services.

---

## Project Structure

```
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
```

---

## Contributing

### Clone the Repository

Using Homebrew (macOS):

```bash
brew install git
```

Clone the repository:

```bash
git clone https://github.com/your-username/ShakerQueue.git
cd ShakerQueue
```

Run the project locally:

```bash
docker compose up --build
```

Contributions, issues, and feature requests are welcome.

---

## Future Improvements

- Add automated tests
- Implement graceful shutdown
- Add observability (metrics and tracing)
- Improve retry logic for failed jobs
- Add CI/CD pipeline
- Add rate limiting

---

## Why This Project Exists

ShakerQueue was built to demonstrate real-world backend engineering concepts:

- Service separation
- Asynchronous processing
- Secure authentication
- Database-driven design
- Containerized deployment

It goes beyond simple CRUD applications and reflects patterns commonly used in
production systems.

---

## License

MIT
