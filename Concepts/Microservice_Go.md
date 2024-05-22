# Building a Microservice in Go with Gin

## Table of Contents
1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Project Structure](#project-structure)
4. [Setting Up the Project](#setting-up-the-project)
5. [Creating the Main Application](#creating-the-main-application)
6. [Defining Routes and Handlers](#defining-routes-and-handlers)
7. [Middleware](#middleware)
8. [Configuration](#configuration)
9. [Running the Application](#running-the-application)
10. [Conclusion](#conclusion)

## Introduction
Gin is a high-performance HTTP web framework written in Go (Golang). It's a popular choice for building microservices due to its speed, simplicity, and flexibility.

## Prerequisites
- Go installed (version 1.16+ recommended)
- Basic understanding of Go programming
- Familiarity with RESTful APIs

## Project Structure
A typical Go project using Gin might look like this:
```
my-microservice/
├── main.go
├── go.mod
├── go.sum
├── config/
│   └── config.go
├── controllers/
│   └── user_controller.go
├── models/
│   └── user.go
├── routes/
│   └── routes.go
└── middleware/
    └── auth.go
```

## Setting Up the Project
First, create a new directory for your project and initialize it as a Go module.

```sh
mkdir my-microservice
cd my-microservice
go mod init my-microservice
```

Install Gin:

```sh
go get -u github.com/gin-gonic/gin
```

## Creating the Main Application

Create a `main.go` file:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "my-microservice/routes"
)

func main() {
    r := gin.Default()
    
    // Setup routes
    routes.SetupRoutes(r)
    
    // Start the server
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
```

## Defining Routes and Handlers

Create a `routes/routes.go` file:

```go
package routes

import (
    "github.com/gin-gonic/gin"
    "my-microservice/controllers"
)

func SetupRoutes(r *gin.Engine) {
    // User routes
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
}
```

Create a `controllers/user_controller.go` file:

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "my-microservice/models"
)

// Mock data -- Simulates what would be the data in a Database. *ID could be UUID
var users = []models.User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// GetUsers handles GET /users
func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

// CreateUser handles POST /users
func CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}
```

Create a `models/user.go` file:

```go
package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

## Middleware

Create a `middleware/auth.go` file for example middleware:

```go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// AuthMiddleware is an example middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "mysecrettoken" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        c.Next()
    }
}
```

To use the middleware in your routes, update the `SetupRoutes` function:

```go
package routes

import (
    "github.com/gin-gonic/gin"
    "my-microservice/controllers"
    "my-microservice/middleware"
)

func SetupRoutes(r *gin.Engine) {
    // Apply middleware
    r.Use(middleware.AuthMiddleware())
    
    // User routes
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
}
```

## Configuration

Create a `config/config.go` file for handling configuration:

```go
package config

import (
    "github.com/spf13/viper"
)

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("json")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
}
```

Add a `config.json` file in the root directory:

```json
{
    "port": "8080"
}
```

Update `main.go` to use the configuration:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "my-microservice/config"
    "my-microservice/routes"
)

func main() {
    config.LoadConfig()
    r := gin.Default()
    
    // Setup routes
    routes.SetupRoutes(r)
    
    // Start the server
    r.Run(":" + viper.GetString("port")) // listen and serve on configured port
}
```

## Running the Application

Run the application:

```sh
go run main.go
```
