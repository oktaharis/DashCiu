package dashboardhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

// Import paket yang diperlukan
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Dekode input JSON dari request body
	var dashboardInput map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dashboardInput); err != nil {
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ambil nilai parameter dari input JSON
	yearmonth := dashboardInput["yearmonth"]
	yearmonthend := dashboardInput["yearmonthend"]
	page := dashboardInput["page"]
	app := dashboardInput["app"]
	app_child := dashboardInput["app_child"]
	// Pastikan parameter page valid, jika tidak, kembalikan error
	if page != "production" && page != "claim" && page != "summary" {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	var yearmonthPtr *string

	// Set nilai default yearmonth jika tidak disediakan
	if yearmonth == "" {
		yearmonthPtr = nil
	} else {
		yearmonthPtr = &yearmonth
	}

	// Koneksi ke database
	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Query untuk mendapatkan data sesuai dengan parameter yang diberikan
	var query string
	var results interface{} // Variabel untuk hasil yang akan dikembalikan
	switch page {
	case "production":
		var query string
		var productionResults []models.ProductionData
		var periodParams string

		// Set periodParams ke 'null' jika yearmonth tidak diset
		if yearmonth == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "'"
		}

		// KPI & Atome
		if app == "kpi" || app == "afi" {
			paramApp := app + "|" + app_child
			if app_child == "All" {
				paramApp = app
			}
			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', 'production', %s)", paramApp, periodParams)
		} else {
			query = "SELECT * FROM dashboard.sp_dashboard('admin', 'production', ?)"
		}

		// Eksekusi query
		rows, err := db.Raw(query, yearmonthPtr).Rows()
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
		query = "SELECT * FROM dashboard.sp_dashboard('admin', 'claim', ?)"
		var claimResults []models.ClaimData
		rows, err := db.Raw(query, yearmonthPtr).Rows()
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
		var query string
		var summaryResults []models.SummaryData

		// Buat parameter periode
		periodParams := ""
		if yearmonth == "" || yearmonthend == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "_" + yearmonthend + "'"
		}

		// Buat query sesuai dengan kondisi aplikasi (kpi, afi, atau lainnya)
		if app == "kpi" || app == "afi" {
			paramApp := app + "|" + app_child
			if app_child == "All" {
				paramApp = app
			}
			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', 'summary_production', %s)", paramApp, periodParams)
		} else {
			query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'summary_production', %s)", periodParams)
			fmt.Println(periodParams)
		}

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
