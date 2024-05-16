package models

// ini adalah stuktur managepolicy
type PolicyData struct {
	PolicyId          string `json:"policy_id"`
	NomorPeminjaman   string `json:"nomor_peminjaman"`
	NomorAkadKredit   string `json:"nomor_akad_kredit"`
	TanggalAwalAkad   string `json:"tanggal_awal_akad"`
	TanggalMulai      string `json:"tanggal_mulai"`
	TanggalAkhir      *string `json:"tanggal_akhir"`
	Tenor             string `json:"tenor"`
	NamaDebitur       string `json:"nama_debitur"`
	TanggalLahir      *string `json:"tanggal_lahir"`
	PolicyNo          *string `json:"policy_no"`
	PokokKredit       string `json:"pokok_kredit"`
	Rate              string `json:"rate"`
	Premium           string `json:"premium"`
	Status            string `json:"status"`
	YearMonth         string `json:"yearmonth"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	JenisKelamin      *string `json:"jenis_kelamin"`
	CertificateNo     *string `json:"certificateno"`
	TanggalCancel     *string `json:"tanggal_cancel"`
	PremiCancel       *string `json:"premi_cancel"`
	Remark            *string `json:"remark"`
	UrlSertifikat     *string `json:"url_sertifikat"`
	KodeProduk        *string `json:"kode_produk"`
	StatusPembayaran  *string `json:"status_pembayaran"`
	TanggalPembayaran *string `json:"tanggal_pembayaran"`
}

// type Paginator struct {
// 	Items       []PolicyData           json:"items"
// 	PerPage     int                    json:"perPage"
// 	CurrentPage int                    json:"currentPage"
// 	Path        string                 json:"path"
// 	Query       map[string]string      json:"query"
// 	Fragment    interface{}            json:"fragment"
// 	PageName    string                 json:"pageName"
// 	OnEachSide  int                    json:"onEachSide"
// 	Options     map[string]interface{} json:"options"
// 	Total       int                    json:"total"
// 	LastPage    int                    json:"lastPage"
// 	Periods     []Period               json:"periods"
// }
