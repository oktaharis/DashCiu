package managepolicy

import (
	"database/sql"
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

func IndexPolicy(w http.ResponseWriter, r *http.Request) {
	// Mendekode body request JSON
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mendapatkan nilai dari body request
	app := requestBody["app"]
	appChild := requestBody["app_child"]
	lengthStr := requestBody["length"]
	yearmonthStr := requestBody["yearmonth"]
	search := requestBody["search"]
	contract := requestBody["contract"]
	status := requestBody["status"]
	risk := requestBody["risk"]

	// Konversi length dan yearmonth menjadi integer
	length, _ := strconv.Atoi(lengthStr)
	fmt.Println("pagination = ",length)
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

	// Query untuk mendapatkan data policy
	var columns []string
	switch app {
	case "flexi":
		columns = []string{
			"policy_number",
			"nomor_polis",
			"kode_produk",
			"kantor_cabang",
			"no_rekening",
			"premi",
			"yearmonth",
			"mulai_asuransi",
			"selesai_asuransi",
			"tanggal_lahir",
			"url_sertifikat",
			"risk",
			"premi_refund",
			"status",
			"remark_refund as remark",
			"DATE_PART('YEAR', AGE(now(), tanggal_lahir )) as usia",
		}
	case "kpi", "afi":
		columns = []string{
			"policy_number",
			"borrower",
			"contract_number",
			"submit_date",
			"loan_start_date",
			"loan_amount",
			"rate",
			"tenor",
			"premium_amount",
			"link_certificate",
			"status",
			"funding_partner",
		}
	case "adk":
		columns = []string{
			"policy_id",
			"nomor_peminjaman",
			"tanggal_awal_akad",
			"tenor",
			"tanggal_lahir",
			"pokok_kredit",
			"url_sertifikat",
			"premium",
			"rate",
			"status",
			"remark",
			"yearmonth",
		}
	default:
		columns = []string{
			"policy_number",
			"product_id",
			"packed_code",
			"harga_pertanggungan",
			"product_key",
			"tanggal_perjanjian_kredit",
			"nomor_aplikasi_pk",
			"url_sertifikat",
			"status_policy",
			"remark",
			"yearmonth",
			"risk",
		}
	}

	// Query untuk mendapatkan data policy
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.policy WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if search != "" && app != "flexi" && app != "afi" && app != "kpi" && app != "adk" {
		filters = append(filters, fmt.Sprintf("nomor_aplikasi_pk = '%s'", search))
	}

	if search != "" && app == "flexi" {
		filters = append(filters, fmt.Sprintf("no_rekening = '%s'", search))
	}

	if search != "" && (app == "kpi" || app == "afi") {
		filters = append(filters, fmt.Sprintf("policy_number = '%s'", search))
	}

	if search != "" && app == "adk" {
		filters = append(filters, fmt.Sprintf("nomor_peminjaman = '%s'", search))
	}

	if contract != "" && (app == "kpi" || app == "afi") {
		filters = append(filters, fmt.Sprintf("contract_number = '%s'", contract))
	}

	if status != "All Status" && app != "flexi" && app != "kpi" && app != "afi" {
		filters = append(filters, fmt.Sprintf("status_policy = '%s'", status))
	}

	if risk != "All Risk" && app != "kpi" && app != "afi" {
		filters = append(filters, fmt.Sprintf("risk = '%s'", risk))
	}

	if status != "All Status" && (app == "kpi" || app == "afi") {
		filters = append(filters, fmt.Sprintf("status = '%s'", status))
	}

	if appChild != "All" && (app == "kpi" || app == "afi") {
		filters = append(filters, fmt.Sprintf("product = '%s'", appChild))
	}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}

	if app == "kpi" {
		filters = append(filters, "funding_partner = 'PT Bank Jago Tbk'")
	}

	if app == "afi" {
		filters = append(filters, "funding_partner = 'PT ATOME FINANCE INDONESIA'")
	}

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("ini filters =", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.policy WHERE "
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
	var policies []models.PolicyData
	for rows.Next() {
		var policy models.PolicyData
		// Pindai nilai kolom ke dalam variabel struktur
		switch app {
		case "flexi":
			var premiRefund, remarkRefund, usia sql.NullString // Tambahkan variabel untuk menangani kolom remark_refund dan usia
			if err := rows.Scan(&policy.PolicyNumber, &policy.NomorPolis, &policy.KodeProduk, &policy.KantorCabang, &policy.NoRekening, &policy.Premi, &policy.YearMonth, &policy.MulaiAsuransi, &policy.SelesaiAsuransi, &policy.TanggalLahir, &policy.URLSertifikat, &policy.Risk, &premiRefund, &policy.Status, &remarkRefund, &usia); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Menangani nilai NULL untuk premiRefund, remarkRefund, dan usia
			setStringIfValid := func(src sql.NullString) string {
				if src.Valid {
					return src.String
				}
				return "" // Atur nilai menjadi string kosong jika nilainya NULL
			}

			// Assign nilai premiRefund, remarkRefund, dan usia ke variabel policy
			policy.PremiRefund = setStringIfValid(premiRefund)
			policy.RemarkRefund = setStringIfValid(remarkRefund)
			policy.Usia = setStringIfValid(usia)

		case "kpi", "afi":
			if err := rows.Scan(&policy.PolicyNumber, &policy.Borrower, &policy.ContractNumber, &policy.SubmitDate, &policy.LoanStartDate, &policy.LoanAmount, &policy.Rate, &policy.Tenor, &policy.PremiumAmount, &policy.LinkCertificate, &policy.Status, &policy.FundingPartner); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		case "adk":
			if err := rows.Scan(&policy.PolicyID, &policy.NomorPeminjaman, &policy.TanggalAwalAkad, &policy.Tenor, &policy.TanggalLahir, &policy.PokokKredit, &policy.URLSertifikat, &policy.Premium, &policy.Rate, &policy.Status, &policy.Remark, &policy.YearMonth); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			var hargaPertanggungan float64
			if err := rows.Scan(&policy.PolicyNumber, &policy.ProductID, &policy.PackedCode, &hargaPertanggungan, &policy.ProductKey, &policy.TanggalPerjanjianKredit, &policy.NomorAplikasiPK, &policy.URLSertifikat, &policy.StatusPolicy, &policy.Remark, &policy.YearMonth, &policy.Risk); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Mengonversi nilai harga_pertanggungan menjadi format yang mudah dibaca
			policy.HargaPertanggungan = fmt.Sprintf("%.0f", hargaPertanggungan)
			// Mengubah format yearmonth menjadi "Month YYYY"
			policy.YearMonth = convertYearMonth(policy.YearMonth)
		}

		policies = append(policies, policy)
	}
// Inisialisasi variabel untuk menyimpan nilai yearmonth dari item pertama (jika tersedia)
var itemYearMonth string

// Pastikan ada setidaknya satu item dalam slice policies
if len(policies) > 0 {
    // Ambil yearmonth dari item pertama dalam slice policies
    itemYearMonth = policies[0].YearMonth
}

// Siapkan data query
queryParams := map[string]string{
    "yearmonth": itemYearMonth,
    "search":    search,
    "length":    lengthStr,
    "risk":      risk,
}

// Kirim respons JSON
helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
    "items":       policies,
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
    "message":     "Berhasil mengambil data policy",
})

}

// Mengonversi format yearmonth dari "yyyymm" menjadi "Month YYYY"
func convertYearMonth(yearMonthStr string) string {
	year, _ := strconv.Atoi(yearMonthStr[:4])  // Ambil 4 digit pertama sebagai tahun
	month, _ := strconv.Atoi(yearMonthStr[4:]) // Ambil 2 digit terakhir sebagai bulan
	monthStr := time.Month(month).String()     // Konversi angka bulan menjadi nama bulan
	return fmt.Sprintf("%s %d", monthStr, year)
}
