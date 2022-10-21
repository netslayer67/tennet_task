package models

type Wallet struct {
	ID     int     `json:"id" gorm:"primary_key:auto_increment"`
	Name   string  `json:"name" gorm:"type:varchar(200)"`
	Assets []Asset `json:"assets"`
}

type WalletAsset struct {
	ID int `json:"id"`
}

func (WalletAsset) TableName() string {
	return "assets"
}
