package handler

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (h *DocsHandler) ServeSwaggerUI(c *gin.Context) {
	tmpl, err := template.ParseFiles(filepath.Join("internal", "templates", "swagger.html"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load documentation"})
		return
	}

	c.Header("Content-Type", "text/html")
	tmpl.Execute(c.Writer, nil)
}

func (h *DocsHandler) ServeSwaggerYAML(c *gin.Context) {
	yamlContent, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read swagger.yaml"})
		return
	}

	c.Header("Content-Type", "text/yaml")
	c.Writer.Write(yamlContent)
}
