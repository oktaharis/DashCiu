package uploadperiodspj

import (
	"log"
	"net/http"
	"time"

	"splcontroller/helper"
	"splcontroller/models"
)
func UploadPeriodSpj(w http.ResponseWriter, r *http.Request) {
	yearMonth := r.URL.Query().Get("yearmonth")

	// Jika yearMonth kosong, beri nilai default "202312"
	if yearMonth == "" {
		yearMonth = "202312"
	}

	var results []models.Period

	parsedTime, err := time.Parse("200601", yearMonth)
	if err != nil {
		log.Println("Error parsing date:", err)
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  false,
			"message": "Internal server error: " + err.Error(),
		})
		return
	}

	for i := 0; i < 13; i++ {
		newMonth := parsedTime.AddDate(0, i, 0)
		yearMonth := newMonth.Format("200601")
		label := newMonth.Format("Jan 2006")
		results = append(results, models.Period{
			YearMonth: yearMonth,
			Label:     label,
			Selected:  false,
		})
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"data":    results,
		"status":  true,
		"message": "success",
	})
}

