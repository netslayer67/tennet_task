package repositories

import "task/models"

func (r *repository) CreateWallet(Wallet models.Wallet) (models.Wallet, error) {
	err := r.db.Create(&Wallet).Error

	return Wallet, err
}

func (r *repository) FindWallet() ([]models.Wallet, error) {
	var wallets []models.Wallet
	err := r.db.Model(&models.Wallet{}).Preload("Assets").Find(&wallets).Error

	return wallets, err
}
