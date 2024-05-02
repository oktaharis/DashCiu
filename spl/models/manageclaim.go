package models

// ini adalah stuktur manageclaim
type ClaimListData struct {
	ClaimId                  string `json:"claim_id"`
	PolicyNumber             string `json:"policy_number"`
	ReferenceId              string `json:"reference_id"`
	ProductID                interface{} `json:"product_id"`
	PackedCode               string `json:"packed_code"`
	Premium                  string  `json:"premium"`
	PhoneNo                  string `json:"phone_no"`
	Email                    string `json:"email"`
	ApplicationNumber        string `json:"application_number"`
	Benefit                  string `json:"benefit"`
	ProductKey               interface{} `json:"product_key"`
	PackageName              string `json:"package_name"`
	PolicyStartDate          string `json:"policy_start_date"`
	StatusClaim              interface{} `json:"status_claim"`
	NoRekening			 string  `json:"no_rekening"`
	NoPerjanjianKredit       string `json:"no_perjanjian_kredit"`
	Nama                     string `json:"nama"`
	TglLahir                 interface{} `json:"tgl_lahir"`
	NoKtp                    string  `json:"no_ktp"`
	NilaiKreditDasar         string  `json:"nilai_kredit_dasar"`
	NilaiKlaim               string    `json:"nilai_klaim"`
	NilaiPokokKredit         string  `json:"nilai_pokok_kredit"`
	TglMulai                 string `json:"tgl_mulai"`
	TglAkhir                 string `json:"tgl_akhir"`
	Tenor                    string  `json:"tenor"`
	TanggalPengajuanKlaimBni string `json:"tanggal_pengajuan_klaim_bni"`
	UploadId                 string  `json:"upload_id"`
	Error                    string `json:"error"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	Total                    string  `json:"total"`
	Filename                 string `json:"filename"`
	Yearmonth                string    `json:"yearmonth"`
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
