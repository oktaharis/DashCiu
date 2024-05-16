package models

type UserData struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	Phone           string `json:"phone"`
	LastLogin       string `json:"last_login"`
	Status          string `json:"status"`
	RememberToken   string `json:"remember_token"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	RoleId          string `json:"role_id"`
	ProductId       string `json:"ProductId"`
	Otp             string `json:"otp"`
	ExpiredAt       string `json:"expired_at"`
	Uid             string `json:"uid"`
	ResetPassword   string `json:"reset_password"`
}

type PaginatorUser struct {
	Items       []UserData             `json:"items"`
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
