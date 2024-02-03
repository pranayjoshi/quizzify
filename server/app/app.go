package app

import (
	"net/http"
	"os"

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
	r.HandleFunc("/signup", healthHandler).Methods("POST")
}

func (a *App) Start() error {
	a.Router = mux.NewRouter()
	RegisterAPIRoutes(a.Router)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: a.Router,
	}

	return server.ListenAndServe()
}
