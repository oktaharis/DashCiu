package models

// ini adalah stuktur manageclaim
type ClaimListData struct {
	ClaimId         string  `json:"claim_id"`
	PolicyId        string  `json:"policy_id"`
	NomorPeminjaman string  `json:"nomor_peminjaman"`
	TanggalKlaim    string  `json:"tanggal_klaim"`
	JumlahKlaim     string  `json:"jumlah_klaim"`
	PokokKredit     string  `json:"pokok_kredit"`
	Status          string  `json:"status"`
	Remark          *string `json:"remark"`
	Yearmonth       string  `json:"yearmonth"`
	BatchPolicy     string  `json:"batch_policy"`
	CreatedAt       string  `json:"created_at"`
}

type PaginatorClaim struct {
	Items       []ClaimListData        `json:"items"`
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
