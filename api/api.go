package api

import (
	"Code_Gen/download"
	"Code_Gen/gen"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /api/download", download.Download)
	router.HandleFunc("POST /api/upload", download.Upload)
	router.HandleFunc("GET /api/compile", gen.Compile)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Specify your React frontend origin
		AllowCredentials: true,                              // Allow cookies and credentials
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Content-Disposition", "application/zip", "application/json", "Authorization"},
	}).Handler(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: corsHandler,
	}
	fmt.Println("Server has started", s.addr)
	return server.ListenAndServe()
}
