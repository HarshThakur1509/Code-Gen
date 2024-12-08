package main

import "Code_Gen/api"


// Example usage
// config := gen.APIConfig{
// 	ModelSchema: []ModelField{
// 		{Name: "ID", Type: "uint", Required: true},
// 		{Name: "Name", Type: "string", Required: true},
// 	},
// 	BaseURL: "/api/v1",
// 	Endpoints: []Endpoint{
// 		{
// 			Path:      "/users",
// 			Method:    "GET",
// 			ModelName: "User",
// 			Operation: "read",
// 		},
// 	},
// }

// generator := gen.NewAPIGenerator("/path/to/project")
// if err := generator.Generate(config); err != nil {
// 	log.Fatal(err)
// }

func main() {
	server := api.NewApiServer(":3000")
	server.Run()
}
