package claimlist

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

func IndexClaimList(w http.ResponseWriter, r *http.Request) {
	// Dekode input JSON dari request body
	var dashboardInput map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dashboardInput); err != nil {
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ambil nilai parameter dari input JSON
	_ = dashboardInput["yearmonth"]
	app := dashboardInput["app"]
	appChild := dashboardInput["app_child"]

	// Koneksi ke database
	db := models.DBConnections[app]
	fmt.Println(db)
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}
	fmt.Println(app)

	// Query untuk mendapatkan data dari database
	var query string
	if app == "kpi" || app == "afi" {
		paramApp := app + "|" + appChild
		if appChild == "All" {
			paramApp = app
		}
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_filter('admin', 'production|period', '%s');", paramApp)
	} else {
		query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
	}

	// Eksekusi query
	rows, err := db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Ambil periode dan labelnya
	var periods []models.Period
	_ = ""

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var yearMonth, label string
		// Pindai nilai kolom ke dalam variabel struktur
		if err := rows.Scan(&yearMonth, &label); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		period := models.Period{
			YearMonth: yearMonth,
			Label:     label,
		}

		_ = append(periods, period)
	}

	// Set parameter periode default jika tidak disediakan
	// if yearmonth == "" && len(periods) > 0 {
	// 	yearmonth = periods[len(periods)-1].YearMonth
	// }

}
