package monitor

import (
	"context"
	"time"

	"encore.app/api"
	"encore.dev/pubsub"
)

var _ = pubsub.NewSubscription(
	api.AuditEvents, "get-audit-events",
	pubsub.SubscriptionConfig[*api.Events]{
		Handler: SetEvents,
	},
)

func SetEvents(ctx context.Context, api *api.Events) error {
	insert(ctx, api.Author)

	return nil
}

// GetAuthor Get retrieves the original URL for the id.
//
//encore:api public method=GET path=/slideshow/:id
func GetAuthor(ctx context.Context, id string) (*api.Events, error) {
	// var author_name string
	var created_at time.Time
	err := db.QueryRow(ctx, `
        SELECT created_at FROM events
        WHERE id = $1
    `, id).Scan(&created_at)
	return &api.Events{
		CreatedAt: created_at,
	}, err
}
