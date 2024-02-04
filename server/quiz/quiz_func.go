package quiz

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pranayjoshi/quizify/user"
)

func getQuiz(w http.ResponseWriter, r *http.Request) {
	store := user.NewStore()
	quiz, _ := store.GetByQuiz("quiz/" + r.URL.Query().Get("quiz_id"))

	fmt.Println(quiz.Author)
	quizMap := map[string]interface{}{
		"content":   string(quiz.Content),
		"author":    quiz.Author,
		"cuizID":    quiz.QuizID,
		"name":      quiz.Name,
		"quiz_name": quiz.QuizName,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(quizMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
