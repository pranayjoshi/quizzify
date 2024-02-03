package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type App struct {
	Router *mux.Router
}

func RegisterAPIRoutes(r *mux.Router) {
	r.HandleFunc("/login", healthHandler).Methods("GET")
	r.HandleFunc("/register", healthHandler).Methods("POST")

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (a *App) Start() error {
	a.Router = mux.NewRouter()
	RegisterAPIRoutes(a.Router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Starting server on port " + port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.Router,
	}

	return server.ListenAndServe()
}
