package models

// TableData merupakan struktur untuk menyimpan data tabel tunggal
type TableData struct {
	Id              string `json:"id"`
	NomorAplikasiPk string `json:"nomor_aplikasi_pk"`
	BatchPolicy     string    `json:"batch_policy"`
	BatchClaim      string    `json:"batch_claim"`
	NilaiClaim      string    `json:"nilai_claim"`
	NilaiSubrogasi  string    `json:"nilai_subrogasi"`
	SisaTerTagih    string    `json:"sisa_tertagih"`
	SisaSubrogasi   string    `json:"sisa_subrogasi"`
	Status          string `json:"status"`
	Remark          string `json:"remark"`
	YearMonth       string    `json:"yearmonth"`
	Filename        string `json:"filename"`
	CreatedAt       string `json:"created_at"`
	Tenor           string    `json:"tenor"`
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
