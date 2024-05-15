package manageclaimafi

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"aficontroller/helper"
	"aficontroller/models"
)

func IndexClaim(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan nilai dari URL query parameters
	app := "afi"
	appChild := r.URL.Query().Get("app_child")
	lengthStr := r.URL.Query().Get("length")
	yearmonthStr := r.URL.Query().Get("yearmonth")
	search := r.URL.Query().Get("search")
	status := r.URL.Query().Get("status")

	// Konversi length dan yearmonth menjadi integer
	length, _ := strconv.Atoi(lengthStr)
	fmt.Println("pagination = ", length)
	yearmonth, _ := strconv.Atoi(yearmonthStr)

	// Koneksi ke database
	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Query untuk mendapatkan periode
	var query string
	query = fmt.Sprintf("SELECT * FROM dashboard.sp_filter('admin', 'production|period', '%s');", app)
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

	// Query untuk mendapatkan data claim
	var columns []string
	columns = []string{
		"claim_id",
		"loan_id",
		"contract_number",
		"nominal_outstanding",
		"dpd",
		"funding_partner",
		"product",
		"tenor",
		"premium_amount",
		"status",
		"remark",
		"yearmonth",
		"batch_policy",
		"load_id",
		"created_at",
	}

	// Query untuk mendapatkan data claim
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.claim WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if search != "" {
		filters = append(filters, fmt.Sprintf("loan_id = '%s'", search))
	}

	if status != "All Status" {
		filters = append(filters, fmt.Sprintf("status = '%s'", status))
	}

	if appChild != "All" {
		filters = append(filters, fmt.Sprintf("product = '%s'", appChild))
	}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}

	filters = append(filters, "funding_partner = 'PT ATOME FINANCE INDONESIA'")

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("inifilters", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.claim WHERE "
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
	var claims []models.ClaimListData
	for rows.Next() {
		var claim models.ClaimListData
		// Pindai nilai kolom ke dalam variabel struktur
		switch app {
		case "afi":
			if err := rows.Scan(&claim.ClaimId, &claim.LoanId, &claim.ContractNumber, &claim.NominalOutstanding, &claim.Dpd, &claim.FundingPartner, &claim.Product, &claim.Tenor, &claim.PremiumAmount, &claim.Status, &claim.Remark, &claim.Yearmonth, &claim.BatchPolicy, &claim.LoanId, &claim.CreatedAt); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		claims = append(claims, claim)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       claims,
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
		"message":     "Berhasil mengambil data claim",
	})
}

// Mengonversi format yearmonth dari "yyyymm" menjadi "Month YYYY"
func convertYearmonth(yearMonthStr string) string {
	year, _ := strconv.Atoi(yearMonthStr[:4])  // Ambil 4 digit pertama sebagai tahun
	month, _ := strconv.Atoi(yearMonthStr[4:]) // Ambil 2 digit terakhir sebagai bulan
	monthStr := time.Month(month).String()     // Konversi angka bulan menjadi nama bulan
	return fmt.Sprintf("%s %d", monthStr, year)
}