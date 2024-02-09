package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Name  string `json:"name" xml:"name" form:"name"`
	Saldo int64  `json:"saldo" xml:"saldo" form:"saldo"`
}
