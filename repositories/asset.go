package repositories

import (
	assetdto "task/dto/asset"
	"task/models"
)

func (r *repository) CreateAsset(Asset models.Asset) (models.Asset, error) {
	err := r.db.Create(&Asset).Error

	return Asset, err
}

func (r *repository) FindAsset() ([]models.Asset, error) {
	var assets []models.Asset
	err := r.db.Find(&assets).Error

	return assets, err
}

func (r *repository) UpdateAsset(assetId int, param *assetdto.UpdateAsset) error {

	err := r.db.Model(&models.Asset{}).Where("id = ?", assetId).
		Updates(&models.Asset{
			Name:    param.Name,
			Network: param.Network,
		}).Error

	return err
}

func (r *repository) DeleteAsset(assetId int) error {

	err := r.db.Delete(&models.Asset{}, assetId).Error

	return err
}
