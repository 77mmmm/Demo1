package Dao

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetUser(username string) (string, error)
}

// 后续加入redis
/*type RedisRepository struct {
	client *redis.Client
}*/
type MysqlRepository struct {
	Db *sql.DB
}

func (m *MysqlRepository) GetUser(username string) (string, error) {

	var password string
	err := m.Db.QueryRow("SELECT password FROM mydb where username=?", username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", fmt.Errorf("error querying MySQL: %s", err)
		}
	}
	return password, nil
}

// 使用多态加入redis，后续加入这个代码
/*func (r *RedisRepository) GetUser(username string) (string, error) {

}
*/
