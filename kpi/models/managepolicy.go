package models

// ini adalah stuktur managepolicy
type PolicyData struct {
	PolicyNumber    string  `json:"policy_number"`
	Borrower        string  `json:"borrower"`
	ContractNumber  string  `json:"contract_number"`
	SubmitDate      string  `json:"submit_date"`
	LoanStartDate   string  `json:"loan_start_date"`
	LoanAmount      string  `json:"loan_amount"`
	Rate            string  `json:"rate"`
	Tenor           int     `json:"tenor"`
	PremiumAmount   string  `json:"premium_amount"`
	LinkCertificate string  `json:"link_certificate"`
	Status          string  `json:"status"`
	FundingPartner  string  `json:"funding_partner"`
	URLSertifikat   *string `json:"url_sertifikat"`
	YearMonth       string  `json:"yearmonth"`
	Risk            string  `json:"risk"`
	PremiRefund     *string `json:"premi_refund"`
	RemarkRefund    *string `json:"remark_refund"`
	Usia            *string `json:"usia"`
}

type PaginatorPolicy struct {
	Items       []PolicyData        `json:"items"`
	PerPage     int                 `json:"perPage"`
	CurrentPage int                 `json:"currentPage"`
	Path        string              `json:"path"`
	Query       map[string]string  `json:"query"`
	Fragment    interface{}        `json:"fragment"`
	PageName    string              `json:"pageName"`
	OnEachSide  int                 `json:"onEachSide"`
	Options     map[string]interface{} `json:"options"`
	Total       int                 `json:"total"`
	LastPage    int                 `json:"lastPage"`
	Periods     []Period            `json:"periods"`
}
