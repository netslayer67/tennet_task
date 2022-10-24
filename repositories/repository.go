package repositories

import (
	assetdto "task/dto/asset"
	walletdto "task/dto/wallet"
	"task/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	//Aset Repository
	CreateAsset(asset models.Asset) error
	FindAsset() ([]models.Asset, error)
	UpdateAsset(assetId int, param *assetdto.UpdateAsset) error
	DeleteAsset(assetId int) error

	//Wallet repository
	CreateWallet(Wallet models.Wallet) (models.Wallet, error)
	FindWallet() ([]models.Wallet, error)
	UpdateWallet(walletId int, param *walletdto.UpdateWallet) error
	DeleteWallet(walletId int) error

	//AssetTransaction Repository
	CreateTransaction(AssetTransaction models.Transaction) error
	GetAssetByID(ID int) (models.Asset, error)
	UpdateAssetBalance(Asset models.Asset, assetBalance float64, isAdd bool) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
