Here's a more polished and structured version of your README:

---

# 🌬️ Pneuma – Intelligent Struct Filling with LLMs

**Pneuma** is a lightweight Go package that automatically populates structs using large language models (LLMs) like OpenAI. By leveraging type definitions and struct field tags, it generates realistic, context-aware data – perfect for prototyping, testing, and dynamic applications.

[![Go Reference](https://pkg.go.dev/badge/github.com/egordigitax/pneuma.svg)](https://pkg.go.dev/github.com/egordigitax/pneuma)

## 🔥 Features

- **Smart Struct Population** – Automatically fills structs using LLM intelligence
- **Type-Driven Generation** – Infers requirements from your Go types and tags
- **Context-Aware** – Optional contextual instructions for precise generation
- **Extensible Architecture** – Bring your own LLM provider
- **JSON Schema Under the Hood** – Converts structs to schemas for reliable generation
- **Rich Tag Support** – Guide generation with field-specific hints

Perfect for:  
✅ Mock data generation  
✅ AI-powered prototyping  
✅ Dynamic form filling  
✅ Contextual classification

## 📦 Installation

```bash
go get github.com/egordigitax/pneuma
```

## 🚀 Quick Start

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/egordigitax/pneuma"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load OPENAI_KEY from .env
	p := pneuma.Init(os.Getenv("OPENAI_KEY"))
	
	type Product struct {
		Name        string  `pneuma:"tech gadget"`
		Price       float64 `pneuma:"50-200 range"`
		InStock     bool    `pneuma:"likely true"`
		Description string  `pneuma:"concise marketing copy"`
	}
	
	var prod Product
	if err := p.Fill(&prod); err != nil {
		panic(err)
	}
	
	fmt.Printf("Generated product: %+v\n", prod)
}
```

## 🧩 Advanced Usage

### Contextual Generation

Provide additional context for more precise results:

```go
type Review struct {
	Rating    int    `pneuma:"1-5 scale"`
	Sentiment string `pneuma:"positive/neutral/negative"`
	Summary   string
}

func analyzeText(text string) {
	var review Review
	ctx := "Analyze this product review:\n" + text
	p.FillWithContext(&review, ctx)
	
	fmt.Printf("Analysis: %+v\n", review)
}
```

### Custom LLM Providers

Implement your own provider:

```go
type MyProvider struct{}

func (p MyProvider) CompleteWithSchema(
	prompt string, 
	schema json.RawMessage,
	objectName string,
) (json.RawMessage, error) {
	// Your custom logic
}

p := pneuma.InitWithProvider(func() pneuma.LLMProvider {
	return MyProvider{}
})
```

## 🏷️ Struct Tag Guide

Add generation hints via `pneuma` tags:

```go
type UserProfile struct {
	Username   string `pneuma:"tech-savvy username"`
	Age        int    `pneuma:"18-35 range"`
	Interests  []string
	Bio        string `pneuma:"short professional bio"`
	IsVerified bool   `pneuma:"20% chance"`
}
```

## 📚 Examples

### Mock API Response

```go
type APIResponse struct {
	UserID    string `json:"user_id"`
	Status    string `pneuma:"success/error"`
	Data      map[string]interface{}
	Timestamp int64
}

var mockResponse APIResponse
p.Fill(&mockResponse)
```

### Classification

```go
type TweetAnalysis struct {
	Topic      string   `pneuma:"main topic"`
	Hashtags   []string `pneuma:"relevant hashtags"`
	Sentiment  float64  `pneuma:"0-1 positivity score"`
	IsPolitical bool
}

p.FillWithContext(&analysis, tweetText)
```

## 🤖 How It Works

1. **Schema Generation**: Converts your struct into JSON Schema
2. **Prompt Construction**: Combines schema with your field hints
3. **LLM Completion**: Sends to OpenAI (or your custom provider)
4. **Result Mapping**: Hydrates your struct with the response

## 📜 License

MIT © [Your Name]

---

This version:
- Uses clearer section headers
- Better visual hierarchy
- More practical examples
- Improved feature descriptions
- Consistent formatting
- Added Go pkg.dev badge
- Better organization of advanced features

Would you like me to adjust any particular section further?