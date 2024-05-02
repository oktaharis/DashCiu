package subrogation

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/jeypc/homecontroller/helper"
	"github.com/jeypc/homecontroller/models"
)

func SubrogationSpl(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter dari URL
	subrogationInput := r.URL.Query()
	yearmonthStr := subrogationInput.Get("yearmonth")
	search := subrogationInput.Get("search")
	status := subrogationInput.Get("status")
	lengthStr := subrogationInput.Get("length")

	// Jika lengthStr kosong, atur length ke nilai default 10
	length := 10
	if lengthStr != "" {
		var err error
		length, err = strconv.Atoi(lengthStr)
		if err != nil {
			response := map[string]interface{}{"message": "Invalid length", "status": false}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
	}

	page := 1 // Halaman default
	if pageStr := subrogationInput.Get("page"); pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	yearmonth, _ := strconv.Atoi(yearmonthStr)

	app := "spl"

	db := models.DBConnections[app]
	if db == nil {
		models.ConnectDatabase(app)
		db = models.DBConnections[app]
	}

	var items []models.TableData
	var totalCount int

	column := []string{
		"id",
		"nomor_aplikasi_pk",
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

	// Query untuk mendapatkan periods
	var periods []models.Period
	if err := db.Table("dashboard.sp_filter('admin', 'subrogation|period')").Find(&periods).Error; err != nil {
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Query untuk mendapatkan data sesuai kondisi
	query := db.Table("dashboard.subrogation")
	if search != "" {
		query = query.Where("nomor_aplikasi_pk = ?", search)
	}
	if status != "All Status" {
		query = query.Where("status = ?", status)
	}
	if search == "" && yearmonth != 0 {
		query = query.Where("yearmonth = ?", yearmonth)
	}

	// Hitung offset untuk paginasi
	offset := (page - 1) * length

	// Tambahkan limit dan offset ke dalam query
	query = query.Select(column).Limit(length).Offset(offset)

	// Mencetak query SQL untuk debug
	fmt.Println(query)

	if err := query.Find(&items).Error; err != nil {
		response := map[string]interface{}{"message": err.Error(), "status": false}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"Items":       items,
		"PerPage":     length,
		"CurrentPage": page,
		"Path":        r.URL.Path,
		"Query":       subrogationInput,
		"Fragment":    r.URL.Fragment,
		"PageName":    "page",
		"OnEachSide":  3,
		"Options": map[string]interface{}{
			"path":     r.URL.Path,
			"pageName": "page",
		},
		"Total":    totalCount,
		"LastPage": int(math.Ceil(float64(totalCount) / float64(length))),
		"Periods":  periods,
	}

	helper.ResponseJSON(w, http.StatusOK, response)
}

