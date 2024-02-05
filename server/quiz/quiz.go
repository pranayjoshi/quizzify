package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/franciscoescher/goopenai"
	"github.com/pranayjoshi/quizify/user"
)

type QuizResponse struct {
	Name     string `json:"Name"`
	ID       string `json:"id"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	QuizName string `json:"quiz_name"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomQuizIDGen() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10) // Change 10 to any length you want
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	// url := "https://chatgpt-api.shn.hk/v1/"
	// reqBody := strings.NewReader(`{
	//     "model": "gpt-3.5-turbo",
	//     "messages": [{"role": "user", "content": "Hello, how are you?"}]
	// }`)

	// req, err := http.NewRequest("POST", url, reqBody)
	// if err != nil {
	// 	log.Fatal("Error reading request. ", err)
	// }

	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// fmt.Println(resp)
	// if err != nil {
	// 	log.Fatal("Error reading response. ", err)
	// }
	// defer resp.Body.Close()

	// fmt.Println("response status:", resp.Status)
	apiKey := os.Getenv("CHATGPT_API_KEY")
	organization := os.Getenv("CHATGPT_ORG_ID")
	body, err := io.ReadAll(r.Body)
	fmt.Println(body)
	var response QuizResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	prompt := response.Content
	fmt.Println(prompt)
	client := goopenai.NewClient(apiKey, organization)

	rrt := goopenai.CreateChatCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: response.Content + "\ncreate 10 quiz quiz question based on the text before in a JSON format, with the question, options and the answer in this format type Quiz struct, use  {Question  Option1  Option2   Option3  Option4 Answer }. The answer must be the Option value not the option index.",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateChatCompletions(context.Background(), rrt)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions.Choices[0].Message.Content)
	content := completions.Choices[0].Message.Content
	content = strings.ReplaceAll(content, "/", "")
	content = strings.ReplaceAll(content, "\n", "")
	// jsonData, err := json.Marshal(content)

	quiz := user.Quiz{
		Content:  content,
		Author:   response.Author,
		QuizID:   RandomQuizIDGen(),
		Name:     response.Name,
		QuizName: response.QuizName,
	}

	store := user.NewStore()
	fmt.Println(quiz.Author)
	quizMap := map[string]interface{}{
		"content":   string(quiz.Content),
		"author":    quiz.Author,
		"quiz_id":   quiz.QuizID,
		"name":      quiz.Name,
		"quiz_name": quiz.QuizName,
	}

	erri := store.Update("users/"+response.Name+"/quiz/"+quiz.QuizID, quizMap)
	_ = store.Update("/quiz/"+quiz.QuizID, quizMap)
	if erri != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	quizIDMap := map[string]string{"quiz_id": quiz.QuizID}
	quizIDJSON, err := json.Marshal(quizIDMap)
	if err != nil {
		fmt.Println("Error marshalling JSON: %v", err)
	}

	w.Write(quizIDJSON)
}
