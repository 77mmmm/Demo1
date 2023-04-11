package mysql

import (
	"fmt"

	"awesomeProject/model"

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

func (m *UserRepository) SelectUser(username string, password string, email string) error {
	tx := m.db.Begin()
	//err := tx.Table("user").Set("gorm:query_option", "FOR UPDATE ").Where("username=?", username).Count(&count1).Error
	//err1 := tx.Table("user").Set("gorm:query_option", "FOR UPDATE").Where("email=?", email).Count(&count2).Error
	list := []model.User{}
	err := tx.Model(&model.User{}).Set("gorm:query_option", "LOCK IN SHARE MODE").Where("username = ?", username).Limit(1).Find(&list).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error querying MySQL: %s", err)

	}
	if len(list) != 0 {
		tx.Rollback()
		return fmt.Errorf("error querying MySQL: %s", err)
	}

	err = tx.Model(&model.User{}).Create(model.User{
		Username: username,
		Password: password,
		Email:    email,
	}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error querying MySQL: %s", err)
	}
	tx.Commit()
	return nil
}
