package users

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"userscontroller/helper"
	"userscontroller/models"
)

func Users(w http.ResponseWriter, r *http.Request) {
	// Mendekode body request JSON
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mendapatkan nilai dari body request
	userInput := r.URL.Query()
	name := userInput.Get("name")
	idUser := userInput.Get("id")

	id, _ := strconv.Atoi(idUser)

	// Koneksi ke database
	models.ConnectDatabase()
	db := models.DB

	var query string

	// Query untuk mendapatkan data user
	columns := []string{
		"id", 
		"name", 
		"email",
		"email_verified_at",
		"password",
		"phone", 
		"last_login",
		"status", 
		"remember_token",
		"created_at", 
		"updated_at", 
		"role_id", 
		"product_id",
		"otp",
		"expired_at",
		"uid",
		"reset_password",
	}

	// Query untuk mendapatkan data user
	query = "SELECT " + strings.Join(columns, ", ") + " FROM dashboard.users WHERE "
	// Buat filter berdasarkan parameter yang diberikan
	filters := []string{}
	
	if name != "" {
		filters = append(filters, fmt.Sprintf("name = '%s'", name))
	}

	if id != 0 {
		filters = append(filters, fmt.Sprintf("id = '%d'", id))
	}

	// Gabungkan semua filter
	if len(filters) > 0 {
		query += strings.Join(filters, " AND ")
	} else {
		query += "1 = 1" // Tambahkan kondisi yang selalu benar jika tidak ada filter
	}
	fmt.Println("inifilters", filters)

	// Query untuk menghitung jumlah total baris yang sesuai dengan kueri
	countQuery := "SELECT COUNT(*) FROM dashboard.users WHERE "
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
	rows, err := db.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Println(query)

	// Iterasi setiap baris hasil query
	var users []models.UserData
	for rows.Next() {
		var user models.UserData
		// Pindai nilai kolom ke dalam variabel struktur
			if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.EmailVerifiedAt, &user.Password, &user.Phone, &user.LastLogin, &user.Status, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.RoleId, &user.ProductId, &user.Otp, &user.ExpiredAt, &user.Uid, &user.ResetPassword); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		users = append(users, user)
	}
	if len(users) == 0 {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "failed, get data users",
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, responseData)
		return
	}

	// Siapkan data untuk ditampilkan dalam format JSON
	// Kirim respons JSON
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"items":       users,
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
		"message":     "Berhasil mengambil data users",
	})
}
