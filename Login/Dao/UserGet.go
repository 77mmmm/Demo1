package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string
	Password string
	Email    string
}

type UserRepository interface {
	GetUser(username string) (string, error)
}

// 后续加入redis
/*type RedisRepository struct {
	client *redis.Client
}*/
type MysqlRepository struct {
	Db *gorm.DB
}

func (m *MysqlRepository) GetUser(username string) (string, error) {
	var user User
	var password string
	err := m.Db.Table("mydb").Where("username=?", username).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("error querying MySQL: %s", err)

	}
	password = user.Password
	return password, nil
}

// 使用多态加入redis，后续加入这个代码
/*func (r *RedisRepository) GetUser(username string) (string, error) {

}
*/
