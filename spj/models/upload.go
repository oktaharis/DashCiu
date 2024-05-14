package models

// ini adalah stuktur manageclaim
type UploadData struct {
	Id                string `json:"id"`
	ProductCode       string `json:"product_code"`
	OriginaFileName   string `json:"origina_file_name"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
	RowStatus         string `json:"row_status"`
	UploadDateTime    string `json:"upload_date_time"`
	Template          string `json:"template"`
	Success           string `json:"success"`
	Failed            string `json:"failed"`
	Duplicate         string `json:"duplicate"`
	Row               string `json:"row"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Path              string `json:"path"`
	Type              string `json:"type"`
	Yearmonth         string `json:"yearmonth"`
	StartInsert       string `json:"start_insert"`
	EndInsert         string `json:"end_insert"`
	TotalRowsInserted string `json:"total_rows_inserted"`
	ProcessingTimeS   *string `json:"processing_time_s"`
}

type PaginatorUpload struct {
	Items       []UploadData           `json:"items"`
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
