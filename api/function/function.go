package function

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/a-h/gcp-data-skeleton/api/function/samples"
	"go.uber.org/zap"
)

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %v", err))
	}
	log.Info("cold start")

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("must have PROJECT_ID set")
	}
	topicID := os.Getenv("TOPIC_ID")
	if topicID == "" {
		log.Fatal("must have PROJECT_ID set")
	}

	ctx := context.Background()
	sh, err := samples.New(ctx, log, projectID, topicID)
	if err != nil {
		log.Fatal("failed to create /samples handler", zap.Error(err))
	}
	functions.HTTP("http", sh.ServeHTTP)
}
