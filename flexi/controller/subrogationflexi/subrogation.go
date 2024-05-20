package subrogation

import (
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

func SubrogationFlexi(w http.ResponseWriter, r *http.Request) {
	// Mendekode body request JSON
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mendapatkan nilai dari body request
	SubrogationFlexi := r.URL.Query()
	yearmonthStr := SubrogationFlexi.Get("yearmonth")
	search := SubrogationFlexi.Get("search")
	status := SubrogationFlexi.Get("status")

	yearmonth, _ := strconv.Atoi(yearmonthStr)

	// Koneksi ke database
	models.ConnectDatabase()
	db := models.DB

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

	// Query untuk mendapatkan data subrogation
	
		columns := []string{
			"id", 
			"nomor_rekening", 
			"batch_policy", 
			"batch_claim", 
			"nilai_claim", 
			"nilai_subrogasi", 
			"sisa_tertagih", 
			"sisa_subrogasi", 
			"status", 
			"remark", 
			"yearmonth", 
			"filename", 
			"created_at", 
			"tenor",
		}

	// Query untuk mendapatkan data subrogation
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.subrogation WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}

	if yearmonth != 0 {
		filters = append(filters, fmt.Sprintf("yearmonth = '%d'", yearmonth))
	}

	if search != "" {
		filters = append(filters, fmt.Sprintf("nomor_rekening = '%s'", search))
	}

	if status != "" {
		filters = append(filters, fmt.Sprintf("status = '%s'", status))
	}

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("inifilters", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.subrogation WHERE "
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
	var subrogations []models.TableData
	for rows.Next() {
		var sub models.TableData
		// Pindai nilai kolom ke dalam variabel struktur
			if err := rows.Scan(&sub.Id, &sub.NomorRekening, &sub.BatchPolicy, &sub.BatchClaim, &sub.NilaiClaim, &sub.NilaiSubrogasi, &sub.SisaTerTagih, &sub.SisaSubrogasi, &sub.Status, &sub.Remark, &sub.Yearmonth, &sub.Filename, &sub.CreatedAt, &sub.Tenor); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Mengubah format yearmonth menjadi "Month YYYY"
			sub.Yearmonth = convertYearmonth(sub.Yearmonth)

		subrogations = append(subrogations, sub)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	if len(subrogations) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data subgrogation",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
	}
	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       subrogations,
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
		"message":     "Berhasil mengambil data subrogation",
	})
}

// Mengonversi format yearmonth dari "yyyymm" menjadi "Month YYYY"
func convertYearmonth(yearMonthStr string) string {
	year, _ := strconv.Atoi(yearMonthStr[:4])  // Ambil 4 digit pertama sebagai tahun
	month, _ := strconv.Atoi(yearMonthStr[4:]) // Ambil 2 digit terakhir sebagai bulan
	monthStr := time.Month(month).String()     // Konversi angka bulan menjadi nama bulan
	return fmt.Sprintf("%s %d", monthStr, year)
}
