package quiz

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pranayjoshi/quizify/user"
)

var store = user.NewStore()

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	quizID := r.URL.Query().Get("quiz_id")
	if quizID == "" {
		http.Error(w, "Missing quiz_id", http.StatusBadRequest)
		return
	}

	quiz, _ := store.GetByQuiz(quizID)

	fmt.Println(quiz.Author)
	quizMap := map[string]interface{}{
		"content":   string(quiz.Content),
		"author":    quiz.Author,
		"quiz_id":   quiz.QuizID,
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

type Results struct {
	Name        string `json:"name"`
	QuizName    string `json:"quiz_name"`
	QuizID      string `json:"quiz_id"`
	ScoredMarks int    `json:"scored_marks"`
	TotalMarks  int    `json:"total_marks"`
	Author      string `json:"author"`
}

func PostResults(w http.ResponseWriter, r *http.Request) {
	var results Results
	err := json.NewDecoder(r.Body).Decode(&results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert results to map[string]interface{}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resultsMap map[string]interface{}
	err = json.Unmarshal(resultsJSON, &resultsMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.Update("users/"+results.Name+"/results/"+results.QuizID, resultsMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
