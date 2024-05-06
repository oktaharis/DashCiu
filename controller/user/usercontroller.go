package user

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/jeypc/homecontroller/helper"
// 	"github.com/jeypc/homecontroller/models"
// )

// func UserMan(w http.ResponseWriter, r *http.Request) {
// 	// Decode input data from request body
// 	var userInput map[string]string
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&userInput); err != nil {
// 		response := map[string]interface{}{"message": err.Error(), "status": false}
// 		helper.ResponseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	// Get app from input
// 	app := userInput["app"]

// 	// Connect to the database
// 	db := models.DBConnections[app]
// 	if db == nil {
// 		models.ConnectDatabase(app)
// 		db = models.DBConnections[app]
// 	}

// 	// Verify user authentication and role
// 	currentUser := getCurrentUser(r) // Implement this function to get current user
// 	if currentUser == nil || currentUser.RoleID != 1 {
// 		http.Redirect(w, r, "/", http.StatusForbidden)
// 		return
// 	}

// 	// Retrieve user data from database
// 	var users []models.User
// 	db.Where("id != ?", currentUser.ID).Find(&users)

// 	// Process users and get roles and products
// 	for i, user := range users {
// 		// Get role
// 		var role models.Role
// 		db.First(&role, user.RoleID)
// 		users[i].Role = role

// 		// Get products if they exist
// 		if user.ProductID != "" {
// 			var productIDs []int
// 			for _, idStr := range strings.Split(user.ProductID, ",") {
// 				id, err := strconv.Atoi(idStr)
// 				if err == nil {
// 					productIDs = append(productIDs, id)
// 				}
// 			}

// 			var products []models.Product
// 			db.Where("id IN (?)", productIDs).Find(&products)
// 			users[i].Products = products
// 		}
// 	}

// 	// Prepare the response
// 	response := map[string]interface{}{
// 		"data": users,
// 	}

// 	// Respond with JSON
// 	helper.ResponseJSON(w, http.StatusOK, response)
// }
