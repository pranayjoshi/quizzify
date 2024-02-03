package quiz

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Quiz struct {
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Answer   string `json:"answer"`
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	url := "https://chatgpt-api.shn.hk/v1/"
	reqBody := strings.NewReader(`{
        "model": "gpt-3.5-turbo",
        "messages": [{"role": "user", "content": "Hello, how are you?"}]
    }`)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
}
