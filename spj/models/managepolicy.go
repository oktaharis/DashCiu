package models

// ini adalah stuktur managepolicy
type PolicyData struct {
	PolicyNumber            string      `json:"policy_number"`
	PackedCode              string      `json:"packed_code"`
	Premium                 string      `json:"premium"`
	StatusPolicy            interface{} `json:"status_policy"`
	Nama                    string      `json:"nama"`
	TanggalLahir            string      `json:"tanggal_lahir"`
	TanggalMulai            string      `json:"tanggal_mulai"`
	TanggalAkhir            string      `json:"tanggal_akhir"`
	Usia                    *string     `json:"usia"`
	JmlBulanKredit          string      `json:"jml_bulan_kredit"`
	HargaPertanggungan      string      `json:"harga_pertanggungan"`
	Kategori                *string      `json:"kategori"`
	NomorRekening           string      `json:"nomor_rekening"`
	TanggalPerjanjianKredit string      `json:"tanggal_perjanjian_kredit"`
	NoKtp                   string      `json:"no_ktp"`
	NomorAplikasiPK         string      `json:"nomor_aplikasi_pk"`
	Alamat                  string      `json:"alamat"`
	CreatedAt               string      `json:"created_at"`
	UpdatedAt               string      `json:"updated_at"`
	Filename                string      `json:"filename"`
	URLSertifikat           interface{} `json:"url_sertifikat"`
	YearMonth               string      `json:"yearmonth"`
	Risk                    bool        `json:"risk"`
	ExpiredDate             *string     `json:"expired_date"`
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