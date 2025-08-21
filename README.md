# 7sol-be-challenge

This is a Go-based backend application built based on Hexagonal Architecture concept. It includes user authentication, and user management which utilize MongoDB integration.

---

## Features
- ✅ User Management (HTTP)
    - User Registration & Login
    - User Creation
    - User Inquiry
    - User Update
    - User Deletion
- ✅ User Creation & Inquiry (gRPC)
- ✅ Overall User Inquiry (Background worker)
---

## Project Structure
### Hexagonal Architecture (Ports & Adapters)
```
├── cmd/                 # Application entry point
│   └── main.go
├── internal/
│   ├── adapters/        # Inbound (HTTP) & Outbound (DB) Adapters
│   ├── app/             # Ports and Services
│   ├── config/          # Configuration loader
│   ├── domain/          # Core domain models and interfaces
├── pkg/                 # Utilities (e.g., error handling, validator)
├── proto
│   └── user.proto
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
### HTTP
#### Public Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | /auth/register | Register a user     |
| POST   | /auth/login    | Login and get token |

#### Private Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| GET   | /api/user/list    | List all users |
| GET   | /api/user/:uid    | Retrieve specific user |
| POST   | /api/user/    | Create user |
| PUT   | /api/user/:uid    | Update user |
| DELETE   | /api/user/:uid    | Delete user |

### gRPC
#### Service: `UserService`

| Method       | Request Message       | Response Message       | Description               |
|--------------|-----------------------|------------------------|---------------------------|
| `CreateUser` | `CreateUserRequest`   | `CreateUserResponse`         | Create a new user         |
| `GetUser`  | `GetUserRequest`    | `GetUserResponse`    | Retrieve specific user via id |

## Playaround
For HTTP server, please feel free to import attached postman collection from following directory 
```/postman/7sol-be-challenge-http.postman_collection.json```

For gRPC server, both methods require ```JWT token``` from http login response and include it in to metadata with the key of ```authorization``` 
<br>eg.

Method: GetUser
```
grpcurl -plaintext -H "authorization: Bearer {JWT_TOKEN}” -d '{"id”:”{uid}}’ localhost:50051 user.UserService/GetUser
```
Note(s): 
- please replace {JWT_TOKEN} with the token from login response
- please replace {uid} with the id of specific user