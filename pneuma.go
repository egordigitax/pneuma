package pneuma

import (
	"encoding/json"

	"github.com/egordigitax/pneuma/internal/jsonschema"
	"github.com/egordigitax/pneuma/internal/llm/providers"
)

type LLMProvider interface {
	CompleteWithSchema(
		prompt string,
		schema json.RawMessage,
		objectName string,
	) (json.RawMessage, error)
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
		provider: providers.NewOpenAIProvider(key),
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

	resp, err := p.provider.CompleteWithSchema("", schema, jsonschema.GetName(s))

	json.Unmarshal([]byte(resp), s)

	return nil
}

func (p *pneuma) FillWithContext(s interface{}, context string) error {
	schema, err := jsonschema.GenerateJSONSchema(s)
	if err != nil {
		return err
	}

	resp, err := p.provider.CompleteWithSchema(context, schema, jsonschema.GetName(s))

	json.Unmarshal([]byte(resp), s)

	return nil
}
