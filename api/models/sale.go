package models

type Sale struct {
	Id        string `json:"id	"`
	UserId    string `json:"user_id"`
	Total     int    `json:"total"`
	Count     int    `json:"count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type SaleRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
}
type SaleResponse struct {
	Count int     `json:"count"`
	Sales []*Sale `json:"sales"`
}

type SalePrimaryKey struct {
	Id string `json:"id"`
}
