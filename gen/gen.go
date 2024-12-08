package gen

import (
	"Code_Gen/edit"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// APIConfig represents the user-defined API configuration
type APIConfig struct {
	ModelSchema []ModelField `json:"modelSchema"`
	BaseURL     string       `json:"baseURL"`
	Endpoints   []Endpoint   `json:"endpoints"`
}

// ModelField defines the structure of a database model
type ModelField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

// Endpoint represents an API endpoint configuration
type Endpoint struct {
	Path      string `json:"path"`
	Method    string `json:"method"`
	ModelName string `json:"modelName"`
	Operation string `json:"operation"` // create, read, update, delete
}

// APIGenerator handles the generation of Go backend code
type APIGenerator struct {
	ProjectPath string
}

// NewAPIGenerator creates a new API generator instance
func NewAPIGenerator(projectPath string) *APIGenerator {
	return &APIGenerator{
		ProjectPath: projectPath,
	}
}

// GenerateModelFile creates a model file based on the schema
func (g *APIGenerator) GenerateModelFile(config APIConfig) (string, error) {
	modelTemplate := `package models

type {{ .ModelName }} struct {
	{{- range .Fields }}
	{{ .Name }} {{ .Type }} {{ .Tags }}
	{{- end }}
}
`
	// Implement model file generation logic
	return modelTemplate, nil
}

// GenerateRoutes creates route handlers based on endpoints
func (g *APIGenerator) GenerateRoutes(config APIConfig) (string, error) {
	routeTemplate := `package routes

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/controllers"
)

func SetupRoutes(r *gin.Engine) {
	{{- range .Endpoints }}
	r.{{ .Method }}("{{ .Path }}", controllers.{{ .ControllerFunc }})
	{{- end }}
}
`
	// Implement route generation logic
	return routeTemplate, nil
}

// GenerateControllers creates controller files for different operations
func (g *APIGenerator) GenerateControllers(config APIConfig) (string, error) {
	controllerTemplate := `package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/models"
	"{{ .ProjectName }}/database"
)

func Create{{ .ModelName }}(c *gin.Context) {
	// Create implementation
}

func Get{{ .ModelName }}(c *gin.Context) {
	// Read implementation
}

func Update{{ .ModelName }}(c *gin.Context) {
	// Update implementation
}

func Delete{{ .ModelName }}(c *gin.Context) {
	// Delete implementation
}
`
	// Implement controller generation logic
	return controllerTemplate, nil
}

// Generate generates the entire backend based on user configuration
func (g *APIGenerator) Generate(config APIConfig) error {
	// Validate input
	if err := g.ValidateConfig(config); err != nil {
		return err
	}

	// Generate files
	model, err := g.GenerateModelFile(config)
	if err != nil {
		return err
	}
	log.Println("Model: ", model)
	route, err := g.GenerateRoutes(config)
	if err != nil {
		return err
	}
	log.Println("Route: ", route)

	controller, err := g.GenerateControllers(config)
	if err != nil {
		return err
	}
	log.Println("Controller: ", controller)
	return nil
}

// validateConfig checks the user-provided configuration
func (g *APIGenerator) ValidateConfig(config APIConfig) error {
	// Add validation logic for API configuration
	if len(config.ModelSchema) == 0 {
		return fmt.Errorf("at least one model field is required")
	}
	return nil
}

func Compile(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request
	var editRequest edit.CodeEditRequest
	err := json.NewDecoder(r.Body).Decode(&editRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform file edit
	err = edit.EditFile(editRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File successfully modified",
	})
}
