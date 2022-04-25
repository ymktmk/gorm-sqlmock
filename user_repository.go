package main

import "gorm.io/gorm"

func Store(db *gorm.DB, u *User) (user *User, err error) {
	if err = db.Create(u).Error; err != nil {
		return
	}
	user = u
	return
}

func FindById(db *gorm.DB, id int) (user *User, err error) {
	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		return
	}
	return
}