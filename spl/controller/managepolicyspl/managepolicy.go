package managepolicy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

func PolicySpl(w http.ResponseWriter, r *http.Request) {
	// Dekode input JSON dari request body
	var dashboardInput map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dashboardInput); err != nil {
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ambil nilai parameter dari input JSON
	yearmonth := dashboardInput["yearmonth"]
	app := "spl"

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

	// Ambil nilai parameter lainnya
	length := 10
	if lenStr := dashboardInput["length"]; lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}
	fmt.Println(length)

	search := dashboardInput["search"]
	status := "All Status"
	if statusStr := dashboardInput["status"]; statusStr != "" {
		status = statusStr
	}

	risk := "All Risk"
	if riskStr := dashboardInput["risk"]; riskStr != "" {
		risk = riskStr
	}

	// Query untuk mendapatkan data dari database
	var query string

	query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"

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

	if search != "" && app != "flexi" && app != "afi" && app != "kpi" && app != "adk" {
		query += fmt.Sprintf("nomor_aplikasi_pk = '%s' AND ", search)
	}

	if status != "All Status" && app != "flexi" && app != "kpi" && app != "afi" {
		query += fmt.Sprintf("status_policy = '%s' AND ", status)
	}

	if risk != "All Risk" && app != "kpi" && app != "afi" {
		query += fmt.Sprintf("risk = '%s' AND ", risk)
	}

	if yearmonth != "" {
		query += fmt.Sprintf("yearmonth = '%s' AND ", yearmonth)
	}

	// Hapus "AND" terakhir dari query
	query = strings.TrimSuffix(query, "AND ")

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
		if err := rows.Scan(&policy.PolicyNumber, &policy.PackedCode, &policy.Premium, &policy.StatusPolicy, &policy.Nama, &policy.TanggalLahir, &policy.TanggalMulai, &policy.TanggalAkhir, &policy.Usia, &policy.JmlBulanKredit, &policy.HargaPertanggungan, &policy.Kategori, &policy.NomorRekening, &policy.TanggalPerjanjianKredit, &policy.NoKtp, &policy.NomorAplikasiPK, &policy.Alamat, &policy.CreatedAt, &policy.UpdatedAt, &policy.Filename, &policy.URLSertifikat, &policy.YearMonth, &policy.Risk, &policy.ExpiredDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
