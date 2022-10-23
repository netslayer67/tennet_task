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
	CreateAsset(asset *assetdto.CreateAsset) ( error)
	FindAsset() ([]models.Asset, error)
	UpdateAsset(assetId int, param *assetdto.UpdateAsset) error
	DeleteAsset(assetId int) error

	//Wallet repository
	CreateWallet(Wallet models.Wallet) (models.Wallet, error)
	FindWallet() ([]models.Wallet, error)

	//AssetTransaction Repository
	CreateTransaction(param interface{}) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
