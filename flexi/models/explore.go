package models

// ini adalah stuktur manageclaim
type ExploreData struct {
	NoPk	                 interface{} `json:"no_pk"`
	Nik             		 interface{} `json:"nik"`
	Name             		 string `json:"name"`
	Policy             		 string `json:"policy"`
	Claim             		 string `json:"claim"`
	PolicyYearmonth          string `json:"policy_yearmonth"`
	ClaimYearmonth           string `json:"claim_yearmonth"`
	LtcByNik                 string `json:"ltc_by_nik"`
}

type PaginatorExplore struct {
	Items       []ExploreData         `json:"items"`
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
