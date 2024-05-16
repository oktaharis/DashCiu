package dashboardflexi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"flexicontroller/helper"
	"flexicontroller/models"
)

func IndexDashFlexi(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	params := r.URL.Query()
	page := params.Get("page")
	yearmonth := params.Get("yearmonth")
	yearmonthend := params.Get("yearmonthend")
	lengthStr := params.Get("length")

	//  parameter page valid, jika tidak, kembalikan error
	if page != "production" && page != "claim" && page != "summary" {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, invald page parameter",
		}
		helper.ResponseJSON(w, http.StatusOK, responseData)
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

	// Ambil nilai parameter app
	app := "flexi"

	// Koneksi ke database
	db := models.DBConnections[app]
	fmt.Println(db)
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Query untuk mendapatkan data sesuai dengan parameter yang diberikan
	var query string
	var results interface{} // Variabel untuk hasil yang akan dikembalikan
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
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'production', %s) LIMIT %d", periodParams, length)
		fmt.Println("Query:", query)

		// Eksekusi query
		rows, err := db.Raw(query).Rows()
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

			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'claim', %s) LIMIT %d", periodParams, length)
			fmt.Println("Query:", query)

			// Eksekusi query
			rows, err := db.Raw(query).Rows()

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

		} else if page == "summary" {
			var summaryResults []models.SummaryData

			// Buat parameter periode
			periodParams := ""
			if yearmonth == "" || yearmonthend == "" {
				periodParams = "null"
			} else {
				periodParams = fmt.Sprintf("'%s_%s'", yearmonth, yearmonthend)
			}

			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'summary_production', %s) LIMIT %d", periodParams, length)
			fmt.Println("Query:", query)

			// Eksekusi query
			rows, err := db.Raw(query).Rows()
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
		}

		results = resultsData

	default:
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, invalid page parameter",
		}
		helper.ResponseJSON(w, http.StatusOK, responseData)
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
