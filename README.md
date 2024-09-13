# ChatGPT Go Client

A simple Go client for interacting with the OpenAI ChatGPT API. This library provides an easy interface for sending and receiving messages to/from OpenAI's GPT models, such as GPT-3.5-turbo or GPT-4.

## Features

- Send messages to OpenAI's GPT models.
- Maintain conversation history (context) between messages.
- Support for different GPT model versions.
- Easy to integrate and customize.

## Installation

To install the library, use the following command:

```bash
go get github.com/acolev/chatgpt-go
```
# Usage
Here is a basic example of how to use the library in your Go project:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"your_project/chatgpt"
)

func main() {
	// Make sure to set your OpenAI API key in environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API key is not set in environment variables")
	}

	// Initialize a new ChatGPT client
	chat := chatgpt.New(apiKey, "gpt-3.5-turbo") // You can use any available model

	// Prepare context (previous conversation history)
	context := []chatgpt.Message{
		{Role: "user", Content: "Hello, who are you?"},
		{Role: "assistant", Content: "I'm an AI created by OpenAI. How can I assist you today?"},
	}

	// Send a new message
	response, err := chat.SendMessage("Can you tell me a joke?", context)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Println("ChatGPT response:", response)
}

```

# Environment Variables
The library requires an OpenAI API key to interact with the OpenAI API. You can set your API key as an environment variable:

```bash
export OPENAI_API_KEY=your_api_key_here
```

# Models
The following GPT models are supported:

* gpt-3.5-turbo
* gpt-4-turbo
You can specify the model when creating the client:

```go
chat := chatgpt.New(apiKey, "gpt-4-turbo")
```
# Handling Context
The library allows you to maintain conversation context by appending previous messages to the current request. Make sure to pass the full context (including both user and assistant messages) in the **SendMessage** function to retain the history.
```go
context := []chatgpt.Message{
    {Role: "user", Content: "Tell me a joke."},
    {Role: "assistant", Content: "Why did the chicken cross the road? To get to the other side!"},
}

response, err := chat.SendMessage("Can you tell me another joke?", context)
```

# Error Handling
All functions return an error if something goes wrong. Always check for errors when calling the API.

Example:

```go
response, err := chat.SendMessage("What is the meaning of life?", context)
if err != nil {
log.Fatalf("Failed to get response: %v", err)
}
```
