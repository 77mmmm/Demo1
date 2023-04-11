package mysql

import (
	"awesomeProject/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

func NewUserRepository() *UserRepository {
	db, err := InitDb()
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (m *UserRepository) GetUser(username string) (string, error) {
	var user model.User
	var password string
	err := m.db.Table("user").Where("username=?", username).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("error querying MySQL: %s", err)

	}
	password = user.Password
	return password, nil
}

func (m *UserRepository) SelectUser(username string, password string, email string) (bool, error) {
	var count1 int
	var count2 int
	tx := m.db.Begin()
	err := tx.Table("user").Set("gorm:query_option", "FOR UPDATE ").Where("username=?", username).Count(&count1).Error
	err1 := tx.Table("user").Set("gorm:query_option", "FOR UPDATE").Where("email=?", email).Count(&count2).Error
	if err != nil || err1 != nil {
		return false, fmt.Errorf("error querying MySQL: %s", err)

	}

	tx.Table("user").Create(model.User{
		Username: username,
		Password: password,
		Email:    email,
	})
	if count1 > 0 {
		tx.Rollback()
		return false, fmt.Errorf("Username  already exists")
	} else if count2 > 0 {
		tx.Rollback()
		return false, fmt.Errorf("email already exists")

	} else {
		tx.Commit()
		return true, nil
	}

}
