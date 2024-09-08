// Package api Service api implements a simple api world REST API.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"encore.dev/pubsub"
)

var AuditEvents = pubsub.NewTopic[*Events]("audit-events", pubsub.TopicConfig{
	DeliveryGuarantee: pubsub.AtLeastOnce,
})

func fetchSlideshowContent() (Slideshow, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", "https://httpbin.org/json", nil)
	if err != nil {
		return Slideshow{}, fmt.Errorf("error creating request: %w", err)
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return Slideshow{}, fmt.Errorf("error sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Slideshow{}, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON response
	var jsonResponse Slideshow
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return Slideshow{}, fmt.Errorf("error parsing JSON response: %w", err)
	}

	return jsonResponse, nil
}

// FetchSlideshow retrieves the slideshow data from external api.
//
//encore:api public path=/api
func FetchSlideshow(ctx context.Context) (*Response, error) {

	jsonResponse, err := fetchSlideshowContent()
	if err != nil {
		return nil, err
	}

	msgID, err := AuditEvents.Publish(ctx, &Events{
		Author: jsonResponse.Slideshow.Author,
	})
	if err != nil {
		return nil, fmt.Errorf("error publishing signal: %w", err)
	}

	// Return the response with both the greeting and the JSON data
	return &Response{
		MessageID: "Event successfully published, Message ID: " + msgID,
	}, nil
}
