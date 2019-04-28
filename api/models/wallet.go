package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go-minikube/api/utils"
	"time"
)

var (
	ErrWalletNotExists  = errors.New("wallet not exists")
	ErrInvalidCash      = errors.New("invalid cash")
	ErrInvalidAction    = errors.New("invalid action")
	ErrInsufficientCash = errors.New("insufficient money")
	ErrInvalidTransfer  = errors.New("invalid transfer")
)

const (
	SUM = "sum"
	SUB = "sub"
)

type Wallet struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	PublicKey string     `gorm:"size:32;not null;unique" json:"public_key"`
	Cash      float64    `gorm:"default: 0.0" json:"cash"`
	Owner     Owner      `gorm:"foreignkey:OwnerID" json:"owner"`
	OwnerID   uint32     `gorm:"not null" json:"owner_id"`
	UpdatedAt *time.Time `gorm:"default:(datetime('now','localtime'))" json:"updated_at"`
}

func (wallet *Wallet) PublicKeyGenerator() error {
	if wallet.Owner.ID == 0 {
		return ErrWalletNotExists
	}
	wallet.PublicKey = fmt.Sprintf("%x", md5.Sum([]byte(wallet.Owner.Email+wallet.Owner.Password)))
	return nil
}

func FilterWallets(wallets []Wallet) []Wallet {
	var filter []Wallet
	for _, wallet := range wallets {
		if wallet.Owner.Status != 0 {
			filter = append(filter, wallet)
		}
	}
	return filter
}

func GetWallets() []Wallet {
	db := Connect()
	defer db.Close()
	var wallets []Wallet
	db.Table("wallets").Select("*").Joins("inner join owners on owners.id = wallets.owner_id").Where("owners.status != ?", 0).Scan(&wallets)
	for i := range wallets {
		db.Model(&wallets[i]).Related(&wallets[i].Owner)
	}
	// remove owners with status 0
	return wallets
}

func PaginateWallets(page, offset, limit int) (interface{}, error) {
	db := Connect()
	defer db.Close()
	var total int64
	err := db.Model(&Owner{}).Where("status != ?", 0).Count(&total).Error
	if err != nil {
		return nil, err
	}
	var wallets []Wallet
	err = db.Table("wallets").Select("*").Joins("inner join owners on owners.id = wallets.owner_id").
		Offset(offset).Limit(limit).Where("owners.status != ?", 0).Scan(&wallets).Error
	if err != nil {
		return nil, err
	}
	for i := range wallets {
		db.Model(&wallets[i]).Related(&wallets[i].Owner)
	}
	return utils.Pagination(wallets, len(wallets), limit, int(total), page), nil
}

func GetWalletById(id uint32) (Wallet, error) {
	db := Connect()
	defer db.Close()
	var wallet Wallet
	db.Where("id = ?", id).Find(&wallet)
	db.Model(&wallet).Related(&wallet.Owner)
	if wallet.ID == 0 || wallet.Owner.Status == 0 {
		return Wallet{}, ErrWalletNotExists
	}
	return wallet, nil
}

func UpdateWallet(wallet Wallet, action string) (int64, error) {
	db := Connect()
	defer db.Close()
	if wallet.Cash < 1 {
		return 0, ErrInvalidCash
	}
	old, err := GetWalletByPublicKey(wallet.PublicKey)
	if err != nil {
		return 0, err
	}
	cash, err := GetCashByAction(old.Cash, wallet.Cash, action)
	if err != nil {
		return 0, err
	}
	rs := db.Model(&wallet).Where("public_key = ?", wallet.PublicKey).UpdateColumns(
		map[string]interface{}{
			"cash":       cash,
			"updated_at": time.Now(),
		})
	return rs.RowsAffected, rs.Error
}

func GetCashByAction(oldValue, newValue float64, action string) (float64, error) {
	if action == SUM {
		return (oldValue + newValue), nil
	} else if action == SUB && oldValue > newValue {
		return (oldValue - newValue), nil
	}
	return oldValue, ErrInvalidAction
}

func GetWalletByPublicKey(publicKey string) (Wallet, error) {
	db := Connect()
	defer db.Close()
	var wallet Wallet
	db.Where("public_key = ?", publicKey).Find(&wallet)
	db.Model(&wallet).Related(&wallet.Owner)
	if wallet.ID == 0 || wallet.Owner.Status == 0 {
		return Wallet{}, ErrWalletNotExists
	}
	return wallet, nil
}

func Transfer(keys []string, cash float64) (Log, error) {
	db := Connect()
	defer db.Close()
	origin, target, err := VerifyWallets(keys, cash)
	if err != nil {
		return Log{}, err
	}
	tx := db.Begin()
	err = tx.Debug().Model(&origin).Where("id = ?", origin.ID).UpdateColumns(map[string]interface{}{
		"cash": (origin.Cash - cash), "updated_at": time.Now()}).Error
	if err != nil {
		tx.Rollback()
		return Log{}, err
	}
	rs := tx.Debug().Model(&target).Where("id = ?", target.ID).UpdateColumns(map[string]interface{}{
		"cash": (target.Cash + cash), "updated_at": time.Now()})
	if rs.Error != nil {
		tx.Rollback()
		return Log{}, err
	}
	log := logHandler(origin, target, cash)
	err = tx.Debug().Create(&log).Error
	if err != nil {
		tx.Rollback()
		return Log{}, err
	}
	return log, tx.Commit().Error
}

func VerifyWallets(keys []string, cash float64) (Wallet, Wallet, error) {
	origin, err := GetWalletByPublicKey(utils.Trim(keys[0]))
	if err != nil {
		return Wallet{}, Wallet{}, err
	}
	target, err := GetWalletByPublicKey(utils.Trim(keys[1]))
	if err != nil {
		return Wallet{}, Wallet{}, err
	}
	if origin.ID == target.ID {
		return Wallet{}, Wallet{}, ErrInvalidTransfer
	}
	if origin.Cash < cash {
		return Wallet{}, Wallet{}, ErrInsufficientCash
	}
	return origin, target, nil
}
