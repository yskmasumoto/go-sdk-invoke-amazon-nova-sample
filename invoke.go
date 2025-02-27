package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

var prompt = "What is the capital of France?"

func buildRequestBody() BedrockRequest {
	reqPayload := BedrockRequest{
		Messages: []Messages{
			{
				Role: "user",
				Content: []map[string]string{
					{
						"text": prompt,
					},
				},
			},
		},
		InferenceConfig: InferenceConfig{
			MaxTokens:   100,
			Temperature: 0.1,
		},
	}
	return reqPayload
}

func InvokeBedrock(ctx context.Context) (string, error) {
	// new config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		return "", err
	}

	// create bedrock client
	client := bedrockruntime.NewFromConfig(cfg)

	// build request body
	reqPayload := buildRequestBody()
	slog.Debug("invoke bedrock request struct", "request", reqPayload)

	// marshal request body
	body, err := json.Marshal(reqPayload)
	if err != nil {
		slog.Error("failed to marshal request body", "error", err)
		return "", err
	}

	// invoke model
	output, err := client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String("amazon.nova-pro-v1:0"), // Amazon Nova ProモデルのIDを環境変数から取得
		ContentType: aws.String("application/json"),
		Body:        body,
	})
	if err != nil {
		slog.Error("failed to invoke model", "error", err)
		return "", err
	}

	// output.Body([]byte) convert to struct
	var responseJson BedrockResponse
	err = json.Unmarshal(output.Body, &responseJson)
	if err != nil {
		slog.Error("failed to marshal response body", "error", err)
		return "", err
	}

	slog.Debug("invoke bedrock response struct", "response", responseJson)
	return responseJson.Output.Message.Content[0]["text"], nil
}
