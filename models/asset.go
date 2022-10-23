package models

type Asset struct {
	ID       int     `json:"id" gorm:"primary_key:auto_increment"`
	WalletID int     `json:"wallet_id"`
	Wallet   Wallet  `json:"-"`
	Name     string  `json:"name" gorm:"type:varchar(255)"`
	Symbol   string  `json:"symbol" gorm:"type:varchar(100)"`
	Network  string  `json:"network" gorm:"type:varchar(100)"`
	Address  string  `json:"address" gorm:"type:varchar(42)"`
	Balance  float64 `json:"balance" gorm:"type:decimal(16,8)"`
	Token    string  `json:"token"`
}

type WalletAsset struct {
	ID int `json:"id"`
}

func (WalletAsset) TableName() string {
	return "assets"
}
