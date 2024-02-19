package models

import (
	"time"
)

type Product struct {
	Id            int64     `json:"id"`
	Name  	      string    `json:"name"`
	Description   string    `json:"description"`
	Price      	  float32   `json:"price"`
	CreatedAt  	  time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt  	  time.Time `orm:"auto_now;type(datetime)" json:"-"`
}