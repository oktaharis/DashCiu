package dashboardhandler

import (
	"encoding/json"
	"net/http"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

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
	page := dashboardInput["page"]

	// Pastikan parameter page valid, jika tidak, kembalikan error
	if page != "production" && page != "claim" && page != "summary" {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	db := models.DBConnections["spl"]
	if db == nil {
		models.ConnectDatabase("spl")
		db = models.DBConnections["spl"]
	}

	// Query untuk mendapatkan data sesuai dengan parameter yang diberikan
	var query string
	switch page {
	case "production":
		query = "SELECT * FROM dashboard.sp_dashboard('admin', 'production', ?)"
	case "claim":
		query = "SELECT * FROM dashboard.sp_dashboard('admin', 'claim', ?)"
	case "summary":
		query = "SELECT * FROM dashboard.sp_dashboard('admin', 'summary', NULL)"
	}

	// Jalankan query
	rows, err := db.Raw(query, yearmonth).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Inisialisasi variabel untuk menyimpan hasil
	var results []models.QueryResult

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var jsonData string
		// Pindai nilai kolom JSON ke dalam variabel string
		if err := rows.Scan(&jsonData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal nilai JSON ke dalam struktur QueryResult
		var result models.QueryResult
		if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Tambahkan hasil ke slice results
		results = append(results, result)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	responseData := map[string]interface{}{
		"data": results,
        "status": true,
        "message": "Berhasil mengambil data dashboard",
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)
}
