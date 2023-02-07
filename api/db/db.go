package db

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/a-h/gcp-data-skeleton/models"
	"google.golang.org/api/iterator"
)

func NewSamples(ctx context.Context, projectID string) (s *Samples, err error) {
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return
	}
	s = &Samples{
		Client:     client,
		Collection: client.Collection("samples"),
	}
	return
}

type Samples struct {
	Client     *firestore.Client
	Collection *firestore.CollectionRef
}

func (s *Samples) Upsert(ctx context.Context, id string, sample models.Sample) (err error) {
	_, err = s.Collection.Doc(id).Set(ctx, sample)
	return
}

func (s *Samples) Get(ctx context.Context, id string) (sample models.Sample, err error) {
	doc, err := s.Collection.Doc(id).Get(ctx)
	if err != nil || !doc.Exists() {
		return
	}
	err = doc.DataTo(&sample)
	return
}

func (s *Samples) Query(ctx context.Context, name string) (samples []models.Sample, err error) {
	q := s.Collection.
		Where("Name", "==", name).
		Documents(ctx)

	for {
		doc, err := q.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return samples, err
		}
		var sample models.Sample
		err = doc.DataTo(&sample)
		if err != nil {
			return samples, err
		}
		samples = append(samples, sample)
	}

	return
}
