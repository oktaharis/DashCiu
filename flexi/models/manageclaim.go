package models

// ini adalah stuktur manageclaim
type ClaimListData struct {
	ClaimId                  string `json:"claim_id"`
	PolicyNumber             string `json:"policy_number"`
	ReferenceId              string `json:"reference_id"`
	ProductID                interface{} `json:"product_id"`
	PackedCode               string `json:"packed_code"`
	Premium                  string  `json:"premium"`
	Premi 	            	 string  `json:"premi"`
	LoanId					 string  `json:"loan_id"`
	PhoneNo                  string `json:"phone_no"`
	Email                    string `json:"email"`
	Jangka					string    `json:"jangka"`
	NoSurat					string    `json:"no_surat"`
	TglKolektibility3       string    `json:"tgl_kolektibility_3"`
	PenyebabKlaim           string    `json:"penyebab_klaim"`
	NilaiPengajuan			string    `json:"nilai_pengajuan"`
	HutangPokok         	string    `json:"hutang_pokok"`
	NominalDisetujui		string    `json:"nominal_disetujui"`
	RekeningKoran 			string    `json:"rekening_koran"`
	DataNasabah  	        string     `json:"data_nasabah"`
	JenisKelamin			string  `json:"jenis_kelamin"`
	NomorPeminjaman         string      `json:"nomor_peminjaman"`
	ApplicationNumber        string `json:"application_number"`
	Benefit                  string `json:"benefit"`
	PremiumAmount           string      `json:"premium_amount"`
	KantorCabang			 string  `json:"kantor_cabang"`
	ProductKey               interface{} `json:"product_key"`
	PackageName              string `json:"package_name"`
	PolicyStartDate          string `json:"policy_start_date"`
	StatusClaim              interface{} `json:"status_claim"`
	Status                   string      `json:"status"`
	ContractNumber          string      `json:"contract_number"`
	NoRekening               string `json:"no_rekening"`
	LoanAmount              string      `json:"loan_amount"`
	NoPolis					 string      `json:"no_polis"`
	Usia					 string       `json:"usia"`
	NoPerjanjianKredit       string `json:"no_perjanjian_kredit"`
	Nama                     string `json:"nama"`
	FundingPartner          string      `json:"funding_partner"`
	TglLahir                 interface{} `json:"tgl_lahir"`
	NoKtp                    string  `json:"no_ktp"`
	NilaiKreditDasar         string  `json:"nilai_kredit_dasar"`
	NilaiKlaim               string    `json:"nilai_klaim"`
	HakKlaim80				string  `json:"hak_klaim_80"`
	HakHutangPokok			string   `json:"hak_hutang_pokok"`
	Tsi						string  `json:"tsi"`
	NilaiPokokKredit         string  `json:"nilai_pokok_kredit"`
	TglMulai                 string `json:"tgl_mulai"`
	TglAkhir                 string `json:"tgl_akhir"`
	Tenor                    string  `json:"tenor"`
	TanggalPengajuanKlaimBni string `json:"tanggal_pengajuan_klaim_bni"`
	UploadId                 string  `json:"upload_id"`
	NominalOutstanding		string  `json:"nominal_outstanding"`
	TglPengajuan			 string  `json:"tgl_pengajuan"`
	Error                    string `json:"error"`
	PolicyId				string  `json:"policy_id"`
	TanggalClaim			string   `json:"tanggal_claim"`
	JumlahKlaim				string    `json:"jumlah_klaim"`
	Product                    string `json:"product"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	Total                    string  `json:"total"`
	Dpd                      string  `json:"dpd"`
	Filename                 string `json:"filename"`
	Yearmonth                string    `json:"yearmonth"`
	PokokKredit              string      `json:"pokok_kredit"`
	Remark                   interface{} `json:"remark"`
	BatchPolicy              string  `json:"batch_policy"`
	Category                 string `json:"category"`
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
