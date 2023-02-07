package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/gcp-data-skeleton/api/samples"
	"go.uber.org/zap"
)

func New(ctx context.Context, log *zap.Logger, projectID, topicID string) (routes http.Handler, err error) {
	mux := http.NewServeMux()

	sh, err := samples.New(ctx, log, projectID, topicID)
	if err != nil {
		return mux, fmt.Errorf("routes: failed to create /samples handler: %w", err)
	}
	mux.Handle("/samples", sh)

	return mux, nil
}
