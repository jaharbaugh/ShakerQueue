package main

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"bytes"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
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