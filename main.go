package main

import (
	"fmt"

	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/database"
	"github.com/MorrisMorrison/retfig/templrender"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Startup retfig")
	database.InitializeDbConnection()
	r := gin.Default()
	r.HTMLRender = &templrender.TemplRender{}
	// r.Use(middleware.HtmxHandler())

	// r.GET("/login", api.GetLogin)
	// r.POST("/login", api.CreateLogin)

	r.GET("/events/:id", api.GetEvent)
	r.POST("/events", api.CreateEvent)

	r.GET("/", api.Index)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
