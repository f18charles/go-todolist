# Todo API in Go

A practical example of building a flexible Go backend with clean architecture. The same business logic works with different routers and databases — no changes required.

---

## Core Idea

Write your business logic once. Swap out frameworks and databases without touching the core code.

**What you can switch:**

- **Routers:** `net/http`, Gin, or Chi
- **Databases:** SQLite or MongoDB

The service layer? It stays exactly the same, no matter what.

---

## Project Structure

```text
todolist-go/
├── cmd/
│   └── server/
│       └── main.go           # App entry point
│
└── internal/
    ├── models/
    │   └── todo.go            # Data structure
    │
    ├── repository/
    │   ├── interface.go       # What a repo must do
    │   ├── sqlite.go          # SQLite implementation
    │   └── mongo.go           # MongoDB implementation
    │
    ├── service/
    │   └── todo.go            # Business logic
    │
    ├── handler/
    │   ├── std.go             # net/http handler
    │   ├── gin.go             # Gin handler
    │   └── chi.go             # Chi handler
    │
    └── router/
        ├── std.go             # net/http routes
        ├── gin.go             # Gin routes
        └── chi.go             # Chi routes
```

---

## How the Layers Work

### 🧱 Model

Just the data structure. No HTTP, no database concerns.

```go
type Todo struct {
    ID        string
    Title     string
    Completed bool
    CreatedAt time.Time
}
```

### 💾 Repository

The data layer. Defines a clear interface so the rest of the app doesn't care what database you use.

```go
type Repository interface {
    Create(todo Todo) error
    GetByID(id string) (Todo, error)
    GetAll() ([]Todo, error)
    Update(todo Todo) error
    Delete(id string) error
}
```

### ⚙️ Service

Where the actual logic lives. It only knows about the repository interface — not the specific database.

```go
type Service struct {
    repo Repository
}

func (s *Service) CreateTodo(title string) (Todo, error) {
    // Business logic here
    // Validation, ID generation, timestamps
}
```

### 🎯 Handler

Turns HTTP requests into function calls. Different handlers for different routers, but they all call the same service methods.

### 🚦 Router

Just connects URLs to handlers. No logic, just wiring.

---

## API Endpoints

| Method   | Path          | Description           |
|----------|---------------|-----------------------|
| `POST`   | `/todos`      | Create a new todo     |
| `GET`    | `/todos`      | List all todos        |
| `PATCH`  | `/todos/{id}` | Mark a todo as complete |
| `DELETE` | `/todos/{id}` | Delete a todo         |

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

---

## Running the Project

### Quick Start

```bash
# Clone and go
git clone https://github.com/yourname/todolist-go
cd todolist-go

# Get dependencies
go mod tidy

# Run the server
go run cmd/server/main.go
```

Server starts at `http://localhost:8080`

---

## Making Changes

### Switch Routers

Just change one line in `main.go`:

```go
// Use Gin
router := router.NewGinRouter(handler)

// Or use Chi
router := router.NewChiRouter(handler)

// Or stick with net/http
router := router.NewStdRouter(handler)
```

> The service layer? Not touched.

### Switch Databases

Change how you create the repository:

```go
// SQLite
repo := repository.NewSQLite(db)

// MongoDB
repo := repository.NewMongo(collection)
```

> The service layer? Still unchanged.

---

## Why This Works

| Layer      | Responsibility       | Knows about                    |
|------------|----------------------|--------------------------------|
| Model      | Data structure       | Nothing else                   |
| Repository | Database operations  | Only the model                 |
| Service    | Business logic       | Only the repository interface  |
| Handler    | HTTP handling        | Only the service               |
| Router     | Route mapping        | Only handlers                  |

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
