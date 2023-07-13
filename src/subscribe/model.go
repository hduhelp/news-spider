package subscribe

import "time"

type Subscribe struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	StaffID   string     `json:"-"`
	Tag       string
	Name      string
}
