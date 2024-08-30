
# Custom HTTP Router in Go

This project is a simple HTTP router implemented in Go. It allows you to define routes, attach middleware, serve static files, and handle errors in a customizable way.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Documentation](#documentation)

## Features

- **Route Matching**: Supports dynamic route matching with parameters.
- **HTTP Methods**: Supports all standard HTTP methods (GET, POST, PUT, DELETE, etc.).
- **Middleware Support**: Allows attaching middleware to routes and route groups.
- **Static File Serving**: Supports serving static files from a directory.
- **Error Handling**: Customizable error handling.
- **Route Groups**: Supports grouping routes with common prefixes and middlewares.

## Installation

To install this package, clone the repository and build it using Go:

```bash
git clone <repository-url>
cd <repository-directory>
go build
```

## Usage

Below is a basic example of how to use the custom router:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Create a new router
    app := NewRouter()

    // Define a simple route
    app.Get("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    }))

    // Serve static files from the "./public" directory
    app.Static("/static", "./public")

    // Start the HTTP server
    fmt.Println("Server starting at :8080")
    http.ListenAndServe(":8080", app)
}
```

## Examples

### Defining Routes

You can define routes using HTTP methods:

```go
app.Get("/users/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Handle GET request for a user with a specific ID
}))
```

### Using Middleware

Middleware can be added to routes to perform actions like logging or authentication:

```go
loggingMiddleware := func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request received")
        next.ServeHTTP(w, r)
    })
}

app.Get("/secure", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Secure Content")
}), loggingMiddleware)
```

### Route Groups

Route groups allow you to apply common middleware or prefixes to a set of routes:

```go
apiGroup := app.Group("/api", loggingMiddleware)
apiGroup.Get("/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "User List")
}))
```

## Documentation

### Core Components

- **Router**: The main entry point to define routes and groups.
- **Middleware**: Functions that wrap HTTP handlers to provide additional functionality.
- **Tree**: Represents the routing tree structure where all routes are stored.
- **RouterGroup**: Represents a group of routes with a common prefix and middleware.

### Key Methods

- `NewRouter() *Router`: Creates a new router instance.
- `Router.Get(path string, handler http.Handler, middleware ...Middleware)`: Registers a new GET route.
- `Router.Post(path string, handler http.Handler, middleware ...Middleware)`: Registers a new POST route.
- `Router.Static(pathPrefix, directory string)`: Serves static files from the given directory.
- `Router.ServeHTTP(w http.ResponseWriter, req *http.Request)`: Handles incoming HTTP requests.
- `Router.Group(prefix string, middlewares ...Middleware) *RouterGroup`: Creates a new route group.

### Error Handling

The router provides a default error handler but allows you to define custom error handlers:

```go
app.SetErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
    // Custom error handling logic
})
```