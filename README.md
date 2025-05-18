# Task Service API

A RESTful task management API built with Go and Gin framework.

## Requirements

- Go 1.18+
- Docker (optional)

## Running the Application

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

The server will start at `http://localhost:8080`

### Using Docker

1. Build the Docker image:
```bash
docker build -t task-service .
```

2. Run the container:
```bash
docker run -p 8080:8080 task-service
```

## API Endpoints

- `GET /tasks` - Get all tasks
- `POST /tasks` - Create a new task
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task

### Task Model

```json
{
  "id": "string",
  "name": "string",
  "status": 0  // 0: incomplete, 1: complete
}
```

## Running Tests

```bash
go test ./... -v
```
