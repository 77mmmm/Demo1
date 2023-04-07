package cache

import (
	"log"
	"time"

	"awesomeProject/model/dao"

	"github.com/go-redis/redis"
)

type UserCache struct {
	client *redis.Client
	user   *dao.Repository
}

func NewUserCache() *UserCache {
	return &UserCache{
		user: dao.NewRepository(),
	}
}

func (u *UserCache) GetUser(username string) (string, error) {
	if u.client == nil {
		return u.user.GetUser(username)
	}
	//当缓存的客户端被初始化之后，则直接从缓存中拿去结果
	strIface, err := u.client.Do("").Result()
	if err != nil {
		return u.user.GetUser(username)
	}
	if strIface == nil {
		var str string
		str, err = u.user.GetUser(username)
		if err != nil {
			return "", err
		}
		err1 := u.client.Set(username, str, 3600*time.Second).Err()
		if err1 != nil {
			log.Println(err)
		}
		return str, err
	}
	return strIface.(string), nil
}
func (u *UserCache) SelectUser(username string, password string, email string) (bool, error) {

	return u.user.SelectUser(username, password, email)
}
