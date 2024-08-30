# Custom HTTP Router in Go

This project is a simple HTTP router implemented in Go. It allows you to define routes, attach middleware, serve static files, and handle errors in a customizable way.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Documentation](#documentation)
- [License](#license)

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