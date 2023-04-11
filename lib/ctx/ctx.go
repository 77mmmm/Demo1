package ctx

import (
	"awesomeProject/lib/mysql"
	"awesomeProject/model"
	"awesomeProject/model/cache"
	"github.com/jinzhu/gorm"
)

type GlobalCtx struct {
	User      model.UserRepository
	DB        *gorm.DB
	Commodity model.CommodityRepository
}

var ctx *GlobalCtx

func GetGlobalCtx() *GlobalCtx {

	if ctx != nil {
		return ctx
	}
	db, err := mysql.InitDb()
	if err != nil {
		panic(err)
	}
	ctx = &GlobalCtx{
		//User: dao.NewRepository(),
		User:      cache.NewUserCache(),
		DB:        db,
		Commodity: cache.NewCommodityCache(),
	}
	return ctx
}
