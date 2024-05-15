package models

// ini adalah stuktur manageclaim
type ClaimListData struct {
	ClaimId           string      `json:"claim_id"`
	NoRekening        string      `json:"no_rekening"`
	NoPolis           string      `json:"no_polis"`
	Nama              string      `json:"nama"`
	Usia              string      `json:"usia"`
	JenisKelamin      string      `json:"jenis_kelamin"`
	Kategori          string      `json:"kategori"`
	Jangka            string      `json:"jangka"`
	KantorCabang      string      `json:"kantor_cabang"`
	TglPengajuan      string      `json:"tgl_pengajuan"`
	NoSurat           string      `json:"no_surat"`
	TglKolektibility3 string      `json:"tgl_kolektibility_3"`
	PenyebabKlaim     string      `json:"penyebab_klaim"`
	NilaiPengajuan    string      `json:"nilai_pengajuan"`
	HutangPokok       string      `json:"hutang_pokok"`
	TunggakanBunga    string      `json:"tunggakan_bunga"`
	TunggakanBiaya    string      `json:"tunggakan_biaya"`
	TunggakanDenda    string      `json:"tunggakan_denda"`
	NominalDisetujui  string      `json:"nominal_disetujui"`
	RekeningKoran     string      `json:"rekening_koran"`
	BuktiDokumen      string      `json:"bukti_dokumen"`
	DataNasabah       string      `json:"data_nasabah"`
	PembayaranKlaim   string      `json:"pembayaran_klaim"`
	Remark            interface{} `json:"remark"`
	Status            string      `json:"status"`
	ClaimSettlement   string      `json:"claim_settlement"`
	Yearmonth         string      `json:"yearmonth"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	BuktiPembayaran   string      `json:"bukti_pembayaran"`
	BatchPolicy       string      `json:"batch_policy"`
	HakKlaim80        string      `json:"hak_klaim_80"`
	HakHutangPokok    string      `json:"hak_hutang_pokok"`
	Tsi               string      `json:"tsi"`
	Premi             string      `json:"premi"`
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
