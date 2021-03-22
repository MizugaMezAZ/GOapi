package models

import "strings"

type Account struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

func (a *Account) CreateAC(ac *Account) (err error) {
	result := db.Create(ac)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (a *Account) DeleteAC(ac *Account) (err error) {
	result := db.Where("username = ?", ac.Username).Delete(ac)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (a *Account) UpdateAC(ac *Account) (err error) {
	result := db.Model(ac).Where("username = ?", ac.Username).Update("password", ac.Password)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (a *Account) FindAC(ac *Account) (err error, tr bool) {
	var checkac Account
	db.Where("username = ?", ac.Username).First(&checkac)
	result := db.Where("username = ?", ac.Username).First(&checkac)

	if result.Error != nil {
		err = result.Error
		return
	}

	if strings.Compare(ac.Password, checkac.Password) == 0 {
		return err, true
	} else {
		return err, false
	}
}
