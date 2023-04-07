package handler

import (
	"awesomeProject/lib/ctx"
	"awesomeProject/logic/commodity"
	"fmt"
	"net/http"
)

func Purchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	commodityname := r.FormValue("commodityname")

	commodity := commodity.Commodity{Ctx: ctx.GetGlobalCtx()}
	Bool, err := commodity.Purchase(commodityname)
	if err != nil {
		http.Error(w, fmt.Sprintf("error checking commodity: %s", err), http.StatusBadRequest)
	}
	if Bool == false {
		http.Error(w, "商品为空", http.StatusBadRequest)
	} else {
		fmt.Println("购买成功")
	}
}
