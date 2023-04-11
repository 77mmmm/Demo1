package cache

import (
	"awesomeProject/model"
	"awesomeProject/model/dao"
	"github.com/go-redis/redis"
	"strconv"
)

type CommodityCache struct {
	client    *redis.Client
	commodity *dao.CommodityRepositoryImpl
}

func NewCommodityCache() *CommodityCache {
	return &CommodityCache{
		commodity: dao.NewCommodityRepositoryImpl(),
	}
}

func (u *CommodityCache) GetCommodity(commodityname string) (model.Commodity, error) {
	if u.client == nil {
		return u.commodity.GetCommodity(commodityname)
	}
	values, err := u.client.HMGet("comodity:"+commodityname, "commodityprice", "stock").Result()
	if err != nil {
		return model.Commodity{}, nil
	}
	strIace1, err := strconv.ParseFloat(values[0].(string), 32)
	if err != nil {
		return model.Commodity{}, nil
	}
	strIace2, err := strconv.Atoi(values[1].(string))
	if err != nil {
		return model.Commodity{}, nil
	}
	return model.Commodity{
		CommodityName:  commodityname,
		Commodityprice: strIace1,
		Stock:          strIace2,
	}, nil

}
func (u *CommodityCache) Purchase(commodityname string) (bool, error) {
	return u.commodity.Purchase(commodityname)
}
