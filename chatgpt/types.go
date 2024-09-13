package chatgpt

// GPTRequest представляет структуру запроса к ChatGPT API
type GPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message представляет структуру сообщения
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GPTResponse представляет структуру ответа от ChatGPT API
type GPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}
