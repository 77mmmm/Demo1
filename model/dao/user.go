package dao

import (
	"fmt"

	"awesomeProject/lib/mysql"
)

// 后续加入redis
/*type RedisRepository struct {
	client *redis.Client
}*/

func NewRepository() *Repository {
	return &Repository{db: mysql.NewUserRepository()}
}

type Repository struct {
	db *mysql.UserRepository
}

func (m *Repository) GetUser(username string) (string, error) {
	password, err := m.db.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("error querying MySQL: %s", err)

	}
	return password, nil
}
func (m *Repository) SelectUser(username string, password string, email string) (bool, error) {
	err := m.db.SelectUser(username, password, email)
	if err != nil {
		return false, fmt.Errorf("error querying MySQL: %s", err)
	}

	return true, nil
}
