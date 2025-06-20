package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	Surname   string `gorm:"type:varchar(100);index;not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	IsActive  bool   `gorm:"default:true"`
	IsBlocked bool   `gorm:"default:false"`
}
