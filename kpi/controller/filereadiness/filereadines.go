package filereadinesskpi

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"kpicontroller/helper"
	"kpicontroller/models"
)

func IndexFilesAfi(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan nilai dari URL query parameters
	queryParams := r.URL.Query()
	yearmonthStr := queryParams.Get("yearmonth")
	appChild := queryParams.Get("app_child")

	yearmonth, _ := strconv.Atoi(yearmonthStr)

	// Koneksi ke database
	models.ConnectDatabase()
	db := models.DB

	// Query untuk mendapatkan periode
	var query string

	query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period', 'kpi');"
	fmt.Println(query)

	var periods []models.Period
	rows, err := db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var period models.Period
		if err := rows.Scan(&period.YearMonth, &period.Label); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		periods = append(periods, period)
	}

	// Set parameter periode default jika tidak disediakan
	if yearmonth == 0 && len(periods) > 0 {
		yearmonthInt, err := strconv.Atoi(periods[len(periods)-1].YearMonth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		yearmonth = yearmonthInt
	}

	// Query untuk mendapatkan data files redines
	query = "SELECT yearmonth, label, policy, claim, updated_at, is_process, type, summary_production, summary_claim FROM dashboard.files_redines WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}

	if appChild != "All" {
		filters = append(filters, fmt.Sprintf("type = '%s'", appChild))
	}

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("inifilters", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.files_redines WHERE "
	// Tambahkan klausa filter sesuai dengan kueri utama
	if len(filters) > 0 {
		countQuery += strings.Join(filters, " AND ")
	} else {
		countQuery += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}

	var totalCount int
	err = db.Raw(countQuery).Row().Scan(&totalCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Paginasi data
	page := 1
	pageLength := 10
	offset := (page - 1) * pageLength

	// Tambahkan limit dan offset ke dalam query
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageLength, offset)

	// Eksekusi query
	rows, err = db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Println(query)

	// Iterasi setiap baris hasil query
	var files []models.FilesRedinesData
	for rows.Next() {
		var file models.FilesRedinesData
		var updatedAt sql.NullString
		var isProcess sql.NullBool

		if err := rows.Scan(&file.Yearmonth, &file.Label, &file.Policy, &file.Claim, &updatedAt, &isProcess, &file.Type, &file.SummaryProduction, &file.SummaryClaim); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Menggunakan nilai default jika updatedAt atau isProcess nil
		if updatedAt.Valid {
			file.UpdatedAt = updatedAt.String
		} else {
			file.UpdatedAt = ""
		}
		if isProcess.Valid {
			file.IsProcess = strconv.FormatBool(isProcess.Bool)
		} else {
			file.IsProcess = "false"
		}

		files = append(files, file)
	}
	// Cek apakah data files kosong
	if len(files) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data fileread",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       files,
		"perPage":     pageLength,
		"currentPage": page,
		"path":        r.URL.Path,
		"query":       queryParams,
		"fragment":    r.URL.Fragment,
		"pageName":    "page",
		"onEachSide":  3,
		"options":     map[string]string{"path": r.URL.Path, "pageName": "page"},
		"total":       totalCount,
		"lastPage":    int(math.Ceil(float64(totalCount) / float64(pageLength))),
		"status":      true,
		"message":     "Berhasil mengambil data files redines",
	})
}
