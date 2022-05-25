package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"

	"gorm.io/gorm"
)

// GetComerByAddress  get comer entity by comer's address
func GetComerByAddress(db *gorm.DB, address string, comer *Comer) error {
	return db.Where("address = ? AND is_deleted = false", address).Find(comer).Error
}

// GetComerByID  get comer entity by comer's ID
func GetComerByID(db *gorm.DB, comerID uint64, comer *Comer) (err error) {
	return db.Where("id = ? AND is_deleted = false", comerID).Find(comer).Error
}

// CreateComer create a comer
func CreateComer(db *gorm.DB, comer *Comer) (err error) {
	return db.Create(comer).Error
}

//UpdateComerAddress update the comer address
func UpdateComerAddress(db *gorm.DB, comerID uint64, address string) (err error) {
	return db.Model(&Comer{Base: model.Base{ID: comerID}}).Update("address", address).Error
}

func GetComerAccount(db *gorm.DB, accountType ComerAccountType, oin string, comerAccount *ComerAccount) error {
	return db.Where("type = ? AND oin = ? AND is_deleted = false", accountType, oin).Find(comerAccount).Error
}

func ListAccount(db *gorm.DB, comerID uint64, accountList *[]ComerAccount) (err error) {
	return db.Where("comer_id = ? AND is_deleted = false", comerID).Find(accountList).Error
}

func CreateAccount(db *gorm.DB, comerAccount *ComerAccount) (err error) {
	return db.Create(comerAccount).Error
}

func DeleteAccount(db *gorm.DB, comerID, accountID uint64) error {
	return db.Where("comer_id = ? AND id = ? AND is_deleted = false", comerID, accountID).Delete(&ComerAccount{}).Error
}

//GetComerProfile update the comer address
func GetComerProfile(db *gorm.DB, comerID uint64, profile *ComerProfile) (err error) {
	return db.Where("comer_id = ? AND is_deleted = false", comerID).Preload("Skills", "category = ?", tag.ComerSkill).Find(profile).Error
}

//CreateComerProfile update the comer address
func CreateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Create(&comerProfile).Error
}

//UpdateComerProfile update the comer address
func UpdateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Where("comer_id = ? AND is_deleted = false", comerProfile.ComerID).Select("avatar", "name", "location", "time_zone", "website", "email", "twitter", "discord", "telegram", "medium", "bio").Updates(comerProfile).Error
}
