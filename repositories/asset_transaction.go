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
