package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Определяем структуру запроса к API
type GPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Определяем структуру ответа API
type GPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// Функция отправки запроса к API
func SendMessage(apiKey, prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	reqBody := GPTRequest{
		Model: "gpt-3.5-turbo", // Используемая модель
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gptResp GPTResponse
	err = json.Unmarshal(body, &gptResp)
	if err != nil {
		return "", err
	}

	if len(gptResp.Choices) > 0 {
		return gptResp.Choices[0].Message.Content, nil
	}

	return "", nil
}
