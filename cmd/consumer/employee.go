package main

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	//"log"

	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func JoinQueue(sessionClient app.Client) error {
	err := queue.SubscribeJSON(
		sessionClient.AMQPConn,
		queue.ExchangeDirect,
		"orders.employee",
		"orders.created",
		queue.SimpleQueueDurable,
		ProcessOrder(sessionClient),
	)
	if err != nil {
		return fmt.Errorf("Failed to subscribe: %v", err)
	}

	return nil

}

func ProcessOrder(sessionClient app.Client) func(models.OrderEvent) queue.Acktype {
	return func(event models.OrderEvent) queue.Acktype {
		fmt.Printf("Processing order %s\n", event.OrderID)

		url := fmt.Sprintf(
			"%s/orders/start?id=%s",
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

		if resp.StatusCode != http.StatusOK {
			fmt.Printf(
				"Failed to complete order %s: status %d\n",
				event.OrderID,
				resp.StatusCode,
			)
			return queue.NackRequeue
		}

		time.Sleep(time.Duration(event.Delay) * time.Second)

		url = fmt.Sprintf(
			"%s/orders/complete?id=%s",
			sessionClient.BaseURL,
			event.OrderID,
		)

		req, err = http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			url,
			nil,
		)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)
			return queue.NackRequeue
		}

		fmt.Printf("Order %s completed\n", event.OrderID)
		return queue.Ack
	}
}

func AddCocktailRecipe(sessionClient app.Client, name string, ingredients map[string]string, buildType string) error {

	params := models.CreateCocktailRecipeRequest{
		Name: name,
		Ingredients: ingredients,
		BuildType: buildType,
	}

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		sessionClient.BaseURL+"/menu/add",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionClient.BearerToken)

	resp, err := sessionClient.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("User role update failed: %s", resp.Status)
	}

	return nil
}

