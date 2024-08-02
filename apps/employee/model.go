package employee

import "time"

type Employee struct {
	Id        int
	Name      string
	Address   string
	NIP       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
