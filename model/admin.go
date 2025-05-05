package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       uint
	PasswordDigest string
	Avatar         string
	Address        string
}
