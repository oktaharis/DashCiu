package models

// ini adalah stuktur managepolicy
type PolicyData struct {
	PolicyNumber       string `json:"policy_number"`
	Periode            string `json:"periode"`
	KantorCabang       *string `json:"kantor_cabang"`
	NoRekening         *string `json:"no_rekening"`
	NoKTP              *string `json:"no_ktp"`
	CIF                *string `json:"cif"`
	NamaDebitur        *string `json:"nama_debitur"`
	TanggalLahir       *string `json:"tanggal_lahir"`
	JenisKelamin       *string `json:"jenis_kelamin"`
	Produk             *string `json:"produk"`
	KodeProduk         *string `json:"kode_produk"`
	SubProduk          *string `json:"sub_produk"`
	ProdukFintech      *string `json:"produk_fintech"`
	Kategori           *string `json:"kategori"`
	NamaPerusahaan     *string `json:"nama_perusahaan"`
	MulaiAsuransi      *string `json:"mulai_asuransi"`
	SelesaiAsuransi    *string `json:"selesai_asuransi"`
	JangkaWaktu        *string `json:"jangka_waktu"`
	LimitPlafond       *string `json:"limit_plafond"`
	NilaiPertanggungan *string `json:"nilai_pertanggungan"`
	RatePremi          *string `json:"rate_premi"`
	Premi              *string `json:"premi"`
	TglPencairan       *string `json:"tgl_pencairan"`
	TglPK              *string `json:"tgl_pk"`
	NoPK               *string `json:"no_pk"`
	NamaProgram        *string `json:"nama_program"`
	IsCBC              *string `json:"is_cbc"`
	Coverage           *string `json:"coverage"`
	NomorPolis         *string `json:"nomor_polis"`
	URLSertifikat      *string `json:"url_sertifikat"`
	YearMonth          string `json:"yearmonth"`
	CreatedAt          *string `json:"created_at"`
	Risk               *string `json:"risk"`
	Status             string `json:"status"`
	PSJT               *string `json:"psjt"`
	SisaBulan          *string `json:"sisa_bulan"`
	PremiRefund        *string `json:"premi_refund"`
	RemarkRefund       *string `json:"remark_refund"`
	ExpiredDate        *string `json:"expired_date"`
}

type PaginatorPolicy struct {
	Items       []PolicyData           `json:"items"`
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
