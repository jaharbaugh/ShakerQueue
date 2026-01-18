package main

import (
	"encoding/json"
	"net/http"
	//"time"
	"fmt"
	"bytes"
	"io"
	//"context"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
)

func Health(sessionClient app.Client) (string, error) {
    req, err := http.NewRequest(
        http.MethodGet,
        sessionClient.BaseURL+"/health",
        bytes.NewBuffer(nil),
    )
    if err != nil {
        return "", err
    }

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ sessionClient.BearerToken)

	resp, err := sessionClient.HTTPClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Health Check failed: %s", resp.Status)
    }

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
    	return "", err
	}

	bodyString := string(bodyBytes)

	return bodyString, nil

}

func UpdateUserRole(sessionClient app.Client, email, newRole string) error{
	
	params := models.UpdateUserRoleRequest{
		Email: email,
		NewRole: newRole,
	}

	body, err := json.Marshal(params)
    if err != nil {
        return err
    }
	
	req, err := http.NewRequest(
        http.MethodPost,
        sessionClient.BaseURL+"/role/set",
        bytes.NewBuffer(body),
    )
    if err != nil {
        return err
    }

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ sessionClient.BearerToken)

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


func ListOrders(sessionClient app.Client) (models.ListOrderResponse, error) {
    req, err := http.NewRequest(
        http.MethodGet,
        sessionClient.BaseURL+"/orders/list",
        bytes.NewBuffer(nil),
    )
    if err != nil {
        return models.ListOrderResponse{}, err
    }

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ sessionClient.BearerToken)

	resp, err := sessionClient.HTTPClient.Do(req)
    if err != nil {
        return models.ListOrderResponse{}, err
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        return models.ListOrderResponse{}, fmt.Errorf("Order Retrieval Failed: %s", resp.Status)
    }

	var listOrderResp models.ListOrderResponse
    if err := json.NewDecoder(resp.Body).Decode(&listOrderResp); err != nil {
        return models.ListOrderResponse{}, err
    }

	return listOrderResp, nil

}