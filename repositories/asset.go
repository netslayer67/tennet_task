package repositories

import (
	"task/models"

	"gorm.io/gorm"
)

type AssetRepository interface {
	CreateAsset(asset models.Asset) (models.Asset, error)
	FindAsset() ([]models.Asset, error)
}

func RepositoryAsset(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAsset(Asset models.Asset) (models.Asset, error) {
	err := r.db.Create(&Asset).Error

	return Asset, err
}

func (r *repository) FindAsset() ([]models.Asset, error) {
	var assets []models.Asset
	err := r.db.Find(&assets).Error

	return assets, err
}
