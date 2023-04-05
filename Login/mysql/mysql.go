package mysql

import (
	"awesomeProject/Login/Dao"
	"awesomeProject/Login/Entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

var DB = Dao.MysqlRepository{}

func init() {
	var err error
	DB.Db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/dbsearch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		err = fmt.Errorf("error:%s", err)
		fmt.Println(err.Error())
	} else {
		fmt.Println("连接成功")
	}
	DB.Db.AutoMigrate(&Dao.User{})
	err = DB.Db.DB().Ping()
	if err != nil {
		err = fmt.Errorf("error:%s", err)
		fmt.Println(err.Error())
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		accountID := r.FormValue("username")
		password := r.FormValue("password")

		correct, err := Entity.ValidateUser(accountID, password, &DB)
		if err != nil {

		}

		if correct == true {
			fmt.Println("用户密码正确")
			token, err1 := Entity.GenerateToken(accountID)
			if err1 != nil {

				err1 = fmt.Errorf("错误信息:%s", err1)
				fmt.Println(err1.Error())
			}
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

		} else {
			http.Error(w, fmt.Sprintf("error checking user: %s", err), http.StatusBadRequest)
		}

	}

}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}
		claims, err := Entity.ParseToken(authorization)
		if err != nil {
			http.Error(w, "expire the token", http.StatusUnauthorized)
		}
		fmt.Println("token正确")
		_, _ = w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
