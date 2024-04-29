package models

import "time"

// Monitoring represents an individual monitoring log entry.
type Monitoring struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Role       string    `json:"role"`
	Activity   string    `json:"activity"`
	DateTime   time.Time `json:"datetime"`
	DateTimeStr string    `json:"datetime_str"` // Tambahkan field untuk menyimpan datetime dalam format string
	Location   string    `json:"location"`
	IP         string    `json:"ip"`
}


// LengthAwarePaginator represents pagination information for a list of monitoring logs.
type LengthAwarePaginator struct {
	Items      []Monitoring `json:"items"`
	PerPage    int          `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	Path       string       `json:"path"`
	Query      map[string]string `json:"query"`
	Fragment   string       `json:"fragment"`
	PageName   string       `json:"page_name"`
	OnEachSide int          `json:"on_each_side"`
	Options    map[string]string `json:"options"`
	Total      int          `json:"total"`
	LastPage   int          `json:"last_page"`
}
