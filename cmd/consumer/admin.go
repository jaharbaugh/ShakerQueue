package main

import (
	//"encoding/json"
	"net/http"
	//"time"
	"fmt"
	"bytes"
	"io"
	//"context"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	//"github.com/jaharbaugh/ShakerQueue/internal/models"
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

	fmt.Println("Server Health Status:")
	fmt.Printf(bodyString)

	return bodyString, nil

}

func UpdateUserRole(){

}

func ListOrders(){

}