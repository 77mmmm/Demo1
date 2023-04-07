package handler

import (
	"fmt"
	"log"
	"net/http"

	"awesomeProject/logic/commodity"
	"awesomeProject/model/cache"
)

func Product_Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	commodity_name := r.FormValue("commodityname")

	commodity := commodity.Productctx{
		Commodity: cache.NewCommodityCache(),
	}
	commodity_Info, err := commodity.GetCommodity(commodity_name)
	if err != nil {
		http.Error(w, fmt.Sprintf("error checking commodity: %s", err), http.StatusBadRequest)
	}
	commodity_price := commodity_Info.Commodityprice
	stock := commodity_Info.Stock
	data := fmt.Sprintf("%s %f %d", commodity_name, commodity_price, stock)
	_, err = w.Write([]byte(data))
	if err != nil {
		log.Println(err)
	}

}
