package commodity

import "awesomeProject/model"

type Productctx struct {
	Commodity model.CommodityRepository
}

func (p *Productctx) GetCommodity(commodityname string) (model.Commodity, error) {
	commodity_info, err := p.Commodity.GetCommodity(commodityname)
	if err != nil {
		return model.Commodity{}, err
	}
	return commodity_info, nil
}
