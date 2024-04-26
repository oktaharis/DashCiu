package models

import "time"

// TableData merupakan struktur untuk menyimpan data tabel tunggal
type TableData struct {
    NomorAplikasiPk   string    `json:"nomor_aplikasi_pk"`
    BatchPolicy       int       `json:"batch_policy"`
    Remark            *string   `json:"remark"` // Gunakan pointer untuk nullable field
    BatchClaim        int       `json:"batch_claim"`
    NilaiClaim        string    `json:"nilai_claim"`
    NilaiSubrogasi    string    `json:"nilai_subrogasi"`
    SisaTerTagih      string    `json:"sisa_tertagih"`
    Paid              string    `json:"paid"`
    Status            string    `json:"status"`
    CreatedAt         time.Time `json:"created_at"`
    CreatedAtFormatted string  `json:"created_at_formatted"` // Properti baru untuk format tanggal CreatedAt
}


// Paginator merupakan struktur untuk menyimpan data paginasi
type Paginator struct {
	Items       []TableData            `json:"items"`
	PerPage     int                    `json:"perPage"`
	CurrentPage int                    `json:"currentPage"`
	Path        string                 `json:"path"`
	Query       map[string]string      `json:"query"`
	Fragment    interface{}            `json:"fragment"`
	PageName    string                 `json:"pageName"`
	OnEachSide  int                    `json:"onEachSide"`
	Options     map[string]interface{} `json:"options"`
	Total       int                    `json:"total"`
	LastPage    int                    `json:"lastPage"`
	Periods     []Period               `json:"periods"`
}
