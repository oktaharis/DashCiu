package filesredines

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	// Mendekode body request JSON
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mendapatkan nilai dari body request
	app := requestBody["app"]
	yearmonthStr := requestBody["yearmonth"]
	appChild := requestBody["app_child"]
	
	yearmonth, _ := strconv.Atoi(yearmonthStr)

	// Koneksi ke database
	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Query untuk mendapatkan periode
	var query string
	if app == "kpi" || app == "afi" {
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_filter('admin', 'production|period', '%s');", app)
		fmt.Println(query)
	} else {
		query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
		fmt.Println(query)
	}

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
	var columns []string
	switch app {
	case "kpi", "afi":
		columns = []string{
			"yearmonth",
			"label",
			"policy",
			"claim",
			"updated_at",
			"is_process",
			"type",
			"summary_production",
			"summary_claim",
		}
	default:
		columns = []string{
			"yearmonth",
			"label",
			"policy",
			"claim",
			"updated_at",
			"is_process",
		}
	}


	// Query untuk mendapatkan data files redines
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.files_redines WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}
	
	if appChild != "All" && (app == "kpi" || app == "afi") {
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
		// Pindai nilai kolom ke dalam variabel struktur
		switch app {
		case "kpi", "afi":
			if err := rows.Scan(&file.Yearmonth, &file.Label, &file.Policy, &file.Claim, &file.UpdatedAt, &file.IsProcess, &file.Type, &file.SummaryProduction, &file.SummaryClaim); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			if err := rows.Scan(&file.Yearmonth, &file.Label, &file.Policy, &file.Claim, &file.UpdatedAt, &file.IsProcess); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Mengubah format yearmonth menjadi "Month YYYY"
			file.Yearmonth = convertYearmonth(file.Yearmonth)
		}

		files = append(files, file)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       files,
		"perPage":     pageLength,
		"currentPage": page,
		"path":        r.URL.Path,
		"query":       r.URL.Query(),
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

// Mengonversi format yearmonth dari "yyyymm" menjadi "Month YYYY"
func convertYearmonth(yearMonthStr string) string {
	year, _ := strconv.Atoi(yearMonthStr[:4])  // Ambil 4 digit pertama sebagai tahun
	month, _ := strconv.Atoi(yearMonthStr[4:]) // Ambil 2 digit terakhir sebagai bulan
	monthStr := time.Month(month).String()     // Konversi angka bulan menjadi nama bulan
	return fmt.Sprintf("%s %d", monthStr, year)
}