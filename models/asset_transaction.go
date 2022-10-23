package models

type Transaction struct {
	ID           int     `json:"id" gorm:"primary_key:auto_increment"`
	SrcWalletID  int32   `json:"src_wallet_id" gorm:"type:bigint(20)"`
	SrcAssetID   int32   `json:"src_asset_id" gorm:"type:bigint(20)"`
	DestWalletID int32   `json:"dest_wallet_id" gorm:"type:bigint(20)"`
	DestAssetID  int32   `json:"dest_asset_id" gorm:"type:bigint(20)"`
	Amount       float64 `json:"amount" gorm:"type:decimal(16,8)"`
	GasFee       float64 `json:"gas_fee" gorm:"type:decimal(16,8)"`
	Total        float64 `json:"total" gorm:"type:decimal(16,8)"`
}
