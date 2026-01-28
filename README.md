# Golang Full CI/CD Template

A production-ready Gin Go framework application with full CI/CD pipeline, dockerized PostgreSQL and Redis, and a complete development environment setup.

## 🚀 Features

- **Gin Web Framework**: Fast and lightweight HTTP web framework
- **PostgreSQL Database**: Fully integrated with GORM ORM
- **Redis Cache**: For high-performance caching
- **Docker Support**: Both production and development environments
- **Hot Reload**: Development mode with automatic code reloading using Air
- **CI/CD Pipeline**: GitHub Actions workflows for testing, linting, and building
- **Health Checks**: Built-in health check endpoints
- **RESTful API**: Clean API structure with versioning
- **Environment Configuration**: Easy configuration management

## 📋 Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- Make (optional, for convenience commands)

## 🏗️ Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── cache/
│   │   └── redis.go          # Redis connection
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── database/
│   │   └── database.go       # Database connection
│   ├── handlers/
│   │   ├── handlers.go       # HTTP handlers
│   │   └── handlers_test.go  # Handler tests
│   └── models/
│       └── user.go           # Data models
├── .github/
│   └── workflows/            # CI/CD workflows
├── docker-compose.yml        # Production compose file
├── docker-compose.dev.yml    # Development compose file
├── Dockerfile                # Production Dockerfile
├── Dockerfile.dev            # Development Dockerfile
├── Makefile                  # Convenience commands
└── .env.example              # Environment variables template
```

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**:
   ```bash
   git clone https://github.com/shubham21155102/golang-full-ci-cd-template.git
   cd golang-full-ci-cd-template
   ```

2. **Start the development environment**:
   ```bash
   make dev-up
   ```
   Or without make:
   ```bash
   docker-compose -f docker-compose.dev.yml up --build
   ```

3. **Access the application**:
   - API: http://localhost:8080
   - Health Check: http://localhost:8080/health

### Local Development

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Start PostgreSQL and Redis** (if not using Docker):
   ```bash
   docker-compose -f docker-compose.dev.yml up postgres redis
   ```

3. **Run the application**:
   ```bash
   go run cmd/api/main.go
   ```

## 📝 Environment Variables

Copy `.env.example` to `.env` and modify as needed:

```bash
cp .env.example .env
```

Available variables:
- `SERVER_PORT`: API server port (default: 8080)
- `ENV`: Environment (development/production)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `REDIS_HOST`: Redis host
- `REDIS_PORT`: Redis port

## 🔧 Available Commands

Using Make:

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application locally
make test          # Run tests
make coverage      # Generate coverage report
make lint          # Run linters
make docker-up     # Start production environment
make docker-down   # Stop production environment
make dev-up        # Start development environment
make dev-down      # Stop development environment
```

## 🌐 API Endpoints

### Health Check
```http
GET /health
```

Response:
```json
{
  "status": "ok",
  "database": "connected",
  "redis": "connected",
  "time": "2024-01-01T12:00:00Z"
}
```

### User Management

#### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### Get User
```http
GET /api/v1/users/:id
```

#### List Users
```http
GET /api/v1/users
```

## 🧪 Testing

Run tests:
```bash
go test -v ./...
```

Run tests with coverage:
```bash
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🐳 Docker

### Production Build

Build and run production environment:
```bash
docker-compose up --build
```

### Development with Hot Reload

The development environment uses Air for automatic code reloading:
```bash
docker-compose -f docker-compose.dev.yml up --build
```

Any changes to `.go` files will automatically rebuild and restart the application.

## 🔄 CI/CD Pipeline

The project includes GitHub Actions workflows for:

1. **Linting** (`lint.yml`): Code quality checks with golangci-lint
2. **Testing** (`test.yml`): Automated tests with PostgreSQL and Redis services
3. **Building** (`build.yml`): Application build verification
4. **Docker** (`docker.yml`): Docker image building and pushing to GitHub Container Registry
5. **CI** (`ci.yml`): Combined workflow for complete CI pipeline

All workflows run automatically on push and pull requests to `main` and `develop` branches.

## 📊 Database Migrations

Migrations run automatically on application startup using GORM's AutoMigrate feature. To add new models:

1. Create the model in `internal/models/`
2. Add it to the AutoMigrate call in `cmd/api/main.go`

## 🔐 Security

- Environment variables for sensitive data
- PostgreSQL with SSL support
- Redis with password authentication
- Docker secrets support ready

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [go-redis](https://github.com/redis/go-redis)
- [Air](https://github.com/air-verse/air) for hot reloading

## 📧 Contact

For questions or feedback, please open an issue on GitHub.