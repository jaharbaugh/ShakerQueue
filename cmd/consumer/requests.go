package main

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"bytes"
	"context"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
)

func Login(serverURL string, creds models.LogInRequest) (*models.LogInResponse, error) {
    body, err := json.Marshal(creds)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(
        http.MethodPost,
        serverURL+"/login",
        bytes.NewBuffer(body),
    )
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("login failed: %s", resp.Status)
    }

    var loginResp models.LogInResponse
    if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
        return nil, err
    }

    return &loginResp, nil
}

func RegisterUser(serverURL string, creds models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	body, err := json.Marshal(creds)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(
        http.MethodPost,
        serverURL+"/register",
        bytes.NewBuffer(body),
    )
    if err != nil {
        return nil, err
    }
	
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return nil, fmt.Errorf("registration failed: %s", resp.Status)
    }

	var registrationResp models.RegisterUserResponse
    if err := json.NewDecoder(resp.Body).Decode(&registrationResp); err != nil {
        return nil, err
    }

    return &registrationResp, nil
}

func CreateOrder(sessionClient app.Client, cocktail string) (*models.CreateOrderResponse, error) {

	params := models.CreateOrderParams{
		Name: cocktail,
	}

	body, err := json.Marshal(params)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(
        http.MethodPost,
        sessionClient.BaseURL+"/createorder",
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