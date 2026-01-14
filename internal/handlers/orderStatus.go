package handlers

import (
	"net/http"
	//"context"

	//"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	//"github.com/jaharbaugh/ShakerQueue/internal/models"
)

func HandleOrderStatus(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//order, err := deps.Queries.GetOrderByID(context.Background(), orderID)

	}
}
