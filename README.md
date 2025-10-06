# 🧩 Go User API — Learning Project

A simple **REST API built with Go (Golang)** to demonstrate basic web server concepts for **Node.js developers transitioning to Go**.

---

## 🚀 What This Project Demonstrates

- **HTTP Server** — Build REST APIs with Go’s standard library (`net/http`)
- **JSON Handling** — Parse and generate JSON responses
- **Struct Methods** — Object-oriented patterns in Go
- **Error Handling** — Go’s explicit approach to managing errors
- **In-Memory Storage** — Simple persistence with slices

---

## 📋 Prerequisites

- Go **1.19+** ([Download here](https://go.dev/dl/))
- Basic understanding of REST APIs
- Familiarity with JSON

Check installation:

```bash
go version
```

---

## 🏃‍♂️ How to Run

### 1️⃣ Clone or navigate to your project

```bash
cd /path/to/go-course
```

### 2️⃣ Initialize Go module

```bash
go mod init go-course
```

### 3️⃣ Run the application

```bash
go run main.go
```

### 4️⃣ Verify it’s running

```bash
curl http://localhost:8080/users
```

Expected response:
```json
[]
```

Or build and run manually:

```bash
go build -o user-api .
./user-api
```

---

## 📡 API Endpoints

### `GET /users`

Returns all users.

**Example:**
```bash
curl http://localhost:8080/users
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
]
```

---

### `POST /users`

Creates a new user.

**Example:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

**Response:** `201 Created`

**Validation Rules**
- `name` is required  
- `email` is required and must be unique

---

## 🔍 Project Structure

```
.
├── main.go     # Application entry point
├── api.go      # HTTP handlers
└── user.go     # User model
```

---

## 🧠 Go Syntax Deep Dive

### 📦 Package Declaration & Imports
```go
// Every Go file starts with a package declaration
package main  // 'main' package = executable program

// Import statements - bring in external packages
import (
    "fmt"          // Single package
    "encoding/json" // Standard library
    "net/http"     // Another standard library package
)

// Alternative single import
import "fmt"
```

### 🏗️ Variable Declarations
```go
// Explicit type declaration
var name string = "John"
var age int = 25
var isActive bool = true

// Type inference with :=
name := "John"          // string inferred
age := 25              // int inferred
isActive := true       // bool inferred

// Multiple variable declaration
var (
    name     string = "John"
    age      int    = 25
    isActive bool   = true
)

// Multiple assignment
a, b := 1, "hello"
name, email := "John", "john@example.com"

// Zero values (default values when not initialized)
var count int     // 0
var message string // ""
var flag bool     // false
var ptr *int      // nil
```

### 🔢 Basic Types
```go
// Integers
var age int = 25           // Platform dependent (32 or 64 bit)
var count int32 = 100      // 32-bit integer
var bigNum int64 = 999999  // 64-bit integer

// Floating point
var price float32 = 99.99
var precise float64 = 3.14159265359

// Strings
var name string = "Go Developer"
var multiline string = `This is a
multi-line string
using backticks`

// Booleans
var isReady bool = true

// Byte and Rune
var letter byte = 'A'        // byte = uint8
var unicode rune = '🚀'      // rune = int32 (Unicode code point)
```

### 🏛️ Structs (Go's "Classes")
```go
// Define a struct
type User struct {
    ID    int    `json:"id"`    // Field with JSON tag
    Name  string `json:"name"`  // Public field (capitalized)
    email string               // Private field (lowercase)
}

// Create struct instances
user1 := User{ID: 1, Name: "John", email: "john@example.com"}
user2 := User{
    ID:   2,
    Name: "Jane",
    email: "jane@example.com",
} // Trailing comma allowed

// Zero value struct
var user3 User // ID: 0, Name: "", email: ""

// Access fields
fmt.Println(user1.Name)  // "John"
user1.Name = "Johnny"    // Modify field
```

### 🎯 Methods (Functions on Structs)
```go
// Method with value receiver (read-only)
func (u User) GetFullInfo() string {
    return fmt.Sprintf("ID: %d, Name: %s", u.ID, u.Name)
}

// Method with pointer receiver (can modify)
func (u *User) SetEmail(email string) {
    u.email = email  // Modifies the original struct
}

// Method usage
user := User{ID: 1, Name: "John"}
info := user.GetFullInfo()     // Call method
user.SetEmail("john@new.com")  // Modify via pointer method
```

### 🚨 Error Handling Pattern
```go
// Functions return multiple values (result, error)
func divideNumbers(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil  // nil means no error
}

// Error handling pattern
result, err := divideNumbers(10, 2)
if err != nil {
    // Handle error
    fmt.Println("Error:", err)
    return
}
// Use result
fmt.Println("Result:", result)

// Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

### 📊 Slices (Dynamic Arrays)
```go
// Declare empty slice
var users []User

// Slice literal
numbers := []int{1, 2, 3, 4, 5}
names := []string{"Alice", "Bob", "Charlie"}

// Add elements with append
users = append(users, User{ID: 1, Name: "John"})
numbers = append(numbers, 6, 7, 8)

// Slice operations
fmt.Println(len(numbers))     // Length: 8
fmt.Println(numbers[0])       // First element: 1
fmt.Println(numbers[1:3])     // Slice [2, 3]
fmt.Println(numbers[:2])      // First 2 elements [1, 2]
fmt.Println(numbers[3:])      // From index 3 to end

// Iterate over slice
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Iterate without index
for _, value := range numbers {
    fmt.Println(value)
}
```

### 🗺️ Maps (Key-Value Pairs)
```go
// Declare and initialize map
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Alternative declaration
var scores map[string]int
scores = make(map[string]int)

// Add/update values
ages["David"] = 28
scores["math"] = 95

// Check if key exists
age, exists := ages["Alice"]
if exists {
    fmt.Printf("Alice is %d years old\n", age)
}

// Delete key
delete(ages, "Bob")

// Iterate over map
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### 🔗 Pointers
```go
// Declare pointer
var ptr *int

// Get address of variable
x := 42
ptr = &x  // ptr points to x's memory address

// Dereference pointer (get value)
fmt.Println(*ptr)  // Prints: 42

// Modify value through pointer
*ptr = 100
fmt.Println(x)     // Prints: 100

// Pointers with structs
user := User{ID: 1, Name: "John"}
userPtr := &user

// Access struct fields through pointer (automatic dereferencing)
fmt.Println(userPtr.Name)  // Same as (*userPtr).Name
userPtr.Name = "Johnny"    // Modifies original struct
```

### 🔄 Control Structures

#### If Statements
```go
// Basic if
if age >= 18 {
    fmt.Println("Adult")
}

// If-else
if score >= 90 {
    fmt.Println("A grade")
} else if score >= 80 {
    fmt.Println("B grade")
} else {
    fmt.Println("Need improvement")
}

// If with initialization
if user, err := getUser(id); err != nil {
    // Handle error
} else {
    // Use user
    fmt.Println(user.Name)
}
```

#### For Loops
```go
// Classic for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    // Do something forever
    if condition {
        break
    }
}

// Range loop
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %v\n", index, value)
}
```

#### Switch Statements
```go
// Basic switch
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}

// Switch with no condition (like if-else chain)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
default:
    grade = "C"
}
```

### 🔧 Functions
```go
// Basic function
func greet(name string) string {
    return "Hello, " + name
}

// Multiple parameters and return values
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return // Named returns
}

// Function with error
func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return errors.New("invalid email format")
    }
    return nil
}

// Function as variable
var operation func(int, int) int
operation = func(a, b int) int {
    return a + b
}
result := operation(5, 3)  // 8

// Anonymous function (closure)
func() {
    fmt.Println("This runs immediately")
}()
```

### 🏷️ Interfaces
```go
// Define interface
type Speaker interface {
    Speak() string
    SetVolume(int)
}

// Implement interface (implicit)
type Dog struct {
    Name string
    volume int
}

func (d Dog) Speak() string {
    return "Woof! I'm " + d.Name
}

func (d *Dog) SetVolume(v int) {
    d.volume = v
}

// Use interface
func makeNoise(s Speaker) {
    fmt.Println(s.Speak())
}

dog := Dog{Name: "Buddy"}
makeNoise(dog)  // Works because Dog implements Speaker
```

### 📄 JSON Handling
```go
// Struct with JSON tags
type Person struct {
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Email    string `json:"email,omitempty"` // Omit if empty
    Password string `json:"-"`               // Never serialize
}

// Marshal (Go struct → JSON)
person := Person{Name: "John", Age: 30, Email: "john@example.com"}
jsonData, err := json.Marshal(person)
if err != nil {
    // Handle error
}
fmt.Println(string(jsonData))

// Unmarshal (JSON → Go struct)
jsonString := `{"name":"Jane","age":25}`
var person2 Person
err = json.Unmarshal([]byte(jsonString), &person2)
if err != nil {
    // Handle error
}
```

---

## 🎯 Next Steps

- Add database integration (PostgreSQL, MongoDB)
- Implement authentication and authorization
- Add logging and validation
- Write unit tests
- Containerize with Docker

---

## 📖 Learning Resources

- [Official Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Web Examples](https://gowebexamples.com/)

---

**Happy Go Coding! 🎉**
