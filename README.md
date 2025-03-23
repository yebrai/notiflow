# NotiFlow

![Maintenance](https://img.shields.io/badge/Status-Under%20Maintenance-yellow) ![Go Version](https://img.shields.io/badge/Go-1.23-blue)

<div align="center">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" width="200">
  <br>
  <em>A modern, event-driven notification management system</em>
</div>

## ⚠️ Development Status

This project is currently **under active development** and is not yet ready for production use. Major architectural changes and features are being implemented.

## 🚀 Features (Planned/In Progress)

- Multi-channel notification delivery (Email, SMS, Push)
- Event-driven architecture
- Configurable templates and delivery rules
- Retry mechanisms for failed notifications
- Delivery status tracking
- REST API for notification management

## 🔧 Technology Stack

- Go 1.23
- Domain-Driven Design (DDD)
- Event-Driven Architecture
- GORM (PostgreSQL)
- Gin Web Framework
- Hexagonal Architecture
- Docker & Docker Compose
- Wire for Dependency Injection

## 📋 Prerequisites

- Go 1.23+
- Docker and Docker Compose
- PostgreSQL 17

## 🚀 Getting Started

### Development Environment

```bash
# Clone the repository
git clone https://github.com/yrebai/notification-service.git
cd notification-service

# Start PostgreSQL (only DB, for local development)
docker-compose -f docker-compose.dev.yaml up -d

# Install dependencies
go mod tidy

# Run the application
go run ./cmd/api
```

### Running with Docker

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f
```

## 🏗️ Project Structure

```
notification-service/
├── cmd/
│   └── api/             # Application entry points
├── internal/
│   ├── domain/          # Domain entities and interfaces
│   ├── application/     # Use cases and application services
│   ├── infrastructure/  # External implementations (DB, email)
│   └── config/          # Configuration
├── scripts/             # DB migrations and utility scripts
└── docker-compose.yaml  # Container orchestration
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

---
