package managepolicy

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"flexicontroller/helper"
	"flexicontroller/models"
)

func PolicyFlexi(w http.ResponseWriter, r *http.Request) {
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
		"periode",
		"kantor_cabang",
		"no_rekening",
		"no_ktp",
		"cif",
		"nama_debitur",
		"tanggal_lahir",
		"jenis_kelamin",
		"produk",
		"kode_produk",
		"sub_produk",
		"produk_fintech",
		"kategori",
		"nama_perusahaan",
		"mulai_asuransi",
		"selesai_asuransi",
		"jangka_waktu",
		"limit_plafond",
		"nilai_pertanggungan",
		"rate_premi",
		"premi",
		"tgl_pencairan",
		"tgl_pk",
		"no_pk",
		"nama_program",
		"is_cbc",
		"coverage",
		"nomor_polis",
		"url_sertifikat",
		"yearmonth",
		"created_at",
		"risk",
		"status",
		"psjt",
		"sisa_bulan",
		"premi_refund",
		"remark_refund",
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
		query += fmt.Sprintf("no_rekening = '%s' AND ", search)
	}

	if status != "" {
		query += fmt.Sprintf("status = '%s' AND ", status)
	}

	if risk != "" {
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
		if err := rows.Scan(&policy.PolicyNumber, &policy.Periode, &policy.KantorCabang, &policy.NoRekening, &policy.NoKTP, &policy.CIF, &policy.NamaDebitur, &policy.TanggalLahir, &policy.JenisKelamin, &policy.Produk, &policy.KodeProduk, &policy.SubProduk, &policy.ProdukFintech, &policy.Kategori, &policy.NamaPerusahaan, &policy.MulaiAsuransi, &policy.SelesaiAsuransi, &policy.JangkaWaktu, &policy.LimitPlafond, &policy.NilaiPertanggungan, &policy.RatePremi, &policy.Premi, &policy.TglPencairan, &policy.TglPK, &policy.NoPK, &policy.NamaProgram, &policy.IsCBC, &policy.Coverage, &policy.NomorPolis, &policy.URLSertifikat, &policy.YearMonth, &policy.CreatedAt, &policy.Risk, &policy.Status, &policy.PSJT, &policy.SisaBulan, &policy.PremiRefund, &policy.RemarkRefund, &policy.ExpiredDate,
			); err != nil {
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


