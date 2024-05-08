package uploadperiodspl

import (
	"log"
	"net/http"
	"time"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)
func UploadPeriodSpl(w http.ResponseWriter, r *http.Request) {
	// Parse tahun dan bulan dari query parameter "yearmonth"
	yearMonth := r.URL.Query().Get("yearmonth")

	// Jika query parameter kosong, gunakan tahun dan bulan saat ini
	if yearMonth == "" {
		yearMonth = time.Now().Format("200601")
	}

	// Variabel untuk menyimpan hasil query
	var results []models.Period

	// Parse tahun dan bulan dari format "YYYYMM"
	parsedTime, err := time.Parse("200601", yearMonth)
	if err != nil {
		// Menangani kesalahan
		log.Println("Error parsing date:", err)
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal server error: " + err.Error(),
		})
		return
	}

	// Format tanggal menjadi "Aug 2021"
	label := parsedTime.Format("Jan 2006")

	// Tambahkan hasil ke slice results
	results = append(results, models.Period{
		YearMonth: yearMonth,
		Label:     label,
	})

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"data":    results,
		"status":  true,
		"message": "success",
	})
}

