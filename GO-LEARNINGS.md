# Go Learnings from Task Manager CLI

## Project Structure

```go
package main  // Entry point package

import (
    "bufio"       // Reading user input
    "encoding/json" // JSON file operations
    "fmt"         // Printing (aliased as 'print' in our code)
    "os"          // File operations
    "strconv"     // String to int conversion
    "strings"     // String manipulation
    "time"        // Date/time handling
)
```

---

## Structs (Custom Types)

```go
type Task struct {
    ID          int       `json:"ID"`      // JSON tag for serialization
    Title       string    `json:"Title"`
    AssignedTo  string    `json:"AssignedTo"`
    Status      bool      `json:"Status"`
    DueDate     time.Time `json:"DueDate"`
    TimeCreated time.Time `json:"TimeCreated"`
}
```

**Key points:**
- Struct fields must be capitalized to be exported (visible outside package)
- JSON tags control how fields appear in JSON files
- `time.Time` is Go's built-in date/time type

---

## Slices (Dynamic Arrays)

```go
// Declaration
var TasksSlice = []Task{}

// Append item
TasksSlice = append(TasksSlice, newTask)

// Delete item by index
TasksSlice = append(TasksSlice[:index], TasksSlice[index+1:]...)

// Length
len(TasksSlice)

// Iterate with index (needed for modification)
for i := range TasksSlice {
    TasksSlice[i].Title = "NEW"  // Modifies original
}

// Iterate with copy (read-only)
for _, task := range TasksSlice {
    fmt.Println(task.Title)  // task is a COPY
}
```

**Critical lesson:** Use `for i := range` when you need to **modify** slice elements. Using `for _, item := range` gives you a copy!

---

## Finding Items by ID vs Index

```go
// WRONG - ID is not the same as slice index!
TasksSlice[id].Title = "NEW"  // Bug! ID 5 != index 5

// CORRECT - Find index first
index := -1
for i := range TasksSlice {
    if TasksSlice[i].ID == id {
        index = i
        break
    }
}
if index == -1 {
    fmt.Println("Not found!")
    return
}
TasksSlice[index].Title = "NEW"  // Now it's correct
```

---

## Error Handling

```go
// Functions return error as last value
dueDate, err := time.Parse("Jan 02, 2006", input)
if err != nil {
    fmt.Println("Wrong format!")
    return
}

// Creating errors
import "errors"
return errors.New("item not found")

// Ignoring errors (use sparingly)
data, _ := reader.ReadString('\n')  // _ ignores the error
```

---

## Reading User Input

```go
reader := bufio.NewReader(os.Stdin)

// Read until newline
input, _ := reader.ReadString('\n')

// Clean up the input
input = strings.TrimSpace(input)           // Remove \n and spaces
input = strings.ToUpper(input)             // Uppercase
input = strings.ToUpper(strings.TrimSpace(input))  // Combined
```

---

## String Conversion

```go
// String to Int
num, err := strconv.Atoi("42")

// Int to String
str := strconv.Itoa(42)
```

---

## Switch Statements

```go
// Cleaner than multiple if/else
switch input {
case "1":
    // do something
case "2":
    // do something else
case "C":
    return
default:
    fmt.Println("Invalid!")
}
```

---

## File Operations (JSON)

```go
// Save to file
func saveTask() {
    data, err := json.MarshalIndent(TasksSlice, "", "  ")  // Pretty print
    if err != nil {
        return
    }
    os.WriteFile("Tasks.json", data, 0644)  // 0644 = file permissions
}

// Load from file
func loadTasks() {
    data, err := os.ReadFile("Tasks.json")
    if err != nil {
        return  // File doesn't exist yet
    }
    json.Unmarshal(data, &TasksSlice)  // & passes pointer
}
```

---

## Time/Date Handling

```go
// Parse string to time
dueDate, err := time.Parse("Jan 02, 2006", "Dec 25, 2024")

// Format time to string
dateStr := task.DueDate.Format("Jan 02, 2006")

// Current time
now := time.Now()

// Sleep/delay
time.Sleep(700 * time.Millisecond)
time.Sleep(1 * time.Second)
```

**Note:** Go uses a specific reference date: `Jan 02, 2006 15:04:05` (1/2 3:4:5 2006)

---

## Constants

```go
const (
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Reset  = "\033[0m"
)

// Usage
fmt.Println(Red + "Error!" + Reset)
```

---

## Early Return Pattern

```go
// Instead of nested if/else
func showAllTask() {
    if len(TasksSlice) == 0 {
        fmt.Println("No tasks!")
        return  // Exit early
    }

    // Main logic here - no nesting needed
    for _, task := range TasksSlice {
        fmt.Println(task.Title)
    }
}
```

---

## Common Patterns Summary

| Pattern | Use Case |
|---------|----------|
| `for i := range slice` | Modify slice elements |
| `for _, item := range slice` | Read-only iteration |
| `if err != nil { return }` | Error handling |
| Early return | Avoid deep nesting |
| `switch` | Multiple conditions |
| `strings.TrimSpace()` | Clean user input |
| `strconv.Atoi()` | String to int |

---

## Gotchas to Remember

1. **Capitalized = exported** - `Title` is public, `title` is private
2. **Slices are references** - but range gives copies
3. **ID != Index** - always find the index first
4. **`nil` vs empty** - `var s []int` is nil, `s := []int{}` is empty (both work)
5. **Unused variables = compile error** - use `_` to ignore
