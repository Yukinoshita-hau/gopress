# Custom HTTP Router Documentation

This Go-based HTTP router provides an efficient way to define routes, middleware, and static file handling, with a customizable error management system. Below is the detailed documentation on its usage and core components.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Core Components](#core-components)
- [Middleware](#middleware)
- [Error Handling](#error-handling)
- [Examples](#examples)

## Features

- **Route Matching:** Supports dynamic routing with URL parameters.
- **HTTP Methods Support:** GET, POST, PUT, DELETE, PATCH, etc.
- **Middleware:** Attach middleware to individual routes or groups.
- **Static File Serving:** Serve static files easily from a directory.
- **Custom Error Handling:** Manage 404 and 405 errors, or create your own handlers.
- **Route Groups:** Create route groups with shared middleware or path prefixes.

## Installation

```bash
git clone https://github.com/Yukinoshita-hau/gopress
cd gopress
```

## Usage

### Create Router

To use the router, first, initialize it using NewRouter(), then define routes, middlewares, and static file handlers.

```go
func main() {
    router := NewRouter()

    // Define a GET route
    router.Get("/hello", HandlerFunction(func(w Response, r *Request) {
        w.Json(map[string]interface{}{
            "message": "Hello, World!",
        }, http.StatusOK)
    }))

    // Serve static files
    router.Static("/static", "./public")

    // Start the server
    router.ListenAndServe(":8080", router)
}
```

## Core Components

### Router

The `Router` struct is the main object used to define routes and handle requests.

**Methods**:
- `NewRouter() *Router`: Initialize a new router.
- `ServeHTTP(Response, *Request)`: Core method handling requests.
- `Get`, `Post`, `Put`, `Delete`, etc.: Register routes for specific HTTP methods.
- `Static(pathPrefix, directory string)`: Serve static files.

### Route Insertion

Routes are stored in a tree structure (`Tree`). Each node in the tree represents a part of the URL, and dynamic parameters (e.g., `/users/:id`) are supported.

**Example**:

```go
router.Get("/users/:id", HandlerFunction(func(w Response, r *Request) {
    id := r.GetParam("id")
    w.Json(map[string]interface{}{"id": id}, http.StatusOK)
}))
```

### Handler

Handlers process incoming HTTP requests. The library provides a `HandlerFunction` type for convenient inline handler creation.

## Middleware

Middleware functions wrap handlers, allowing pre-processing (like logging, authentication, etc.) before the handler is executed. Middlewares can be applied globally, to individual routes, or to route groups.

### Middleware Definition

```go
func loggingMiddleware(next Handler) Handler {
    return HandlerFunction(func(w Response, r *Request) {
        log.Println("Request received")
        next.ServeHTTP(w, r)
    })
}
```

**Usage**:

```go
router.Get("/secure", HandlerFunction(func(w Response, r *Request) {
    w.Json(map[string]interface{}{"message": "Secure Page"}, http.StatusOK)
}), loggingMiddleware)
```

## Error Handling

### Default Error Responses

The router provides basic 404 (Not Found) and 405 (Method Not Allowed) responses with customizable messages:

```go
var (
    ErrNotFound           = errors.New("Not Found: 404")
    ErrMethodNotAllowed   = errors.New("Method Not Allowed: 405")
    Http404Response       = []byte("Page not found")
    Http405Response       = []byte("Method not allowed")
)
```

### Custom Error Handling

You can define a custom error handler using `SetErrorHandler`:

```go
router.SetErrorHandler(func(w Response, r *Request, err error) {
    status, body := handleErr(err)
    JsonErrorResponse(w, status, string(body))
})
```

## Examples

### Route Grouping

Group routes under a common path with shared middleware:

```go
apiGroup := router.Group("/api", loggingMiddleware)

apiGroup.Get("/users", HandlerFunction(func(w Response, r *Request) {
    w.Json(map[string]interface{}{"users": "List of users"}, http.StatusOK)
}))
```

### Static Files

Serve static files easily from a directory:

```go
router.Static("/assets", "./static")
```

### Handling Parameters

Dynamic URL parameters are accessible using `GetParam()`:

```go
router.Get("/posts/:id", HandlerFunction(func(w Response, r *Request) {
    postID := r.GetParam("id")
    w.Json(map[string]interface{}{"postID": postID}, http.StatusOK)
}))
```

This documentation should give you a clear understanding of how to use this HTTP router, as well as the flexibility it provides in handling routes, middleware, and errors.