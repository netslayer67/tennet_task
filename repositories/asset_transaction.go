package repositories

import "task/models"

func (r *repository) CreateTransaction(AssetTransaction models.Transaction) error {
	err := r.db.Create(&AssetTransaction).Error

	return err
}

func (r *repository) GetAssetByID(ID int) (models.Asset, error) {
	var Asset models.Asset
	err := r.db.Find(&Asset, ID).Error

	return Asset, err
}

func (r *repository) UpdateAssetBalance(Asset models.Asset, assetBalance float64, isAdd bool) error {

	r.db.Find(&Asset)

	if isAdd {
		Asset.Balance += assetBalance
	} else {
		Asset.Balance -= assetBalance
	}

	err := r.db.Save(&Asset).Error

	return err

}
