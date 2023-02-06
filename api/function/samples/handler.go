package samples

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/a-h/gcp-data-skeleton/api/function/models"
	"github.com/a-h/gcp-data-skeleton/api/function/pubsub"
	"github.com/a-h/respond"
	"go.uber.org/zap"
)

func New(ctx context.Context, log *zap.Logger, projectID, topicID string) (h *Handler, err error) {
	client, err := pubsub.NewClient[models.Sample](ctx, projectID, topicID)
	return &Handler{Log: log, Client: client}, err
}

type Handler struct {
	Log    *zap.Logger
	Client pubsub.Client[models.Sample]
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		respond.WithError(w, "missing body", http.StatusBadRequest)
		return
	}
	var sample models.Sample
	err := json.NewDecoder(r.Body).Decode(&sample)
	if err != nil {
		respond.WithError(w, "could not decode request body", http.StatusBadRequest)
		return
	}
	if sample.Name == "" {
		respond.WithError(w, "invalid sample name", http.StatusBadRequest)
		return
	}

	h.Log.Info("publishing message")
	serverID, err := h.Client.Publish(r.Context(), pubsub.Message[models.Sample]{Data: sample})
	if err != nil {
		h.Log.Error("failed to publish sample", zap.Error(err))
		respond.WithError(w, "failed to publish sample", http.StatusInternalServerError)
		return
	}

	h.Log.Info("published sample", zap.String("serverId", serverID))
	respond.WithJSON(w, models.SamplePostResponse{OK: true}, http.StatusOK)
	return
}
