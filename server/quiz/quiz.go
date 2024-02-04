package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/franciscoescher/goopenai"
)

type QuizResponse struct {
	Name   string `json:"Name"`
	ID     string `json:"id"`
	Prompt string `json:"prompt"`
	Author string `json:"author"`
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
	var response QuizResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	prompt := response.Prompt
	fmt.Println(prompt)
	client := goopenai.NewClient(apiKey, organization)

	rrt := goopenai.CreateChatCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "I followed along here to get a general understanding of the algorithm: https://www.techiedelight.com/binary-search/ Things I learned from this: First off, I completely forgot you had to re-evaluate variables even if the dependencies change. For example, although the startIndex and endIndex variables were changing over time dynamically, I had to re-evaluate midIndex to have the correct value. Remember to sort list in ascending order. Binary Search only works on ascending sorted lists. Note that if you have multiples of the same value (say an array of [1, 2, 5, 5]), the algorithm is not particular about which 5 it chooses if that’s your target. Wherever the midIndex lands, it will choose that. So in this example, the first instance of 5 will be chosen in the array (index 2). You can check this when running the code. I saw in the article you can reformat this to be recursive. That made sense to me so I didn’t do it. The scope was to understand how the algorithm worked. I still have to read about how this saves time but I think it intuitively makes sense compared to the linear approach. They explained that linearly, the worst case is that the value is at the end of the list. I can see why that would take a long time if the list was millions of items long." + "\ncreate 10 quiz quiz question based on this arctile in a JSON format, with the question, options and the answer in this format type Quiz struct, use  {Question  Option1  Option2   Option3  Option4 Answer }",
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
	jsonData, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
