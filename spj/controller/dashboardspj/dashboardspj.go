package dashboardhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"splcontroller/helper"
	"splcontroller/models"
)

// Import paket yang diperlukan
func IndexDashSpj(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	params := r.URL.Query()
	page := params.Get("page")
	yearmonth := params.Get("yearmonth")
	yearmonthend := params.Get("yearmonthend")
	lengthStr := params.Get("length")

	//  parameter page valid, jika tidak, kembalikan error
	if page != "production" && page != "claim" && page != "summary" {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// Konversi lengthStr ke integer
	length := 10 // Default length
	if lengthStr != "" {
		var err error
		length, err = strconv.Atoi(lengthStr)
		if err != nil {
			responseData := map[string]interface{}{
				"status":  false,
				"message": "failed, invalid length",
			}
			helper.ResponseJSON(w, http.StatusOK, responseData)
			return
		}
	}

	// Koneksi ke database
	models.ConnectDatabase()
	db := models.DB

	// Query untuk mendapatkan data sesuai dengan parameter yang diberikan
	var query string
	var results interface{} // Variabel untuk hasil yang akan dikembalikan
	var isEmptyResults bool // Variabel untuk memeriksa apakah hasil kosong
	switch page {
	case "production":
		var productionResults []models.ProductionData
		var periodParams string

		// Set periodParams ke 'null' jika yearmonth tidak diset
		if yearmonth == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "'"
		}
		yearmonthPtr := ""
		fmt.Println(yearmonthPtr)
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'production', %s) LIMIT ?", periodParams)
		fmt.Println("Query:", query)

		// Eksekusi query
		rows, err := db.Raw(query, length).Rows()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Iterasi setiap baris hasil query produksi
		for rows.Next() {
			var jsonData string
			// Pindai nilai kolom JSON ke dalam variabel string
			if err := rows.Scan(&jsonData); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Unmarshal nilai JSON ke dalam struktur ProductionData
			var result models.ProductionData
			if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Tambahkan hasil ke slice productionResults
			productionResults = append(productionResults, result)
		}

		results = productionResults
		isEmptyResults = len(productionResults) == 0

	case "claim", "summary":
		var resultsData interface{}

		// Buat parameter periode
		periodParams := ""
		if yearmonth == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "'"
		}

		if page == "claim" {
			var claimResults []models.ClaimData

			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'claim', %s) LIMIT ?", periodParams)
			fmt.Println("Query:", query)

			// Eksekusi query
			rows, err := db.Raw(query, length).Rows()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			// Iterasi setiap baris hasil query klaim
			for rows.Next() {
				var jsonData string
				// Pindai nilai kolom JSON ke dalam variabel string
				if err := rows.Scan(&jsonData); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Unmarshal nilai JSON ke dalam struktur ClaimData
				var result models.ClaimData
				if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Tambahkan hasil ke slice claimResults
				claimResults = append(claimResults, result)
			}

			if len(claimResults) == 0 {
				responseData := map[string]interface{}{
					"status":  false,
					"message": "failed",
				}
				helper.ResponseJSON(w, http.StatusOK, responseData)
				return
			}

			resultsData = claimResults
			isEmptyResults = len(claimResults) == 0

		} else if page == "summary" {
			var summaryResults []models.SummaryData

			// Buat parameter periode
			periodParams := ""
			if yearmonth == "" || yearmonthend == "" {
				periodParams = "null"
			} else {
				periodParams = fmt.Sprintf("'%s_%s'", yearmonth, yearmonthend)
			}

			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'summary_production', %s) LIMIT ?", periodParams)
			fmt.Println("Query:", query)

			// Eksekusi query
			rows, err := db.Raw(query, length).Rows()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			// Iterasi setiap baris hasil query summary
			for rows.Next() {
				var jsonData string
				// Pindai nilai kolom JSON ke dalam variabel string
				if err := rows.Scan(&jsonData); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Unmarshal nilai JSON ke dalam struktur SummaryData
				var result models.SummaryData
				if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Tambahkan hasil ke slice summaryResults
				summaryResults = append(summaryResults, result)
			}

			if len(summaryResults) == 0 {
				responseData := map[string]interface{}{
					"status":  false,
					"message": "failed",
				}
				helper.ResponseJSON(w, http.StatusOK, responseData)
				return
			}

			resultsData = summaryResults
			isEmptyResults = len(summaryResults) == 0
		}

		results = resultsData

	default:
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, invalid parameter",
		}
		helper.ResponseJSON(w, http.StatusOK, responseData)
		return
	}
	// Cek apakah results kosong atau tidak
	if results == nil || isEmptyResults {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data dashboard",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	responseData := map[string]interface{}{
		"data":    results,
		"status":  true,
		"message": "success",
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)
}
