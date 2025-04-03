package pneuma

import (
	"encoding/json"

	"github.com/egordigitax/pneuma/internal/jsonschema"
	"github.com/egordigitax/pneuma/internal/llm"
)

type LLMProvider interface {
	CompleteWithSchema(prompt string, schema json.RawMessage) json.RawMessage
}

type pneumaOpts struct {
	key string
}

type pneuma struct {
	opts     pneumaOpts
	provider LLMProvider
}

func Init(key string) *pneuma {
	return &pneuma{
		opts:     pneumaOpts{},
		provider: llm.NewOpenAIProvider(key),
	}
}

func InitWithProvider(newProvider func() LLMProvider) *pneuma {
	return &pneuma{
		opts:     pneumaOpts{},
		provider: newProvider(),
	}
}

func (p *pneuma) Fill(s interface{}) error {
	schema, err := jsonschema.GenerateJSONSchema(s)
	if err != nil {
		return err
	}

	resp := llm.InitLLM(p.opts.OpenAIKey, schema, jsonschema.GetName(s))
	json.Unmarshal([]byte(resp), s)

	return nil
}
