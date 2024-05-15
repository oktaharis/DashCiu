package managepolicyafi

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"aficontroller/helper"
	"aficontroller/models"
)

func IndexPolicyAfi(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	dashboardInput := r.URL.Query()
	yearmonth := dashboardInput.Get("yearmonth")
	search := dashboardInput.Get("search")
	status := dashboardInput.Get("status")
	risk := dashboardInput.Get("risk")
	appChild := dashboardInput.Get("app_child")

	length := 10
	if lenStr := dashboardInput.Get("length"); lenStr != "" {
		length, _ = strconv.Atoi(lenStr)
	}

	app := "afi"

	// Koneksi ke database
	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	// Set kolom yang akan diambil dari tabel
	columns := []string{
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

	// Query untuk mendapatkan data dari database
	query := "SELECT " + strings.Join(columns, ",") + " FROM dashboard.policy WHERE 1 = 1"

	// Tambahkan kondisi pencarian ke query
	var conditions []string
	if yearmonth != "" {
		conditions = append(conditions, fmt.Sprintf("yearmonth = '%s'", yearmonth))
	}
	if search != "" {
		conditions = append(conditions, fmt.Sprintf("policy_number = '%s'", search))
	}
	if status != "All Status" {
		conditions = append(conditions, fmt.Sprintf("status = '%s'", status))
	}
	if risk == "" {
		risk = "All Risk"
	}
	if risk != "All Risk" {
		conditions = append(conditions, fmt.Sprintf("risk = '%s'", risk))
	}
	if appChild != "All" {
		conditions = append(conditions, fmt.Sprintf("product = '%s'", appChild))
	}

	// Gabungkan semua kondisi ke dalam satu query
	// if len(conditions) > 0 {
	// 	query += " AND " + strings.Join(conditions, " AND ")
	// }

	// Tambahkan LIMIT
	query += fmt.Sprintf(" LIMIT %d OFFSET 0", length)

	// Hapus "AND" terakhir dari query
	query = strings.TrimSuffix(query, "AND ")

	fmt.Println("Query:", query) // Print the query for debugging

	// Eksekusi query
	rows, err := db.Raw(query).Rows()
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
		if err := rows.Scan(
			&policy.PolicyNumber,
			&policy.Borrower,
			&policy.ContractNumber,
			&policy.SubmitDate,
			&policy.LoanStartDate,
			&policy.LoanAmount,
			&policy.Rate,
			&policy.Tenor,
			&policy.PremiumAmount,
			&policy.LinkCertificate,
			&policy.Status,
			&policy.FundingPartner,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		policies = append(policies, policy)
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	// Jika tidak ada data yang ditemukan
	if len(policies) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "Tidak ada data yang ditemukan",
		}
		helper.ResponseJSON(w, http.StatusOK, responseData)
		return
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	responseData := map[string]interface{}{
		"data":    policies,
		"status":  true,
		"message": "Berhasil mengambil data policy",
	}

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)

	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, responseData)
}
