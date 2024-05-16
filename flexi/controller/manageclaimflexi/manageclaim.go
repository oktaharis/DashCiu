package manageclaim

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"flexicontroller/helper"
	"flexicontroller/models"
)

func ClaimFlexi(w http.ResponseWriter, r *http.Request) {
	// Mendekode body request JSON
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app := "flexi"
	// Mendapatkan nilai parameter dari URL
	claimInput := r.URL.Query()
	lengthStr := claimInput.Get("length")
	yearmonthStr := claimInput.Get("yearmonth")
	search := claimInput.Get("search")
	status := claimInput.Get("status")

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

	// Query untuk mendapatkan data claim
	columns := []string{
		"claim_id",
		"no_rekening",
		"no_polis",
		"nama",
		"usia",
		"jenis_kelamin",
		"jangka",
		"kantor_cabang",
		"tgl_pengajuan",
		"tgl_kolektibility_3",
		"nilai_pengajuan",
		"hutang_pokok",
		"nominal_disetujui",
		"rekening_koran",
		"data_nasabah",
		"status",
		"yearmonth",
		"created_at",
		"updated_at",
		"batch_policy",
		"hak_klaim_80",
		"tsi",
		"premi",
	}

	// Query untuk mendapatkan data claim
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.claim WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if search != "" {
		filters = append(filters, fmt.Sprintf("no_rekening = '%s'", search))
	}

	if status != "" {
		filters = append(filters, fmt.Sprintf("status = '%s'", status))
	}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}

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
		var usia sql.NullString // Tambahkan variabel untuk menangani kolom usia

		if err := rows.Scan(&claim.ClaimId, &claim.NoRekening, &claim.NoPolis, &claim.Nama, &claim.Usia, &claim.JenisKelamin, &claim.Jangka, &claim.KantorCabang, &claim.TglPengajuan, &claim.TglKolektibility3, &claim.NilaiPengajuan, &claim.HutangPokok, &claim.NominalDisetujui, &claim.RekeningKoran, &claim.DataNasabah, &claim.Status, &claim.Yearmonth, &claim.CreatedAt, &claim.UpdatedAt, &claim.BatchPolicy, &claim.HakKlaim80, &claim.Tsi, &claim.Premi); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Menangani nilai NULL untuk usia
		setStringIfValid := func(src sql.NullString) string {
			if src.Valid {
				return src.String
			}
			return "" // Atur nilai menjadi string kosong jika nilainya NULL
		}

		// Assign nilai usia ke variabel claim
		claim.Usia = setStringIfValid(usia)
		// Mengubah format yearmonth menjadi "Month YYYY"
		claim.Yearmonth = convertYearmonth(claim.Yearmonth)

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
