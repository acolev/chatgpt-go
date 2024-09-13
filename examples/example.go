package main

import (
	"fmt"
	"github.com/acolev/chatgpt-go/chatgpt"
	"log"
)

func main() {
	apiKey := "your_openai_api_key"
	chat := chatgpt.New(apiKey, "gpt-3.5-turbo") // Можно изменить на любую доступную модель

	// Предыдущие сообщения (контекст)
	context := []chatgpt.Message{
		{Role: "user", Content: "Hello, who are you?"},
		{Role: "assistant", Content: "I'm an AI created by OpenAI. How can I assist you today?"},
	}

	// Отправка нового сообщения с сохранением контекста
	response, err := chat.SendMessage("Can you tell me a joke?", "chat123", context)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ChatGPT response:", response)
}
