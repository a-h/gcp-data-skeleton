package function

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/a-h/gcp-data-skeleton/api"
	"github.com/a-h/gcp-data-skeleton/onevent/samplepublished"
	"go.uber.org/zap"
)

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %v", err))
	}
	log.Info("cold start")

	target := os.Getenv("FUNCTION_TARGET")

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("must have PROJECT_ID set")
	}
	ctx := context.Background()

	// Set up HTTP handlers.
	if target == "http" {
		topicID := os.Getenv("TOPIC_ID")
		if topicID == "" {
			log.Fatal("must have PROJECT_ID set")
		}

		h, err := api.New(ctx, log, projectID, topicID)
		if err != nil {
			log.Fatal("failed to create handlers", zap.Error(err))
		}
		functions.HTTP("http", h.ServeHTTP)
	}

	// Set up event handlers.
	if target == "on-sample-published" {
		osp := samplepublished.New(log)
		functions.CloudEvent("on-sample-published", osp.Handle)
	}
}
