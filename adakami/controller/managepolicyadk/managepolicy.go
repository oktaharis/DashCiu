package managepolicy

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"adkcontroller/helper"
	"adkcontroller/models"
)

func IndexPolicyAdk(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	dashboardInput := r.URL.Query()
	yearmonth := dashboardInput.Get("yearmonth")
	search := dashboardInput.Get("search")
	status := dashboardInput.Get("status")
	length := 10
	if lenStr := dashboardInput.Get("length"); lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}

	app := "adk"

	// Koneksi ke database
	db := models.DBConnections[app]
	fmt.Println(db)
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}
	fmt.Println(app)

	// Set kolom yang akan diambil dari tabel
	columns := []string{
		"policy_id",
		"nomor_peminjaman",
		"nomor_akad_kredit",
		"tanggal_awal_akad",
		"tanggal_mulai",
		"tanggal_akhir",
		"tenor",
		"nama_debitur",
		"tanggal_lahir",
		"policyno",
		"pokok_kredit",
		"rate",
		"premium",
		"status",
		"yearmonth",
		"created_at",
		"updated_at",
		"jenis_kelamin",
		"certificateno",
		"tanggal_cancel",
		"premi_cancel",
		"remark",
		"url_sertifikat",
		"kode_produk",
		"status_pembayaran",
		"tanggal_pembayaran",
	}

	// Query untuk mendapatkan data dari database
	query := "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
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
		query += fmt.Sprintf("nomor_peminjaman = '%s' AND ", search)
	}

	if status != "All Status" {
		query += fmt.Sprintf("status = '%s' AND ", status)
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
		if err := rows.Scan(&policy.PolicyId, &policy.NomorPeminjaman, &policy.NomorAkadKredit, &policy.TanggalAwalAkad, &policy.TanggalMulai, &policy.TanggalAkhir, &policy.Tenor, &policy.NamaDebitur, &policy.TanggalLahir, &policy.PolicyNo, &policy.PokokKredit, &policy.Rate, &policy.Premium, &policy.Status, &policy.YearMonth, &policy.CreatedAt, &policy.UpdatedAt, &policy.JenisKelamin, &policy.CertificateNo, &policy.TanggalCancel, &policy.PremiCancel, &policy.Remark, &policy.UrlSertifikat, &policy.KodeProduk, &policy.StatusPembayaran, &policy.TanggalPembayaran); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		policies = append(policies, policy)
	}
	if len(policies) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data policy",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
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
