package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pranayjoshi/quizify/quiz"
	"github.com/pranayjoshi/quizify/user"

	"github.com/gorilla/mux"
)

// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("OK"))
// }

type App struct {
	Router *mux.Router
}

func RegisterAPIRoutes(r *mux.Router) {
	r.HandleFunc("/login", user.LoginUser).Methods("POST")
	r.HandleFunc("/register", user.RegisterUser).Methods("POST")
	r.HandleFunc("/create_quiz", quiz.CreateQuiz).Methods("POST")
	// r.HandleFunc("/get_quiz:quiz_id", quiz.GetQuiz).Methods("GET")
	// r.HandleFunc("/post_quiz_results:quiz_id", quiz.GetQuiz).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
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

	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed under request %v", err)
		} else {
			log.Printf("Server closed unexpect: %v", err)
		}
	}
	return err
}
