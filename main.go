package main

import (
	"context"
	"log/slog"
)

func main() {
	ctx := context.Background()
	result, err := InvokeBedrock(ctx)
	if err != nil {
		slog.Error("failed to invoke bedrock", "error", err)
	}

	slog.Info("invoke bedrock result", "result", result)
}
