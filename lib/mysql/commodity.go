package mysql

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

func NewCommodityRepository() *CommodityRepository {
	db, err := InitDb()
	if err != nil {
		panic(err)
	}
	return &CommodityRepository{
		db: db,
	}
}

type CommodityRepository struct {
	db *gorm.DB
}

func (m *CommodityRepository) GetCommodity(commodityname string) (model.Commodity, error) {
	var commodity model.Commodity
	err := m.db.Table("commodity").Where("commodityname=?", commodityname).First(&commodity).Error
	if err != nil {
		return model.Commodity{}, err
	}
	return commodity, nil

}
func (m *CommodityRepository) Purchase(commodityname string) (bool, error) {
	var commodity model.Commodity
	err := m.db.Table("commodity").Where("commodityname=?", commodityname).First(&commodity).Error
	if err != nil {
		return false, err
	}
	if commodity.Stock == 0 {
		return false, nil
	}
	err = m.db.Table("commodity").Where("commodityname=?", commodityname).Update("stock", commodity.Stock-1).Error
	if err != nil {
		return false, err

	}
	return true, nil
}
