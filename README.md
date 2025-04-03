# 🚂 Pneuma

**Pneuma** is a lightweight Go package for filling structs using [LLMs](https://platform.openai.com/) like OpenAI, guided by type definitions and struct field tags. It auto-generates JSON Schemas from your struct and prompts the model to produce realistic data.

---

## ✨ Features

- 🧩 Struct filling via OpenAI LLM
- 🧠 Automatic JSON Schema generation
- 🏷️ Prompt hints via struct tags
- 🔌 Custom LLM provider support
- 🧪 Great for prototyping, testing, mocking, and generative UIs

---

## 📦 Installation

```bash
go get github.com/egordigitax/pneuma
```

---

## ⚡ Quick Start

```go
package main

import (
	"fmt"
	"os"

	"github.com/egordigitax/pneuma"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env") // make sure OPENAI_KEY is in your .env
	p := pneuma.Init(os.Getenv("OPENAI_KEY"))

	type Dog struct {
		Name          string `pneuma:"use russian language"`
		Age           int    `pneuma:"more than 6"`
		FavouriteFood string `pneuma:"more like fruits"`
	}

	d := Dog{}
	p.Fill(&d)

	fmt.Printf("%+v\n", d)
}
```

---

## 🛠️ API

### Initialize

```go
p := pneuma.Init("your-openai-key")
```

### Fill Struct

```go
err := p.Fill(&yourStruct)
```

- Field types supported: `string`, `int`, `float`, `bool`, `struct`, `slice`
- Tags: Use `pneuma:"some hint"` to guide value generation

### Custom LLM Provider (Optional)

You can use a different LLM backend by implementing:

```go
type LLMProvider interface {
	CompleteWithSchema(prompt string, schema json.RawMessage) json.RawMessage
}
```

And initializing:

```go
p := pneuma.InitWithProvider(func() LLMProvider {
	return yourCustomProvider{}
})
```

---

## 🧪 Use Cases

- Auto-generating mock/test data
- LLM-driven prototyping
- Pre-filled forms / UX mocking
- Dynamic content creation
- AI-powered code generation scaffolds

---

## 📂 Example Struct Hints

```go
type Book struct {
	Title     string `pneuma:"classic english literature"`
	Pages     int    `pneuma:"more than 200"`
	Published bool   `pneuma:"likely true"`
}
```

---

## 🧱 Internals

- Uses reflection + `jsonschema` to describe the struct
- Embeds instructions via struct tags into the prompt
- Delegates generation to OpenAI via completion endpoint

---

## 📄 License

MIT — feel free to use, fork, extend.

---

## 🔮 Coming Soon

- Support for more LLM providers
- CLI for batch generation
- Custom schema hooks
