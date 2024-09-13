package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GPTClient представляет клиент для взаимодействия с ChatGPT API
type GPTClient struct {
	apiKey string
	model  string
}

// New создает новый экземпляр GPTClient
func New(apiKey string, version string) *GPTClient {
	if version == "" {
		version = "gpt-3.5-turbo" // Модель по умолчанию
	}
	return &GPTClient{
		apiKey: apiKey,
		model:  version, // Модель по умолчанию, можно вынести в параметр
	}
}

// SendMessage отправляет запрос к ChatGPT API с возможностью использования chat ID
func (client *GPTClient) SendMessage(prompt, chatID string, context []Message) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// Формируем запрос с учетом контекста предыдущих сообщений (если есть)
	messages := append(context, Message{
		Role:    "user",
		Content: prompt,
	})

	reqBody := GPTRequest{
		Model:    client.model,
		Messages: messages,
		ChatID:   chatID, // Поддержка чата с сохранением истории
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiKey))

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
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
