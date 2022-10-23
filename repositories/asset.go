package repositories

import (
	assetdto "task/dto/asset"
	"task/models"
)

func (r *repository) CreateAsset(asset models.Asset) error {
	err := r.db.Create(&asset).Error

	return err
}

func (r *repository) FindAsset() (assets []models.Asset, err error) {

	err = r.db.Find(&assets).Error

	return
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
