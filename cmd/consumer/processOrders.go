package main

import (
	"context"
	"fmt"
	"net/http"
	//"time"

	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func ProcessOrder(sessionClient app.Client) func(models.OrderEvent) queue.Acktype {
	return func(event models.OrderEvent) queue.Acktype {
		fmt.Printf("Processing order %s\n", event.OrderID)

		url := fmt.Sprintf(
			"%s/orders/complete?id=%s",
			sessionClient.BaseURL,
			event.OrderID,
		)

		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			url,
			nil,
		)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)
			return queue.NackRequeue
		}

		// Attach auth (example: bearer token)
		req.Header.Set("Authorization", "Bearer "+sessionClient.BearerToken)

		resp, err := sessionClient.HTTPClient.Do(req)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
			return queue.NackRequeue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			fmt.Printf(
				"Failed to complete order %s: status %d\n",
				event.OrderID,
				resp.StatusCode,
			)
			return queue.NackRequeue
		}

		fmt.Printf("Order %s completed\n", event.OrderID)
		return queue.Ack
	}
}