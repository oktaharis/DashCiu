package models

// ini adalah stuktur manageclaim
type FilesRedinesData struct {
	Yearmonth	             string `json:"yearmonth"`
	Label             		 string `json:"label"`
	Policy             		 string `json:"policy"`
	Claim             		 string `json:"claim"`
	UpdatedAt	             string `json:"updated_at"`
	IsProcess	             string `json:"is_process"`
}

type PaginatorFilesRedines struct {
	Items       []FilesRedinesData     `json:"items"`
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
