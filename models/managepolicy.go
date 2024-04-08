package models

// ini adalah stuktur managepolicy
type PolicyData struct {
	PolicyNumber            string      `json:"policy_number"`
	ProductID               interface{} `json:"product_id"`
	PackedCode              string      `json:"packed_code"`
	HargaPertanggungan      string      `json:"harga_pertanggungan"`
	ProductKey              interface{} `json:"product_key"`
	TanggalPerjanjianKredit string      `json:"tanggal_perjanjian_kredit"`
	NomorAplikasiPK         string      `json:"nomor_aplikasi_pk"`
	URLSertifikat           interface{} `json:"url_sertifikat"`
	StatusPolicy            interface{} `json:"status_policy"`
	Remark                  interface{} `json:"remark"`
	YearMonth               string      `json:"yearmonth"`
	Risk                    bool        `json:"risk"`
	NomorPolis              string      `json:"nomor_polis"`
	KodeProduk              string      `json:"kode_produk"`
	KantorCabang            string      `json:"kantor_cabang"`
	NoRekening              string      `json:"no_rekening"`
	Premi                   string      `json:"premi"`
	MulaiAsuransi           string      `json:"mulai_asuransi"`
	SelesaiAsuransi         string      `json:"selesai_asuransi"`
	TanggalLahir            string      `json:"tanggal_lahir"`
	PremiRefund             string      `json:"premi_refund"`
	Status                  string      `json:"status"`
	RemarkRefund            string      `json:"remark_refund"`
	Usia                    string      `json:"usia"`
	ContractNumber          string      `json:"contract_number"`
	SubmitDate              string      `json:"submit_date"`
	LoanStartDate           string      `json:"loan_start_date"`
	LoanAmount              string      `json:"loan_amount"`
	Rate                    string      `json:"rate"`
	Tenor                   string      `json:"tenor"`
	PremiumAmount           string      `json:"premium_amount"`
	LinkCertificate         string      `json:"link_certificate"`
	FundingPartner          string      `json:"funding_partner"`
	PolicyID                string      `json:"policy_id"`
	NomorPeminjaman         string      `json:"nomor_peminjaman"`
	TanggalAwalAkad         string      `json:"tanggal_awal_akad"`
	PokokKredit             string      `json:"pokok_kredit"`
	Borrower                string      `json:"borrower"` // Tambahkan field Borrower
	Premium                 string      `json:"premium"`  // Tambahkan field Premium
}

// type Paginator struct {
// 	Items       []PolicyData           `json:"items"`
// 	PerPage     int                    `json:"perPage"`
// 	CurrentPage int                    `json:"currentPage"`
// 	Path        string                 `json:"path"`
// 	Query       map[string]string      `json:"query"`
// 	Fragment    interface{}            `json:"fragment"`
// 	PageName    string                 `json:"pageName"`
// 	OnEachSide  int                    `json:"onEachSide"`
// 	Options     map[string]interface{} `json:"options"`
// 	Total       int                    `json:"total"`
// 	LastPage    int                    `json:"lastPage"`
// 	Periods     []Period               `json:"periods"`
// }
