package models

import (
	"fmt"
	"go-minikube/api/utils"
	"time"
)

type Log struct {
	ID             uint64     `gorm:"primary_key;auto_increment" json:"id"`
	WalletOriginID uint32     `gorm:"not null" json:"wallet_origin_id"`
	WalletTargetID uint32     `gorm:"not null" json:"wallet_target_id"`
	Amount         float64    `gorm:"not null" json:"amount"`
	Description    string     `gorm:"size:255;not null" json:"description"`
	CreatedAt      *time.Time `gorm:"default:(datetime('now','localtime'))" json:"created_at"`
}

func GetLogs() []Log {
	db := Connect()
	defer db.Close()
	var logs []Log
	db.Order("id asc").Find(&logs)
	return logs
}

func PaginateLogs(page, offset, limit int) (interface{}, error) {
	db := Connect()
	defer db.Close()
	var total int64
	err := db.Model(&Log{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	var logs []Log
	err = db.Offset(offset).Limit(limit).Order("id desc").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return utils.Pagination(logs, len(logs), limit, int(total), page), nil
}

func logHandler(origin, target Wallet, cash float64) Log {
	return Log{WalletOriginID: origin.ID, WalletTargetID: target.ID, Amount: cash,
		Description: fmt.Sprintf("%s transferred $%.2f to %s",
			origin.Owner.FirstName, cash, target.Owner.FirstName)}
}
