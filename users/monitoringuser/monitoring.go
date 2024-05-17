package monitoring

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"userscontroller/helper"
	"userscontroller/models"
)
func IndexMonuser(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from URL
	queryValues := r.URL.Query()
	start := queryValues.Get("start")
	end := queryValues.Get("end")
	email := queryValues.Get("email")

	// Retrieve the length parameter
	length := 10
	if lenStr := queryValues.Get("length"); lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}
	fmt.Println("pagination = ", length)

	// Calculate end date
	var endDate time.Time
	var endUpdate string // Change endUpdate to string

	if end != "" {
		var err error
		endDate, err = time.Parse("2006-01-02", end) // Format tanggal sesuai dengan "YYYY-MM-DD"
		if err != nil {
			response := map[string]interface{}{"message": "Invalid date format", "status": false}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		endDate = endDate.Add(24 * time.Hour)
		endUpdate = endDate.Format("2006-01-02 15:04:05") // Atur format tanggal dan waktu yang sesuai
	}

	// Koneksi ke database
	models.ConnectDatabase()
	db := models.DB

	// Query the database based on filters
	var data []models.Monitoring
	var count int64

	// Query for count
	countQuery := db.Table("dashboard.activity_log").Select("COUNT(*)")
	if email != "" {
		countQuery = countQuery.Where("email = ?", email)
	}
	if start != "" && endUpdate != "" {
		countQuery = countQuery.Where("datetime BETWEEN ? AND ?", start, endUpdate)
	}
	countQuery.Count(&count)

	// Debugging count query
	countQueryStr := countQuery.Debug().Statement.SQL.String()
	fmt.Println("Count Query:", countQueryStr)

	// Query for data
	query := db.Table("dashboard.activity_log").Select("id", "email", "name", "role", "activity", "datetime", "location", "ip").Order("datetime DESC")
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if start != "" && endUpdate != "" {
		query = query.Where("datetime BETWEEN ? AND ?", start, endUpdate)
	}
	query = query.Limit(length)

	// Print SQL query
	sqlStr := query.Debug().Find(&data).Statement.SQL.String()
	fmt.Println("Data Query:", sqlStr)

	// Execute query
	if err := query.Find(&data).Error; err != nil {
		fmt.Println("Error executing query:", err)
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Print result
	fmt.Println("Data:", data)

	// Convert time.Time to string format
	layout := "2006-01-02 15:04:05"
	for i := range data {
		data[i].DateTimeStr = data[i].DateTime.Format(layout)
	}

	// Create response
	response := models.LengthAwarePaginator{
		Items:       data,
		PerPage:     length,
		CurrentPage: 1,
		Path:        "http://localhost:8000/monitoring",
		Query:       map[string]string{},
		Fragment:    "",
		PageName:    "page",
		OnEachSide:  3,
		Options:     map[string]string{},
		Total:       int(count),
		LastPage:    int(math.Ceil(float64(count) / float64(length))),
	}

	// Respond with the data
	helper.ResponseJSON(w, http.StatusOK, response)
}
