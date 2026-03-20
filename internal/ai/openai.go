package ai

import (
	"context"
	"errors"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAI struct {
	client *openai.Client
	model  string
}

func NewOpenAI(apiKey string, model string) *OpenAI {

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAI{
		client: &client,
		model:  model,
	}
}

func (o *OpenAI) ImproveCommitMessage(
	ctx context.Context,
	prompt string,
) (string, error) {

	if prompt == "" {
		return "", errors.New("prompt cannot be empty")
	}

	resp, err := o.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: o.model,

			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(
					"You are an assistant that improves git commit messages following Conventional Commits.",
				),

				openai.UserMessage(prompt),
			},

			Temperature: openai.Float(0.3),
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}

	content := resp.Choices[0].Message.Content

	if content == "" {
		return "", errors.New("empty response from OpenAI")
	}

	return content, nil
}
