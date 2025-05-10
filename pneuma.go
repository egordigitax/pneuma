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

type Pneuma struct {
	opts     pneumaOpts
	provider LLMProvider
}

func Init(key string) *Pneuma {
	return &Pneuma{
		opts:     pneumaOpts{},
		provider: providers.NewOpenAIProvider(key),
	}
}

func InitWithProvider(newProvider LLMProvider) *Pneuma {
	return &Pneuma{
		opts:     pneumaOpts{},
		provider: newProvider,
	}
}

func (p *Pneuma) Fill(s interface{}) error {
	schema, err := jsonschema.GenerateJSONSchema(s)
	if err != nil {
		return err
	}

	resp, err := p.provider.CompleteWithSchema("", schema, jsonschema.GetName(s))

	json.Unmarshal([]byte(resp), s)

	return nil
}

func (p *Pneuma) FillWithContext(s interface{}, context string) error {
	schema, err := jsonschema.GenerateJSONSchema(s)
	if err != nil {
		return err
	}

	resp, err := p.provider.CompleteWithSchema(context, schema, jsonschema.GetName(s))

	json.Unmarshal([]byte(resp), s)

	return nil
}
