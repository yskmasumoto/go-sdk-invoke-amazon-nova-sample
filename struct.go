package main

// BedrockRequest represents the request payload for Amazon Nova Pro.
type BedrockRequest struct {
	Messages        []Messages      `json:"messages"`
	InferenceConfig InferenceConfig `json:"inferenceConfig"`
}

type InferenceConfig struct {
	MaxTokens   int     `json:"maxTokens"`
	Temperature float64 `json:"temperature"`
}

type Messages struct {
	Role    string              `json:"role"`
	Content []map[string]string `json:"content"`
}

type BedrockResponse struct {
	Output     BedrockOutput  `json:"output"`
	StopReason string         `json:"stopReason"`
	Usage      map[string]int `json:"usage"`
}

type BedrockOutput struct {
	Message MessageResponse `json:"message"`
}

type MessageResponse struct {
	Content []map[string]string `json:"content"`
	Role    string              `json:"role"`
}
