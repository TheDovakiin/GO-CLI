# Go Learning Roadmap

## Current Progress: ~30%

You've built a working CLI app with structs, slices, file I/O, error handling, and user input. Solid foundation.

---

## Phase 1: Enhance This CLI (30% → 65%)

### 1.1 Methods & Pointers
Add methods to your Task struct instead of standalone functions.

```go
// Before (what you have)
func markComplete(id int) { ... }

// After (methods)
func (t *Task) MarkComplete() {
    t.Status = true
}

func (t *Task) IsOverdue() bool {
    return time.Now().After(t.DueDate)
}

// Usage
task.MarkComplete()
if task.IsOverdue() { ... }
```

**Why:** This is how Go does "object-oriented" programming.

---

### 1.2 Maps
Replace slice with map for O(1) lookups by ID.

```go
// Before
var TasksSlice = []Task{}

// After
var TasksMap = map[int]Task{}

// Usage
TasksMap[newID] = newTask           // Add
task, exists := TasksMap[id]        // Get
delete(TasksMap, id)                // Delete
```

**Why:** No more looping to find by ID.

---

### 1.3 Packages (Project Structure)
Split your code into multiple files/folders.

```
GO-CLI/
├── main.go              // Entry point only
├── models/
│   └── task.go          // Task struct & methods
├── storage/
│   └── json.go          // Save/Load functions
└── ui/
    └── menu.go          // Display functions
```

**Why:** Real Go projects are organized this way.

---

### 1.4 Goroutines & Channels
Add background auto-save every 30 seconds.

```go
// Goroutine - runs in background
go func() {
    for {
        time.Sleep(30 * time.Second)
        saveTask()
        fmt.Println("Auto-saved!")
    }
}()

// Channel - stop the goroutine on quit
quit := make(chan bool)

go func() {
    for {
        select {
        case <-quit:
            return  // Stop goroutine
        case <-time.After(30 * time.Second):
            saveTask()
        }
    }
}()

// When user quits
quit <- true
```

**Why:** Concurrency is Go's superpower.

---

### 1.5 Interfaces
Make Task implement the Stringer interface.

```go
import "fmt"

// Stringer interface (built-in)
type Stringer interface {
    String() string
}

// Task implements it
func (t Task) String() string {
    return fmt.Sprintf("[%d] %s (Due: %s)",
        t.ID, t.Title, t.DueDate.Format("Jan 02, 2006"))
}

// Now this works automatically
fmt.Println(task)  // Calls task.String()
```

**Why:** Interfaces enable polymorphism and are used everywhere in Go.

---

### 1.6 Testing
Create `task_test.go` alongside your code.

```go
package main

import "testing"

func TestMarkComplete(t *testing.T) {
    task := Task{ID: 1, Status: false}
    task.MarkComplete()

    if task.Status != true {
        t.Errorf("Expected true, got %v", task.Status)
    }
}

func TestIsOverdue(t *testing.T) {
    past := time.Now().AddDate(0, 0, -1)  // Yesterday
    task := Task{DueDate: past}

    if !task.IsOverdue() {
        t.Error("Expected task to be overdue")
    }
}
```

Run with: `go test`

**Why:** Testing is essential for any serious Go project.

---

### 1.7 Command-Line Flags
Add flags for quick actions.

```go
import "flag"

func main() {
    addFlag := flag.String("add", "", "Quick add a task")
    listFlag := flag.Bool("list", false, "List all tasks")
    flag.Parse()

    if *addFlag != "" {
        // Add task with title from flag
    }
    if *listFlag {
        // Show all tasks
    }
}
```

Usage: `go run main.go -add "Buy milk" -list`

**Why:** Makes CLI tools more powerful.

---

## Phase 2: Build a REST API (65% → 85%)

### Project: Task Manager API

Take your same task manager and serve it over HTTP.

```go
package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/tasks", handleTasks)
    http.HandleFunc("/tasks/", handleTaskByID)
    http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(TasksSlice)
    case "POST":
        var task Task
        json.NewDecoder(r.Body).Decode(&task)
        // Add task...
    }
}
```

### What You'll Learn:
| Concept | Description |
|---------|-------------|
| `net/http` | Built-in HTTP server |
| JSON APIs | Request/response handling |
| Routing | URL patterns & methods |
| Middleware | Logging, auth, CORS |
| Database | PostgreSQL or SQLite |
| Authentication | JWT tokens |
| Environment variables | Config management |
| Docker | Containerization |

### Suggested Framework
Start with standard library, then try:
- **Chi** - lightweight router
- **Gin** - popular, fast
- **Echo** - similar to Gin

---

## Phase 3: Advanced Project (85% → 100%)

### Project: Real-Time Collaborative Task Board

Multiple users sharing tasks with live updates.

### What You'll Learn:
| Concept | Description |
|---------|-------------|
| WebSockets | Real-time communication |
| Context | Request cancellation, timeouts |
| Worker pools | Concurrent task processing |
| gRPC | High-performance RPC |
| Caching | Redis integration |
| Rate limiting | Protect your API |
| Observability | Logging, metrics, tracing |
| Generics | Type-safe data structures |
| Reflection | Advanced metaprogramming |
| Performance | Profiling & optimization |

---

## Summary: 3 Projects to 100%

| Project | Coverage | Key Concepts |
|---------|----------|--------------|
| CLI Task Manager (current) | 30% → 65% | Structs, slices, methods, goroutines, testing |
| REST API | 65% → 85% | HTTP, database, auth, middleware |
| Real-Time Collab App | 85% → 100% | WebSockets, gRPC, caching, performance |

---

## Quick Reference: What's Left After Each Phase

### After CLI (65%)
- [x] Basics (variables, types, functions)
- [x] Structs & methods
- [x] Slices & maps
- [x] Error handling
- [x] File I/O
- [x] Goroutines & channels
- [x] Interfaces
- [x] Testing
- [x] Packages
- [ ] HTTP/Web
- [ ] Database
- [ ] Advanced concurrency

### After API (85%)
- [x] Everything above
- [x] HTTP servers
- [x] JSON APIs
- [x] Database (SQL)
- [x] Authentication
- [x] Middleware
- [ ] WebSockets
- [ ] gRPC
- [ ] Advanced patterns

### After Collab App (100%)
- [x] Everything Go has to offer

---

## Resources

- [Go Tour](https://go.dev/tour/) - Interactive basics
- [Go by Example](https://gobyexample.com/) - Code snippets
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
- [Let's Go](https://lets-go.alexedwards.net/) - Web apps book
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) - Deep dive

---

## Estimated Timeline

This is self-paced. Focus on understanding, not speed.

| Phase | Suggested Focus |
|-------|-----------------|
| Phase 1 (CLI) | Until it feels natural |
| Phase 2 (API) | Until you can build any CRUD API |
| Phase 3 (Advanced) | Ongoing - pick concepts as needed |

You don't need 100% to get a job or build real apps. **~70-80% covers most production Go code.**
