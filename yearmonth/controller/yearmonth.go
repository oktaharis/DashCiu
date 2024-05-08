package yearmonth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

func IndexYear(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	params := r.URL.Query()
	page := params.Get("page")
	yearmonth := params.Get("yearmonth")
	yearmonthend := params.Get("yearmonthend")
	lengthStr := params.Get("length")

	// Konversi lengthStr ke integer
	length := 10 // Default length
	if lengthStr != "" {
		var err error
		length, err = strconv.Atoi(lengthStr)
		if err != nil {
			http.Error(w, "Invalid length parameter", http.StatusBadRequest)
			return
		}
	}

	// Verifikasi parameter `page` agar valid
	if page != "production" && page != "claim" && page != "summary" {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	app := "spl"
	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Definisikan query berdasarkan parameter `page`
	var query string
	if page == "summary" {
		// Untuk halaman "summary", pastikan untuk menangani `yearmonth` dan `yearmonthend`
		periodParams := ""
		if yearmonth == "" || yearmonthend == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "_" + yearmonthend + "'"
		}

		// Query untuk mendapatkan data summary
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', 'summary_production', %s) LIMIT %d", periodParams, length)
	} else {
		// Query untuk halaman lain ("production" atau "claim")
		periodParams := ""
		if yearmonth == "" {
			periodParams = "null"
		} else {
			periodParams = "'" + yearmonth + "'"
		}

		// Query untuk halaman lain
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_dashboard('admin', '%s', %s) LIMIT %d", page, periodParams, length)
	}

	// Eksekusi query
	rows, err := db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Variabel untuk menyimpan hasil query
	var results []models.Period

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var jsonData string
		// Pindai nilai kolom JSON ke dalam variabel string
		if err := rows.Scan(&jsonData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal nilai JSON ke dalam struktur `Period`
		var result models.Period
		if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Tambahkan hasil ke slice results
		results = append(results, result)
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



