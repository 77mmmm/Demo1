package commodity

import "awesomeProject/lib/ctx"

type Commodity struct {
	Ctx *ctx.GlobalCtx
}

func (u *Commodity) Purchase(commodityname string) (bool, error) {
	Bool, err := ctx.GetGlobalCtx().Commodity.Purchase(commodityname)
	if err != nil {
		return false, err
	}
	if Bool == false {
		return Bool, nil
	}
	return Bool, nil
}
