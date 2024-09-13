package chatgpt

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	version := "gpt-3.5-turbo"
	client := New(apiKey, version)

	if client.apiKey != apiKey {
		t.Errorf("expected apiKey %s, got %s", apiKey, client.apiKey)
	}

	if client.model != version {
		t.Errorf("expected model %s, got %s", version, client.model)
	}
}

func TestNewClient_DefaultModel(t *testing.T) {
	apiKey := "test-api-key"
	client := New(apiKey, "")

	if client.model != "gpt-3.5-turbo" {
		t.Errorf("expected default model gpt-3.5-turbo, got %s", client.model)
	}
}

func TestSendMessage(t *testing.T) {
	// Мокаем API запросы с помощью mock сервера или пишем отдельные тесты для реального API
	client := New("test-api-key", "gpt-3.5-turbo")

	// Мокаем контекст сообщений
	context := []Message{
		{Role: "user", Content: "Hello"},
		{Role: "assistant", Content: "Hi, how can I help you?"},
	}

	// Пример запроса
	response, err := client.SendMessage("Can you tell me a joke?", "test-chat-id", context)
	if err != nil {
		t.Errorf("error while sending message: %v", err)
	}

	// Проверяем, что ответ не пуст
	if response == "" {
		t.Errorf("expected non-empty response, got %s", response)
	}
}
