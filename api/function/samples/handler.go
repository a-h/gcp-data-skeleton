package samples

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/gcp-data-skeleton/api/function/db"
	"github.com/a-h/gcp-data-skeleton/api/function/models"
	"github.com/a-h/gcp-data-skeleton/api/function/pubsub"
	"github.com/a-h/respond"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

func New(ctx context.Context, log *zap.Logger, projectID, topicID string) (h *Handler, err error) {
	ps, err := pubsub.NewClient[models.Sample](ctx, projectID, topicID)
	if err != nil {
		return nil, fmt.Errorf("samples: could not create pubsub client: %w", err)
	}
	sdb, err := db.NewSamples(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("samples: could not create samples database client: %w", err)
	}
	h = &Handler{
		Log:    log,
		PubSub: ps,
		DB:     sdb,
	}
	return
}

type Handler struct {
	Log    *zap.Logger
	PubSub pubsub.Client[models.Sample]
	DB     *db.Samples
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

	h.Log.Info("saving to database")
	err = h.DB.Upsert(r.Context(), ksuid.New().String(), sample)
	if err != nil {
		h.Log.Error("failed to save to database", zap.Error(err))
		respond.WithError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.Log.Info("publishing message")
	serverID, err := h.PubSub.Publish(r.Context(), pubsub.Message[models.Sample]{Data: sample})
	if err != nil {
		h.Log.Error("failed to publish sample", zap.Error(err))
		respond.WithError(w, "failed to publish sample", http.StatusInternalServerError)
		return
	}

	h.Log.Info("published sample", zap.String("serverId", serverID))
	respond.WithJSON(w, models.SamplePostResponse{OK: true, ServerID: serverID}, http.StatusOK)
	return
}
