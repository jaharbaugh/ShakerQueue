package handlers

import(
	//"context"
	//"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	//"github.com/jaharbaugh/ShakerQueue/internal/models"
	//"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func HandleStartOrder(deps app.Dependencies) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        if req.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        if !RequireRole(req, database.UserRoleEmployee, database.UserRoleAdmin) {
            RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
            return
        }

		orderIDStr := req.URL.Query().Get("id")
		if orderIDStr == "" {
			RespondWithError(w, http.StatusBadRequest, "Missing order id", nil)
			return
		}

		orderID, err := uuid.Parse(orderIDStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid order id", err)
			return
		}

        err = deps.Queries.UpdateOrderStatus(
            req.Context(),
            database.UpdateOrderStatusParams{
                ID:     orderID,
                Status: database.StatusPending,
            },
        )        
        if err != nil {
            RespondWithError(w, http.StatusInternalServerError, "Could not complete order", err)
            return
        }
/*
        order, err := deps.Queries.GetOrderByID(
            req.Context(),
            orderID,
        )


        switch order.Build{
        case database.BuildTypeShaken:
            time.Sleep(45 * time.Second)
        case database.BuildTypeCollins:
            time.Sleep(15 *time.Second)
        case database.BuildTypeSour:
            time.Sleep(time.Minute)
        case database.BuildTypeStirred:
            time.Sleep(30 *time.Second)
        }

        err = deps.Queries.UpdateOrderStatus(
            req.Context(),
            database.UpdateOrderStatusParams{
                ID:     orderID,
                Status: database.StatusComplete,
            },
        )        
        if err != nil {
            RespondWithError(w, http.StatusInternalServerError, "Could not complete order", err)
            return
        }
    */
       // RespondWithJSON(w, http.StatusOK, order)
        w.WriteHeader(http.StatusOK)
    }
}

func HandleCompleteOrder(deps app.Dependencies) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        if req.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        if !RequireRole(req, database.UserRoleEmployee, database.UserRoleAdmin) {
            RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
            return
        }

		orderIDStr := req.URL.Query().Get("id")
		if orderIDStr == "" {
			RespondWithError(w, http.StatusBadRequest, "Missing order id", nil)
			return
		}

		orderID, err := uuid.Parse(orderIDStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid order id", err)
			return
		}
/*
        err = deps.Queries.UpdateOrderStatus(
            req.Context(),
            database.UpdateOrderStatusParams{
                ID:     orderID,
                Status: database.StatusPending,
            },
        )        
        if err != nil {
            RespondWithError(w, http.StatusInternalServerError, "Could not complete order", err)
            return
        }
*/
/*
        order, err := deps.Queries.GetOrderByID(
            req.Context(),
            orderID,
        )

   */
        err = deps.Queries.UpdateOrderStatus(
            req.Context(),
            database.UpdateOrderStatusParams{
                ID:     orderID,
                Status: database.StatusComplete,
            },
        )        
        if err != nil {
            RespondWithError(w, http.StatusInternalServerError, "Could not complete order", err)
            return
        }
 
        w.WriteHeader(http.StatusNoContent)
    }
}

