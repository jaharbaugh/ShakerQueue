package main

import (
	"encoding/json"
	"net/http"
	//"time"
	"bytes"
	"fmt"
	//"context"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func GetRecpies(sessionClient app.Client) (models.ListMenuResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		sessionClient.BaseURL+"/menu",
		bytes.NewBuffer(nil),
	)
	if err != nil {
		return models.ListMenuResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionClient.BearerToken)

	resp, err := sessionClient.HTTPClient.Do(req)
	if err != nil {
		return models.ListMenuResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ListMenuResponse{}, fmt.Errorf("Recipe Retrieval Failed: %s", resp.Status)
	}

	var listMenuResp models.ListMenuResponse
	if err := json.NewDecoder(resp.Body).Decode(&listMenuResp); err != nil {
		return models.ListMenuResponse{}, err
	}

	return listMenuResp, nil

}

func CreateOrder(sessionClient app.Client, cocktail string) (*models.CreateOrderResponse, error) {

	params := models.CreateOrderRequest{
		Name: cocktail,
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		sessionClient.BaseURL+"order/create",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionClient.BearerToken)

	//client := &http.Client{Timeout: 30 * time.Second}
	resp, err := sessionClient.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Order creation failed: %s", resp.Status)
	}

	var createOrderResp models.CreateOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&createOrderResp); err != nil {
		return nil, err
	}

	return &createOrderResp, nil
}

func OrderStatus(sessionClient app.Client) (models.OrderStatusResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		sessionClient.BaseURL+"/order/status",
		bytes.NewBuffer(nil),
	)
	if err != nil {
		return models.OrderStatusResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionClient.BearerToken)

	resp, err := sessionClient.HTTPClient.Do(req)
	if err != nil {
		return models.OrderStatusResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.OrderStatusResponse{}, fmt.Errorf("Recipe Retrieval Failed: %s", resp.Status)
	}

	var OrderStatusResp models.OrderStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&OrderStatusResp); err != nil {
		return models.OrderStatusResponse{}, err
	}

	return OrderStatusResp, nil
}
