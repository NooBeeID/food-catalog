package menu

import "time"

// ada tag `json`
// adalah sebagai paramter bahwa struct ini akan mengambil
// value sesuai dengan json yang diinputan
type createMenuRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"description"`
	Price    int    `json:"price"`
}

type listMenuResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"description"`
	Price    int    `json:"price"`
}

type singleMenuResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Desc      string    `json:"description"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
