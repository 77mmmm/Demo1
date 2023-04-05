package mysql

import (
	"awesomeProject/Login/Dao"
	"awesomeProject/Login/Entity"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var DB = Dao.MysqlRepository{}
var err error

func init() {
	DB.Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/dbsearch")
	if err != nil {
		fmt.Errorf("error:%s", err)
	} else {
		fmt.Println("连接成功")
	}

	err = DB.Db.Ping()
	if err != nil {
		fmt.Errorf("error:%s", err)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		accountID := r.FormValue("username")
		password := r.FormValue("password")

		correct, err := Entity.ValidateUser(accountID, password, &DB)

		if correct == true {
			fmt.Println("用户密码正确")
			token, err := Entity.GenerateToken(accountID)
			if err != nil {
				fmt.Errorf("错误信息:%s", err)
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
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
