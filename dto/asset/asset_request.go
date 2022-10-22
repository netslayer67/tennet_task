package assetdto

type CreateAsset struct {
	WalletID int     `json:"wallet_id"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Network  string  `json:"network"`
	Address  string  `json:"address"`
	Balance  float64 `json:"balance"`
}

type UpdateAsset struct {
	Name    string `json:"name"`
	Network string `json:"network"`
}
