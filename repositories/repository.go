package repositories

import (
	assetdto "task/dto/asset"
	"task/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	//Aset Repository
	CreateAsset(asset models.Asset) (models.Asset, error)
	FindAsset() ([]models.Asset, error)
	UpdateAsset(assetId int, param *assetdto.UpdateAsset) error
	DeleteAsset(assetId int) error

	//AssetTransaction Repository
	CreateTransaction(param interface{}) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
