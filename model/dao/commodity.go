package dao

import (
	"awesomeProject/lib/mysql"
	"awesomeProject/model"
	"fmt"
)

func NewCommodityRepositoryImpl() *CommodityRepositoryImpl {
	return &CommodityRepositoryImpl{
		db: mysql.NewCommodityRepository(),
	}
}

type CommodityRepositoryImpl struct {
	db *mysql.CommodityRepository
}

func (m *CommodityRepositoryImpl) GetCommodity(commodityname string) (model.Commodity, error) {
	commodityInfo, err := m.db.GetCommodity(commodityname)

	if err != nil {
		return model.Commodity{}, fmt.Errorf("error querying MySQL: %s", err)
	}
	return commodityInfo, nil
}

func (m *CommodityRepositoryImpl) Purchase(commodityname string) (bool, error) {
	Bool, err := m.db.Purchase(commodityname)
	if err != nil {
		return false, fmt.Errorf("error querying MySQL: %s", err)
	}
	return Bool, nil
}
