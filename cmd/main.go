package main

import (
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve the form
	r.GET("/", func(c *gin.Context) {
		var filepath = path.Join("views", "form.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		err = tmpl.Execute(c.Writer, nil)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})

	// Handle the form submission
	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")

		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var data = map[string]interface{}{
			"title": "EAI",
			"name":  name,
			"date":  time.Now().Format("02 January 2006"),
			"time":  time.Now().Format("15:04:05"),
		}

		err = tmpl.Execute(c.Writer, data)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})

	// Serve static files
	r.Static("/static", "./assets")

	// Start the server
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
