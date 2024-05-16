package managepolicy

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"splcontroller/helper"
	"splcontroller/models"
)

func PolicySpl(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	dashboardInput := r.URL.Query()
	yearmonth := dashboardInput.Get("yearmonth")
	search := dashboardInput.Get("search")
	status := dashboardInput.Get("status")
	risk := dashboardInput.Get("risk")
	length := 10
	if lenStr := dashboardInput.Get("length"); lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}
// Koneksi ke database
models.ConnectDatabase()
db := models.DB

	// Set kolom yang akan diambil dari tabel
	columns := []string{
		"policy_number",
		"packed_code",
		"premium",
		"status_policy",
		"nama",
		"tanggal_lahir",
		"tanggal_mulai",
		"tanggal_akhir",
		"usia",
		"jml_bulan_kredit",
		"harga_pertanggungan",
		"kategori",
		"nomor_rekening",
		"tanggal_perjanjian_kredit",
		"no_ktp",
		"nomor_aplikasi_pk",
		"alamat",
		"created_at",
		"updated_at",
		"filename",
		"url_sertifikat",
		"yearmonth",
		"risk",
		"expired_date",
	}

	// Query untuk mendapatkan data dari database
	var query string

	query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
	fmt.Println("Query 1:", query)

	// Eksekusi query
	rows, err := db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Ambil periode dan labelnya
	var periods []models.Period
	selectedPeriod := ""

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

		periods = append(periods, period)
	}

	// Set parameter periode default jika tidak disediakan
	if yearmonth == "" && len(periods) > 0 {
		yearmonth = periods[len(periods)-1].YearMonth
	}

	// Query untuk mendapatkan data policy dari database
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.policy WHERE "

	if search != "" {
		query += fmt.Sprintf("nomor_aplikasi_pk = '%s' AND ", search)
	}

	if status != "All Status" {
		query += fmt.Sprintf("status_policy = '%s' AND ", status)
	}

	if risk != "All Risk" {
		query += fmt.Sprintf("risk = '%s' AND ", risk)
	}

	if yearmonth != "" {
		query += fmt.Sprintf("yearmonth = '%s' AND ", yearmonth)
	}

	// Hapus "AND" terakhir dari query
	query = strings.TrimSuffix(query, "AND ")

	// Tambahkan batasan jumlah data yang diambil
	query += fmt.Sprintf("LIMIT %d", length)
	fmt.Println("Query 2:", query)

	// Eksekusi query
	rows, err = db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var policies []models.PolicyData

	// Iterasi setiap baris hasil query
	for rows.Next() {
		var policy models.PolicyData
		// Pindai nilai kolom ke dalam variabel struktur
		var usia, expiredDate sql.NullString
		if err := rows.Scan(&policy.PolicyNumber, &policy.PackedCode, &policy.Premium, &policy.StatusPolicy, &policy.Nama, &policy.TanggalLahir, &policy.TanggalMulai, &policy.TanggalAkhir, &usia, &policy.JmlBulanKredit, &policy.HargaPertanggungan, &policy.Kategori, &policy.NomorRekening, &policy.TanggalPerjanjianKredit, &policy.NoKtp, &policy.NomorAplikasiPK, &policy.Alamat, &policy.CreatedAt, &policy.UpdatedAt, &policy.Filename, &policy.URLSertifikat, &policy.YearMonth, &policy.Risk, &expiredDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if usia.Valid {
			usiaStr := usia.String
			policy.Usia = &usiaStr
		} else {
			policy.Usia = nil
		}
		if expiredDate.Valid {
			expiredDateStr := expiredDate.String
			policy.ExpiredDate = &expiredDateStr
		} else {
			policy.ExpiredDate = nil
		}

		policies = append(policies, policy)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	responseData := map[string]interface{}{
		"data":            policies,
		"periods":         periods,
		"selected_period": selectedPeriod,
		"status":          true,
		"message":         "Berhasil mengambil data policy",
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)
}


