# ChatGPT Go Client

This package provides a simple client for interacting with OpenAI's ChatGPT API in Go.

## Installation


## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/acolev/chatgpt-go/chatgpt"
)

func main() {
    apiKey := "your_openai_api_key"
    prompt := "Tell me a joke."

    response, err := chatgpt.SendMessage(apiKey, prompt)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("ChatGPT response:", response)
}
```

