# 🚂 Pneuma

**Pneuma** is a lightweight Go package for filling structs using [LLMs](https://platform.openai.com/) like OpenAI, guided by type definitions and struct field tags. It auto-generates JSON Schemas from your struct and prompts the model to produce realistic data.

---

## ✨ Features

- 🧩 Struct filling via OpenAI LLM  
- 🧠 Automatic JSON Schema generation  
- 🏷️ Prompt hints via struct tags  
- 🔌 Custom LLM provider support  
- 🧪 Great for prototyping, testing, mocking, and generative UIs  
- 🌐 **New:** Context-aware struct filling (`FillWithContext`)

---

## 📦 Installation

~~~bash
go get github.com/egordigitax/pneuma
~~~

---

## ⚡ Quick Start

Here's a minimal usage example:

~~~go
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
        Name          string `pneuma:"use english language"`
        Age           int    `pneuma:"more than 6"`
        FavouriteFood string `pneuma:"prefers fruits"`
    }

    d := Dog{}
    err := p.Fill(&d)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Dog: %+v\n", d)
}
~~~

---

## 🛠️ API

### Initialize

Creates a Pneuma client with an OpenAI API key:

~~~go
p := pneuma.Init("your-openai-key")
~~~

Or, if you want to use a custom LLM provider:

~~~go
p := pneuma.InitWithProvider(func() LLMProvider {
    return yourCustomProvider{}
})
~~~

Where `LLMProvider` is:

~~~go
type LLMProvider interface {
    CompleteWithSchema(prompt string, schema json.RawMessage, objectName string) (json.RawMessage, error)
}
~~~

### Fill Struct

~~~go
err := p.Fill(&yourStruct)
~~~

- Inspects `yourStruct` with reflection  
- Generates a JSON Schema  
- Submits it to the OpenAI (or custom) LLM  
- Populates `yourStruct` with generated data  
- Field types supported: `string`, `int`, `float`, `bool`, `struct`, `slice`  
- Tags: Use `pneuma:"some hint"` to guide value generation  

### Fill Struct With Context

In some cases, you may want to provide additional text context or instructions to influence how data is filled. Use `FillWithContext`:

~~~go
err := p.FillWithContext(&yourStruct, "some extra context or instructions here")
~~~

- Works like `Fill` but includes the extra context when prompting the LLM  
- Perfect for classification or more targeted generation  

#### Example

Below is a more advanced example showcasing `FillWithContext` for message moderation:

~~~go
package main

import (
    "fmt"
    "os"

    "github.com/egordigitax/pneuma"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env") // ensure OPENAI_KEY is in your .env
    p := pneuma.Init(os.Getenv("OPENAI_KEY"))

    type ModeratedMessage struct {
        OffensiveLevel  int    `pneuma:"0-10 — how offensive the content is"`
        SpamProbability int    `pneuma:"0-10 — how likely the content is spam"`
        NudityLevel     int    `pneuma:"0-10 — how explicit the content is"`
        Language        string `pneuma:"the language used in the content, e.g. 'english', 'spanish'"`
        Tone            string `pneuma:"general tone, e.g. 'aggressive', 'neutral', 'friendly'"`
    }

    prompt := `
- You are analyzing a user-submitted message to determine its moderation levels.
  Evaluate how offensive, spammy, or explicit the content is. Identify the predominant language 
  and determine the general tone
`
    // First example
    m := ModeratedMessage{}
    message := "Buy cheap products now!!!"
    err := p.FillWithContext(&m, message + prompt)
    if err != nil {
        panic(err)
    }

    fmt.Printf("First ModeratedMessage: %+v\n", m)

    // Second example
    m2 := ModeratedMessage{}
    message2 := "I hate you so much."
    err = p.FillWithContext(&m2, message2 + prompt)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Second ModeratedMessage: %+v\n", m2)
}
~~~

---

## 🧪 Use Cases

- Auto-generating mock/test data  
- LLM-driven prototyping  
- Pre-filled forms / UX mocking  
- Dynamic content creation  
- AI-powered code generation scaffolds  
- **New:** Passing contextual instructions for more precise data generation (e.g. classification)  

---

## 📂 Example Struct Hints

~~~go
type Book struct {
    Title     string `pneuma:"classic english literature"`
    Pages     int    `pneuma:"more than 200"`
    Published bool   `pneuma:"likely true"`
}
~~~

You can combine multiple hints:

~~~go
type Post struct {
    Title   string `pneuma:"funny, short"`
    Content string `pneuma:"light-hearted, about daily life"`
}
~~~

---

## 🧱 Internals

- Uses reflection + `jsonschema` to describe the struct  
- Embeds instructions via struct tags into the prompt  
- Delegates generation to OpenAI (or your custom LLM provider) via the completion endpoint  

---

## 📄 License

MIT — feel free to use, fork, extend. 
