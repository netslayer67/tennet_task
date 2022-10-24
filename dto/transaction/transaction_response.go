package transactiondto

type TransactionResponse struct {
	SrcWalletID  int32   `json:"src_wallet_id"`
	SrcAssetID   int32   `json:"src_asset_id"`
	DestWalletID int32   `json:"dest_wallet_id"`
	DestAssetID  int32   `json:"dest_asset_id"`
	Amount       float64 `json:"amount"`
	GasFee       float64 `json:"gas_fee"`
	Total        float64 `json:"total"`
}
