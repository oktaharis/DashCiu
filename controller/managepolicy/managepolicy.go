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

func IndexPolicy(w http.ResponseWriter, r *http.Request) {
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
	app := dashboardInput["app"]
	appChild := dashboardInput["app_child"]

	// Koneksi ke database
	db := models.DBConnections[app]
	fmt.Println(db)
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}
	fmt.Println(app)
	
	// Set kolom yang akan diambil dari tabel
	var columns string
	switch app {
	case "flexi":
		columns = "policy_number, nomor_polis, kode_produk, kantor_cabang, no_rekening, premi, yearmonth, mulai_asuransi, selesai_asuransi, tanggal_lahir, url_sertifikat, risk, premi_refund, status, remark_refund, usia"
	case "kpi", "afi":
		columns = "policy_number, borrower, contract_number, submit_date, loan_start_date, loan_amount, rate, tenor, premium_amount, link_certificate, status, funding_partner"
	case "adk":
		columns = "policy_id, nomor_peminjaman, tanggal_awal_akad, tenor, tanggal_lahir, pokok_kredit, url_sertifikat, premium, rate, status, remark, yearmonth"
	}

	// Ambil nilai parameter lainnya
	length := 10
	if lenStr := dashboardInput["length"]; lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}
	fmt.Println(length)

	search := dashboardInput["search"]
	contract := dashboardInput["contract"]
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
	if app == "kpi" || app == "afi" {
		paramApp := app + "|" + appChild
		if appChild == "All" {
			paramApp = app
		}
		query = fmt.Sprintf("SELECT * FROM dashboard.sp_filter('admin', 'production|period', '%s');", paramApp)
	} else {
		query = "SELECT * FROM dashboard.sp_filter('admin', 'production|period');"
	}

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
	query = "SELECT " + columns + " FROM dashboard.policy WHERE "
	if search != "" && app != "flexi" && app != "afi" && app != "kpi" && app != "adk" {
		query += fmt.Sprintf("nomor_aplikasi_pk = '%s' AND ", search)
	}

	if search != "" && app == "flexi" {
		query += fmt.Sprintf("no_rekening = '%s' AND ", search)
	}

	if search != "" && (app == "kpi" || app == "afi") {
		query += fmt.Sprintf("policy_number = '%s' AND ", search)
	}

	if search != "" && app == "adk" {
		query += fmt.Sprintf("nomor_peminjaman = '%s' AND ", search)
	}

	if contract != "" && (app == "kpi" || app == "afi") {
		query += fmt.Sprintf("contract_number = '%s' AND ", contract)
	}

	if status != "All Status" && app != "flexi" && app != "kpi" && app != "afi" {
		query += fmt.Sprintf("status_policy = '%s' AND ", status)
	}

	if risk != "All Risk" && app != "kpi" && app != "afi" {
		query += fmt.Sprintf("risk = '%s' AND ", risk)
	}

	if status != "All Status" && (app == "kpi" || app == "afi") {
		query += fmt.Sprintf("status = '%s' AND ", status)
	}

	if appChild != "All" && (app == "kpi" || app == "afi") {
		query += fmt.Sprintf("product = '%s' AND ", appChild)
	}

	if yearmonth != "" {
		query += fmt.Sprintf("yearmonth = '%s' AND ", yearmonth)
	}

	if app == "kpi" {
		query += "funding_partner = 'PT Bank Jago Tbk' AND "
	}

	if app == "afi" {
		query += "funding_partner = 'PT ATOME FINANCE INDONESIA' AND "
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
	switch app {
	case "flexi":
		if err := rows.Scan(&policy.PolicyNumber, &policy.NomorPolis, &policy.KodeProduk, &policy.KantorCabang, &policy.NoRekening, &policy.Premi, &policy.YearMonth, &policy.MulaiAsuransi, &policy.SelesaiAsuransi, &policy.TanggalLahir, &policy.URLSertifikat, &policy.Risk, &policy.PremiRefund, &policy.Status, &policy.RemarkRefund, &policy.Usia); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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