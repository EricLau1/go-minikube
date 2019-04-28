package models

import (
	"errors"
	"go-minikube/api/security"
	"go-minikube/api/utils"
	"time"
)

var (
	ErrOwnerNotFound = errors.New("owner not found")
)

type Owner struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string     `gorm:"size:35;not null" json:"first_name"`
	LastName  string     `gorm:"size:35;not null" json:"last_name"`
	Email     string     `gorm:"size:40;not null;unique" json:"email"`
	Password  string     `gorm:"size:60;not null" json:"password"`
	Gender    string     `gorm:"size:1;not null" json:"gender"`
	Status    uint8      `gorm:"default: 0" json:"status"`
	CreatedAt *time.Time `gorm:"default:(datetime('now','localtime'))" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:(datetime('now','localtime'))" json:"updated_at"`
}

func (owner *Owner) HashedPassword() error {
	hash, err := security.Hash(owner.Password)
	if err != nil {
		return err
	}
	owner.Password = string(hash)
	return nil
}

func NewOwner(owner Owner) (Wallet, error) {
	db := Connect()
	defer db.Close()
	tx := db.Begin()
	err := owner.HashedPassword()
	if err != nil {
		return Wallet{}, err
	}
	err = tx.Debug().Create(&owner).Error
	if err != nil {
		tx.Rollback()
		return Wallet{}, err
	}
	wallet := Wallet{Owner: owner}
	err = wallet.PublicKeyGenerator()
	if err != nil {
		return Wallet{}, err
	}
	err = tx.Debug().Create(&wallet).Error
	if err != nil {
		tx.Rollback()
		return Wallet{}, err
	}
	return wallet, tx.Commit().Error
}

func GetOwners() []Owner {
	db := Connect()
	defer db.Close()
	var owners []Owner
	db.Order("id asc").Where("status != ?", 0).Find(&owners)
	return owners
}

func PaginateOwners(page, offset, limit int) (interface{}, error) {
	db := Connect()
	defer db.Close()
	var total int64
	err := db.Model(&Owner{}).Where("status != ?", 0).Count(&total).Error
	if err != nil {
		return nil, err
	}
	var owners []Owner
	err = db.Offset(offset).Limit(limit).Where("status != ?", 0).Find(&owners).Error
	if err != nil {
		return nil, err
	}
	return utils.Pagination(owners, len(owners), limit, int(total), page), nil
}

func GetOwnerById(id uint32) (Owner, error) {
	db := Connect()
	defer db.Close()
	var owner Owner
	db.Where("id = ?", id).Find(&owner)
	if owner.ID == 0 || owner.Status == 0 {
		return Owner{}, ErrOwnerNotFound
	}
	return owner, nil
}

func UpdateOwner(owner Owner) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&owner).Where("id = ?", owner.ID).UpdateColumns(
		map[string]interface{}{
			"first_name": owner.FirstName,
			"last_name":  owner.LastName,
			"email":      owner.Email,
			"gender":     owner.Gender,
			"status":     1,
			"updated_at": time.Now(),
		},
	)
	return rs.RowsAffected, rs.Error
}

func DisableOwner(id uint32) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&Owner{}).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"status": 0,
		},
	)
	return rs.RowsAffected, rs.Error
}
