package model

type UserRepository interface {
	GetUser(username string) (string, error)
	SelectUser(username string, password string, email string) (bool, error)
}

type CommodityRepository interface {
	GetCommodity(commodityname string) (Commodity, error)
	Purchase(commodityname string) (bool, error)
}
