# 7sol-be-challenge

This is a Go-based backend application built based on Hexagonal Architecture concept. It includes user authentication, and user management which utilize MongoDB integration.

---

## Features
- ✅ Hexagonal Architecture (Ports & Adapters)
- ✅ User Registration & Login
- ✅ Bcrypt Password Hashing
- ✅ JWT Token Generation
- ✅ MongoDB as the data store
- ✅ Docker & Docker Compose setup
- ✅ Unit tests with GoMock and Testify

---

## Project Structure
```
├── cmd/                 # Application entry point
│   └── main.go
├── internal/
│   ├── adapters/        # Inbound (HTTP) & Outbound (DB) Adapters
│   ├── app/             # Ports and Services
│   ├── config/          # Configuration loader
│   ├── domain/          # Core domain models and interfaces
├── pkg/                 # Utilities (e.g., error handling, validator)
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## Running Locally without Docker
1. set these variables in a local .env file. (make sure mongo is running)

```
MONGO_URI=mongodb://localhost:27017/
MONGO_NAME=database
LOGGER_FORMAT=${status} - ${method} ${path} | ${time}
LOGGER_TIME_FORMAT=02-Jan-2006 15:04:05
LOGGER_TIME_ZONE=UTC
APP_TOKEN_TIMEOUT=5
APP_SECRET_KEY=example_secret_key
```

2. execute following commands
```
go mod download
go run ./cmd/main.go
```

## Running with Docker
1. execute following command
```
docker-compose up --build
```
This will start:
- api (Go backend)
- mongo (MongoDB database)

By default, the API runs on http://localhost:8080

## Running Test
```
go test -coverprofile=coverage.out ./internal/app/services
go tool cover -html=coverage.out
```

## API endpoints
### Public Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | /auth/register | Register a user     |
| POST   | /auth/login    | Login and get token |

### Private Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| GET   | /api/user/list    | List all users |
| GET   | /api/user/:uid    | Retrieve specific user |
| POST   | /api/user/    | Create user |
| PUT   | /api/user/:uid    | Update user |
| DELETE   | /api/user/:uid    | Delete user |