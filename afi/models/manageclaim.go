package models

// ini adalah stuktur manageclaim
type ClaimListData struct {
	ClaimId           string      `json:"claim_id"`
	LoanId            string      `json:"loan_id"`
	ContractNumber    string      `json:"contract_number"`
	NominalOutstanding string     `json:"nominal_outstanding"`
	Dpd               string      `json:"dpd"`
	FundingPartner    string      `json:"funding_partner"`
	Product           string      `json:"product"`
	Tenor             string      `json:"tenor"`
	PremiumAmount     string      `json:"premium_amount"`
	Status            string      `json:"status"`
	Remark            interface{} `json:"remark"`
	Yearmonth         string      `json:"yearmonth"`
	BatchPolicy       string      `json:"batch_policy"`
	CreatedAt         string      `json:"created_at"`
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
