package dashboardkpi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kpicontroller/helper"
	"kpicontroller/models"
)

// Import paket yang diperlukan
func IndexDashKpi(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	yearmonth := queryParams.Get("yearmonth")
	yearmonthend := queryParams.Get("yearmonthend")
	page := queryParams.Get("page")
	app_child := queryParams.Get("app_child")

	//  parameter page valid, jika tidak, kembalikan error
	if page != "production" && page != "claim" && page != "summary" {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// Set nilai default yearmonth jika tidak disediakan
	var periodParams interface{}
	if yearmonth == "" || yearmonthend == "" {
		periodParams = "null"
	} else {
		periodParams = fmt.Sprintf("'%s_%s'", yearmonth, yearmonthend)
	}

		// Koneksi ke database
		models.ConnectDatabase()
		db := models.DB

	// Jika app_child kosong, beri nilai default "All"
	if app_child == "" {
		app_child = "All"
	}

	// Query untuk mendapatkan data sesuai dengan parameter yang diberikan
	var query string
	var results interface{} // Variabel untuk hasil yang akan dikembalikan
	switch page {
	case "production":
		var productionResults []models.ProductionData

		// KPI & Atome
		paramApp := "kpi" + "|" + app_child
		if app_child == "All" {
			paramApp = "kpi"
		}
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', 'production', %s)", paramApp, periodParams)
		fmt.Println(query)

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

	case "claim":
		var claimResults []models.ClaimData

		// KPI & Atome
		paramApp := "kpi" + "|" + app_child
		if app_child == "All" {
			paramApp = "kpi"
		}
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', 'claim', %s)", paramApp, periodParams)
		fmt.Println(query)

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

		results = claimResults

	case "summary":
		var summaryResults []models.SummaryData

		// KPI & Atome
		paramApp := "kpi" + "|" + app_child
		if app_child == "All" {
			paramApp = "kpi"
		}
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', 'summary_production', %s)", paramApp, periodParams)
		fmt.Println(query)

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

		results = summaryResults

	default:
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	responseData := map[string]interface{}{
		"data":    results,
		"status":  true,
		"message": "Berhasil mengambil data dashboard",
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)
}
