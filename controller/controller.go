package controller

import (
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

// Handle Form is a controller that handles the form submission
func HandleForm(c *gin.Context) {
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
}

// Handle Home Views
func Home(c *gin.Context) {
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
	}

	err = tmpl.Execute(c.Writer, data)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
