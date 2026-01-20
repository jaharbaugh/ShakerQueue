package handlers

import(
	//"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	//"github.com/jaharbaugh/ShakerQueue/internal/auth"
	"github.com/jaharbaugh/ShakerQueue/internal/database"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	//"github.com/jaharbaugh/ShakerQueue/internal/queue"
)

func HandleCreateCocktailRecipe(deps app.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

/*		userID, ok := req.Context().Value(auth.UserIDKey).(uuid.UUID)
		if !ok {
			RespondWithError(w, http.StatusUnauthorized, "User not authenticated", nil)
			return
		}
*/
		if !RequireRole(req, database.UserRoleAdmin, database.UserRoleEmployee) {
			RespondWithError(w, http.StatusForbidden, "Insufficient permissions", nil)
			return
		}

		newRecipeRequest := models.CreateCocktailRecipeRequest{}

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&newRecipeRequest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
			return
		}
	
		var buildType database.BuildType
		switch newRecipeRequest.BuildType{
		case "shaken":
			buildType = database.BuildTypeShaken
		case "stirred":
			buildType = database.BuildTypeStirred
		case "collins":
			buildType = database.BuildTypeCollins
		case "sour":
			buildType = database.BuildTypeSour
		default:
			RespondWithError(w, http.StatusBadRequest, "Unknown build type", nil)
			return
		} 

		ingredientsData, err := json.Marshal(newRecipeRequest.Ingredients)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not marshal ingredients", err)
			return
		}
		rawIngredientsData := json.RawMessage(ingredientsData)

		recipe := database.CreateCocktailRecipeParams{
			ID: uuid.New(),
			Name: newRecipeRequest.Name,
			Ingredients: rawIngredientsData,
			Build: buildType,
		}

		newCocktailRecipe, err := deps.Queries.CreateCocktailRecipe(req.Context(), recipe)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not insert recipe", err)
			return
		}

		RespondWithJSON(w, http.StatusOK, models.CreateCocktailRecipeResponse{
			Recipe: newCocktailRecipe,
		})
	}
}