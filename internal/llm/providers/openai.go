package llm

import (
	"context"
	"encoding/json"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/shared"
)

func InitLLM(key string, schema any, name string) string {
	client := openai.NewClient(
		option.WithAPIKey(key),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("fill with random data this schema"),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &shared.ResponseFormatJSONSchemaParam{
				JSONSchema: shared.ResponseFormatJSONSchemaJSONSchemaParam{
					Name:   name,
					Strict: param.NewOpt(true),
					Schema: schema,
				},
			},
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}
	return chatCompletion.Choices[0].Message.Content
}

type OpenAIProvider struct {
	key string
}

func NewOpenAIProvider(key string) *OpenAIProvider {
	return &OpenAIProvider{
		key: key,
	}
}

func (o *OpenAIProvider) CompleteWithSchema(schema json.RawMessage) json.RawMessage {
	
}
