package main

import (
    "httpdemo/handlers"
    "httpdemo/models"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    models.InitDB()
    
    api := r.Group("/api/v1")
    {
        api.GET("/users", handlers.GetUsers)
        api.GET("/users/:id", handlers.GetUser)
        api.POST("/users", handlers.CreateUser)
        api.PUT("/users/:id", handlers.UpdateUser)
        api.DELETE("/users/:id", handlers.DeleteUser)
    }
    
    r.Run(":8080")
}