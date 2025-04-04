package providers

import (
	"context"
	"encoding/json"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/shared"
)

type OpenAIProvider struct {
	key string
}

func NewOpenAIProvider(key string) *OpenAIProvider {
	return &OpenAIProvider{
		key: key,
	}
}

func (o *OpenAIProvider) CompleteWithSchema(
	prompt string,
	schema json.RawMessage,
	objectName string,
) (json.RawMessage, error) {

	promptWithContext := "fill with random data this schema"
	if prompt != "" {
		promptWithContext = promptWithContext + ", context: " + prompt
	}

	client := openai.NewClient(
		option.WithAPIKey(o.key),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(promptWithContext),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &shared.ResponseFormatJSONSchemaParam{
				JSONSchema: shared.ResponseFormatJSONSchemaJSONSchemaParam{
					Name:   objectName,
					Strict: param.NewOpt(true),
					Schema: schema,
				},
			},
		},
		Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		return nil, err
	}

	return json.RawMessage(chatCompletion.Choices[0].Message.Content), nil
}
