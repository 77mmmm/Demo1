package handler

import (
	"awesomeProject/lib/ctx"
	"awesomeProject/logic/user"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	accountID := r.FormValue("username")
	password := r.FormValue("password")
	user := user.User{Ctx: ctx.GetGlobalCtx()}
	token, err := user.Login(accountID, password)
	if err != nil {
		http.Error(w, fmt.Sprintf("error checking user: %s", err), http.StatusBadRequest)
	}
	_, err = w.Write([]byte(token))
	if err != nil {
		log.Println(err)
	}
}
