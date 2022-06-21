package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
	"gorm.io/gorm"
	"strings"
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

func GetComerAccountById(db *gorm.DB, accountId uint64, comerAccount *ComerAccount) error {
	return db.Where("id=? AND is_deleted = false", accountId).Find(comerAccount).Error
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

// CreateComerFollowRel create comer relation for comer and target comer
func CreateComerFollowRel(db *gorm.DB, comerID, targetComerID uint64) error {
	return db.Create(&FollowRelation{ComerID: comerID, TargetComerID: targetComerID}).Error
}

// DeleteComerFollowRel delete comer relation for comer and target comer
func DeleteComerFollowRel(db *gorm.DB, input *FollowRelation) error {
	return db.Where("comer_id = ? AND target_comer_id = ?", input.ComerID, input.TargetComerID).Delete(input).Error
}

// ComerFollowIsExist check startup and comer is existed
func ComerFollowIsExist(db *gorm.DB, comerID, targetComerID uint64) (isExist bool, err error) {
	isExist = false
	var count int64
	err = db.Table("comer_follow_rel").Where("comer_id = ? AND target_comer_id = ?", comerID, targetComerID).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func ListFollowComer(db *gorm.DB, comerID uint64, output *[]FollowComer) (total int64, err error) {
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	err = db.Order("created_at ASC").Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills").Find(output).Count(&total).Error
	return
}

func ListFollowedComer(db *gorm.DB, comerID uint64, output *[]FollowedComer) (total int64, err error) {
	if comerID != 0 {
		db = db.Where("target_comer_id = ?", comerID)
	}
	err = db.Order("created_at ASC").Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills").Find(output).Count(&total).Error
	return
}

//BindComerAccountToComerId bind comerAccount to comer
func BindComerAccountToComerId(db *gorm.DB, comerAccountId, comerID uint64) (err error) {
	var crtAccount ComerAccount
	db.First(&crtAccount, comerAccountId)
	if crtAccount.ComerID == comerID || crtAccount.ComerID == 0 {
		return db.Model(&ComerAccount{Base: model.Base{ID: comerAccountId}}).Updates(ComerAccount{ComerID: comerID, IsLinked: true}).Error
	}
	return db.Transaction(func(tx *gorm.DB) (err error) {
		var comer Comer
		tx.First(&comer, crtAccount.ComerID)
		var accounts []ComerAccount
		if err = db.Where("comer_id = ? and is_deleted = false", comer.ID).Find(accounts).Error; err != nil {
			return
		}
		if accounts == nil || (len(accounts) == 1 && comer.Address == nil || strings.TrimSpace(*comer.Address) == "") {
			if err = tx.Delete(&comer).Error; err != nil {
				return
			}
			if err = tx.Model(&ComerAccount{Base: model.Base{ID: comerAccountId}}).Updates(ComerAccount{ComerID: comerID, IsLinked: true}).Error; err != nil {
				return
			}
		}
		return nil
	})
}

func GetComerAccountsByComerId(db *gorm.DB, comerId uint64, accounts *[]ComerAccount) (err error) {
	return db.Where("comer_id = ? and is_deleted = false", comerId).Find(accounts).Error
}
