package entity

import "time"

type Orchard struct {
	ID       string
	Name     string
	Location string
	Owner    string
	Active   bool
	CreateAt time.Time
}
