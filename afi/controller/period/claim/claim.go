package claimperiodafi

import (
	"fmt"
	"net/http"

	"aficontroller/helper"
	"aficontroller/models"
)

func ClaimPeriodAfi(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL

models.ConnectDatabase()
db := models.DB

	// Query untuk mendapatkan semua data periode
	query := fmt.Sprintf("SELECT yearmonth, label FROM dashboard.sp_filter('admin', 'claim|period', 'afi')")

	// Eksekusi query
	rows, err := db.Raw(query).Rows()
	if err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	// Variabel untuk menyimpan hasil query
	var results []models.Period

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var period models.Period
		// Scan nilai kolom ke struct
		if err := rows.Scan(&period.YearMonth, &period.Label); err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		// Tambahkan hasil ke slice results
		results = append(results, period)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	var responseData []map[string]interface{}
	for _, period := range results {
		responseData = append(responseData, map[string]interface{}{
			"YearMonth": "Batch - " + period.Label,
			"Label":     period.Label,
		})
	}
	if len(responseData) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data period claim",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"data":    responseData,
		"status":  true,
		"message": "success",
	})
}