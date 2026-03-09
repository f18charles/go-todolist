# TodoList API in Go

A practical example of building a flexible Go backend with clean architecture. The same business logic works with different routers and databases — no changes required.

---

## Core Idea

Write your business logic once. Swap out frameworks and databases without touching the core code.

**What you can switch:**

- **Routers:** `net/http`, Gin, or Chi
- **Databases:** SQLite or MongoDB

> Think of your phone. You can change the screen protector and phone case. In the same way, you can change the router and repository(database facing layer) but the service logic will remain the same.

---

## Project Structure

```text
go-todolist/
├── cmd/
│   ├── server_sql/
│   │   └── main.go           # Entry point — SQLite + Chi
│   └── server_mongo/
│       └── main.go           # Entry point — MongoDB + net/http
│
└── internal/
    ├── database/
    │   └── db.go              # SQLite and MongoDB init
    │
    ├── models/
    │   └── todo.go            # Data structure
    │
    ├── repository/
    │   ├── todo_repository.go # Repository interface
    │   ├── sqlite_repo.go     # SQLite implementation
    │   └── mongo_repo.go      # MongoDB implementation
    │
    ├── service/
    │   └── todo_service.go    # Business logic
    │
    ├── handler/
    │   ├── std_handler.go     # net/http handler
    │   ├── gin_handler.go     # Gin handler
    │   └── chi_handler.go     # Chi handler
    │
    └── router/
        ├── std_router.go      # net/http routes
        ├── gin_router.go      # Gin routes
        └── chi_router.go      # Chi routes
```

---

## How the Layers Work

### 🧱 Model

Just the data structure.

```go
type Todo struct {
    ID        string
    Title     string
    Completed bool
    CreatedAt time.Time
}
```

### 💾 Repository

The data layer. Defines a clear interface for the rest of the app.

```go
type TodoRepository interface {
    Create(todo Todo) error
    GetByID(id string) (Todo, error)
    GetAll() ([]Todo, error)
    Update(todo Todo) error
    Delete(id string) error
}
```

### ⚙️ Service

Where the actual logic lives. It only knows about the repository interface.

```go
type TodoService struct {
    Repo repository.TodoRepository
}

func (s *TodoService) Create(title string) (Todo, error) {
    // Validation, ID generation, timestamps
}
```

### 🎯 Handler

Turns HTTP requests into function calls. Different handlers for different routers, but they all call the same service methods.

### 🚦 Router

Just connects URLs to handlers. No logic, just wiring.

---

## API Endpoints

| Method   | Path          | Description             |
|----------|---------------|-------------------------|
| `POST`   | `/todos`      | Create a new todo       |
| `GET`    | `/todos`      | List all todos          |
| `PATCH`  | `/todos/{id}` | Mark a todo as complete |
| `DELETE` | `/todos/{id}` | Delete a todo           |

---

## Setup & Running

### Prerequisites

- Go 1.21+
- Docker (optional, for MongoDB)

### Install Dependencies

```bash
git clone https://github.com/f18charles/go-todolist.git
cd go-todolist
go mod tidy
```

## Running the Servers

### Run with SQLite + Chi

```bash
go run cmd/server_sql/main.go
```

This uses:
- **Database:** SQLite (creates `todos.db` in the project root automatically)
- **Router:** Chi
- **Handler:** Chi handler

### Run with MongoDB + net/http

First, start MongoDB. The easiest way is with Docker:

```bash
docker run -d --name mongo -p 27017:27017 mongo:latest
```

Then run the server:

```bash
go run cmd/server_mongo/main.go
```

This uses:
- **Database:** MongoDB (`localhost:27017`, database `todolist`, collection `todos`)
- **Router:** net/http standard library
- **Handler:** Standard handler

> **Note:** These two entry points use different router/handler combinations intentionally —
> to demonstrate that any database can be paired with any router without changing the service layer.

Both servers start at `http://localhost:8080`.

---

## Example Requests

**Create a todo**

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"learn Go architecture"}'
```

**Response**

```json
{
  "id": "abc-123",
  "title": "learn Go architecture",
  "completed": false,
  "created_at": "2024-01-15T10:30:00Z"
}
```

**List all todos**

```bash
curl http://localhost:8080/todos
```

**Mark a todo as complete**

```bash
curl -X PATCH http://localhost:8080/todos/abc-123
```

**Delete a todo**

```bash
curl -X DELETE http://localhost:8080/todos/abc-123
```

---

## All Valid Combinations

Any handler and router can be paired with any database. Here's the full matrix:

| Database | Router     | Handler     | How to wire it up                                    |
|----------|------------|-------------|------------------------------------------------------|
| SQLite   | Chi        | Chi         | `server_sql/main.go` (default)                       |
| SQLite   | Gin        | Gin         | swap `NewChiHandler` → `NewGinHandler`, `NewChiRouter` → `NewGinRouter`  |
| SQLite   | net/http   | Std         | swap `NewChiHandler` → `NewStdHandler`, `NewChiRouter` → `NewStdRouter`  |
| MongoDB  | net/http   | Std         | `server_mongo/main.go` (default)                     |
| MongoDB  | Gin        | Gin         | swap `NewStdHandler` → `NewGinHandler`, `NewStdRouter` → `NewGinRouter`  |
| MongoDB  | Chi        | Chi         | swap `NewStdHandler` → `NewChiHandler`, `NewStdRouter` → `NewChiRouter`  |

All swaps happen in `main.go` only. The service layer is never touched.

---

## Why This Works

| Layer      | Responsibility      |
|------------|---------------------|
| Model      | Data structure      |
| Repository | Database operations |
| Service    | Business logic      |
| Handler    | HTTP handling       |
| Router     | Route mapping       |

Each layer has one job and only depends on the layer below it.

---

## What You Can Learn

This project demonstrates:

- **Separation of concerns** — each part has a clear role
- **Interface-driven design** — depend on abstractions, not implementations
- **Testability** — mock repositories make service testing easy
- **Flexibility** — swap databases and routers without rewriting

---

## Next Steps & Ideas

Want to build on this? Try adding:

- ✅ In-memory repository for faster testing
- ✅ Pagination for listing todos
- ✅ Input validation
- ✅ Unit tests for services
- ✅ Docker setup
- ✅ Update todo endpoint (edit title)

---

## License

MIT — Use it, change it, break it, learn from it.