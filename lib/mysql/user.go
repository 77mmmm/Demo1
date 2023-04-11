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
<<<<<<< HEAD
	tx := m.db.Begin()
	err := tx.Table("user").Set("gorm:query_option", "FOR UPDATE ").Where("username=?", username).Count(&count1).Error
	err1 := tx.Table("user").Set("gorm:query_option", "FOR UPDATE").Where("email=?", email).Count(&count2).Error
=======
	err := m.db.Table("user").Where("username=?", username).Count(&count1).Error
	err1 := m.db.Table("user").Where("email=?", email).Count(&count2).Error
>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
	if err != nil || err1 != nil {
		return false, fmt.Errorf("error querying MySQL: %s", err)

	}
<<<<<<< HEAD

	tx.Table("user").Create(model.User{
=======
	if count1 > 0 {
		return false, fmt.Errorf("Username  already exists")
	}
	if count2 > 0 {
		return false, fmt.Errorf("email already exists")
	}
	m.db.Table("user").Create(model.User{
>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
		Username: username,
		Password: password,
		Email:    email,
	})
<<<<<<< HEAD
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

=======

	return true, nil
>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
}
