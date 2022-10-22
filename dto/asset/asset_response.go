package assetdto

type AssetResponse struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Network string  `json:"network"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}
