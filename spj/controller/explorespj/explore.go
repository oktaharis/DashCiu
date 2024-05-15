package explore

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"splcontroller/helper"
	"splcontroller/models"
)

func ExploreSpj(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	queryValues := r.URL.Query()
	yearmonthStr := queryValues.Get("yearmonth")
	nik := queryValues.Get("nik")
	name := queryValues.Get("name")
	noPk := queryValues.Get("no_pk")
	lengthStr := queryValues.Get("length")

	// Konversi length dan yearmonth menjadi integer
	length, _ := strconv.Atoi(lengthStr)
	fmt.Println("pagination = ", length)
	yearmonth, _ := strconv.Atoi(yearmonthStr)

	app := "spj"

	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Query untuk mendapatkan periode
	query := "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
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

	columns := []string{
		"no_pk",
		"nik",
		"name",
		"policy",
		"claim",
		"policy_yearmonth",
		"ltc_by_nik",
	}

	// Query untuk mendapatkan data explore
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.summary_explore WHERE "

	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if nik != "" {
		filters = append(filters, fmt.Sprintf("nik = '%s'", nik))
	}

	if noPk != "" {
		filters = append(filters, fmt.Sprintf("no_pk = '%s'", noPk))
	}

	if name != "" {
		filters = append(filters, fmt.Sprintf("name = '%s'", name))
	}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("policy_yearmonth = %d", yearmonth))
	}

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("filters:", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.summary_explore WHERE "

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
	if lengthStr != "" {
		pageLength, _ = strconv.Atoi(lengthStr)
	}

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
	var explores []models.ExploreData
	for rows.Next() {
		var explore models.ExploreData
		// Pindai nilai kolom ke dalam variabel struktur
		if err := rows.Scan(&explore.NoPk, &explore.Nik, &explore.Name, &explore.Policy, &explore.Claim, &explore.PolicyYearmonth, &explore.LtcByNik); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		explore.PolicyYearmonth = convertYearmonthPolicy(explore.PolicyYearmonth)
		explore.ClaimYearmonth = convertYearmonthClaim(explore.ClaimYearmonth)

		explores = append(explores, explore)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       explores,
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
		"message":     "Berhasil mengambil data explore",
	})
}

// Mengonversi format yearmonth dari "yyyymm" menjadi "Month YYYY"
func convertYearmonthPolicy(policyYearmonth string) string {
	if policyYearmonth == "" {
        return "" // Jika policyYearmonth kosong, kembalikan nilai kosong
    }
    year, _ := strconv.Atoi(policyYearmonth[:4])
    month, _ := strconv.Atoi(policyYearmonth[4:])
    monthStr := time.Month(month).String()
    return fmt.Sprintf("%s %d", monthStr, year)
}


func convertYearmonthClaim(claimYearmonth string) string {
	if claimYearmonth == "" {
        return "" // Jika claimYearmonth kosong, kembalikan nilai kosong
    }
	year, _ := strconv.Atoi(claimYearmonth[:4])  // Ambil 4 digit pertama sebagai tahun
	month, _ := strconv.Atoi(claimYearmonth[4:]) // Ambil 2 digit terakhir sebagai bulan
	monthStr := time.Month(month).String()     // Konversi angka bulan menjadi nama bulan
	return fmt.Sprintf("%s %d", monthStr, year)
}
