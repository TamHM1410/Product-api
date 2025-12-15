package main

import "go.uber.org/zap"

func main() {
	sugar := zap.NewExample().Sugar()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", "1s",
	)
	sugar.Infof("Failed to fetch URL: %s", "http://example.com")
	sugar.Errorw("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", "1s",
	)

	logger := zap.NewExample()
	logger.Info("Info log example")
	logger.Error("Error log example")

}
