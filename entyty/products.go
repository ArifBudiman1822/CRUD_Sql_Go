package entity

import "time"

type Products struct {
	Id           int32
	Product_name string
	Price        int32
	Quantity     int32
	Created_at   *time.Time
}
