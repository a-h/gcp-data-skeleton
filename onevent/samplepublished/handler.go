package samplepublished

import (
	"context"

	"github.com/a-h/gcp-data-skeleton/models"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
)

func New(log *zap.Logger) *Handler {
	return &Handler{
		Log: log,
	}
}

type Handler struct {
	Log *zap.Logger
}

func (h *Handler) Handle(ctx context.Context, event cloudevents.Event) error {
	h.Log.Info("processing event", zap.Any("event", event.Data))
	var sample models.Sample
	err := event.DataAs(&sample)
	if err != nil {
		h.Log.Error("failed to unmarshal sample", zap.Error(err))
		return err
	}
	h.Log.Info("processed event", zap.Any("sample", sample))
	return nil
}
