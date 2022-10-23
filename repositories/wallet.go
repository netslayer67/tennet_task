package repositories

import (
	walletdto "task/dto/wallet"
	"task/models"
)

func (r *repository) CreateWallet(Wallet models.Wallet) (models.Wallet, error) {
	err := r.db.Create(&Wallet).Error

	return Wallet, err
}

func (r *repository) FindWallet() ([]models.Wallet, error) {
	var wallets []models.Wallet
	err := r.db.Model(&models.Wallet{}).Preload("Assets").Find(&wallets).Error

	return wallets, err
}

func (r *repository) UpdateWallet(walletId int, param *walletdto.CreateWallet) error {

	err := r.db.Model(&models.Wallet{}).Where("id = ?", walletId).
		Updates(&models.Wallet{
			Name: param.Name,
		}).Error

	return err
}
