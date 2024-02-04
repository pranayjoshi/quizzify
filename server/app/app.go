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
	"github.com/rs/cors"
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
	r.HandleFunc("/get_quiz", quiz.GetQuiz).Methods("GET")
	r.HandleFunc("/post_results", quiz.PostResults).Methods("POST")
	r.HandleFunc("/get_results", user.GetResultsbyName).Methods("GET")
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // All headers
		AllowCredentials: true,
	})

	handler := c.Handler(a.Router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Starting server on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, handler))

	return nil
}
