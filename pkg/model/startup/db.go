package startup

import (
	"ceres/pkg/model/tag"
	"database/sql"

	"gorm.io/gorm"
)

// GetStartup  get startup
func GetStartup(db *gorm.DB, startupID uint64, startup *Startup) error {
	return db.Where("is_deleted = false AND id = ?", startupID).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Follows").Find(&startup).Error
}

// CreateStartup  create startup
func CreateStartup(db *gorm.DB, startup *Startup) (err error) {
	return db.Create(startup).Error
}

// CreateStartupWallet  create startup wallet
func CreateStartupWallet(db *gorm.DB, wallets []Wallet) (err error) {
	return db.Create(&wallets).Error
}

// BatchUpdateStartupWallet  batch update startup wallets
func BatchUpdateStartupWallet(db *gorm.DB, wallets []Wallet) (err error) {
	return db.Save(&wallets).Error
}

//FirstOrCreateWallet first or create wallet
func FirstOrCreateWallet(db *gorm.DB, wallet *Wallet) error {
	return db.Where("startup_id = ? AND wallet_name = ? ", wallet.StartupID, wallet.WalletName).FirstOrCreate(&wallet).Error
}

//DeleteStartupWallet delete startup wallet where not used
func DeleteStartupWallet(db *gorm.DB, startupID uint64, walletIds []uint64) error {
	return db.Delete(&Wallet{}, "startup_id = ? AND id NOT IN ?", startupID, walletIds).Error
}

// ListStartups  list startups
func ListStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false")
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Follows").Find(startups).Error
	return
}

// CreateStartupFollowRel create comer relation for startup and comer
func CreateStartupFollowRel(db *gorm.DB, comerID, startupID uint64) error {
	return db.Create(&FollowRelation{ComerID: comerID, StartupID: startupID}).Error
}

// DeleteStartupFollowRel delete comer relation for startup and comer
func DeleteStartupFollowRel(db *gorm.DB, input *FollowRelation) error {
	return db.Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Delete(input).Error
}

// ListFollowedStartups  list followed startups
func ListFollowedStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_follow_rel ON startup_follow_rel.comer_id = ? AND startup_id = startup.id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Follows").Find(startups).Error
	return
}

// StartupNameIsExist check startup's  name is existed
func StartupNameIsExist(db *gorm.DB, name string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("name = ?", name).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// StartupTokenContractIsExist check startup's  token contract is existed
func StartupTokenContractIsExist(db *gorm.DB, tokenContract string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("token_contract_address = ?", tokenContract).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// UpdateStartupBasicSetting  update startup security and social setting
func UpdateStartupBasicSetting(db *gorm.DB, startupID uint64, input *BasicSetting) (err error) {
	return db.Table("startup").Where("id = ?", startupID).Select("kyc", "contract_audit", "website", "discord", "twitter", "telegram", "docs").Updates(input).Error
}

// UpdateStartupFinanceSetting  update startup finance setting
func UpdateStartupFinanceSetting(db *gorm.DB, startupID uint64, input *FinanceSetting) (err error) {
	var fieldMap map[string]interface{}
	fieldMap = make(map[string]interface{})
	fieldMap["token_contract_address"] = input.TokenContractAddress
	fieldMap["launch_network"] = input.LaunchNetwork
	fieldMap["token_name"] = input.TokenName
	fieldMap["token_symbol"] = input.TokenSymbol
	fieldMap["total_supply"] = input.TotalSupply
	if input.PresaleStart.IsZero() {
		fieldMap["presale_start"] = sql.NullTime{}
	} else {
		fieldMap["presale_start"] = input.PresaleStart
	}
	if input.PresaleEnd.IsZero() {
		fieldMap["presale_end"] = sql.NullTime{}
	} else {
		fieldMap["presale_end"] = input.PresaleEnd
	}
	if input.LaunchDate.IsZero() {
		fieldMap["launch_date"] = sql.NullTime{}
	} else {
		fieldMap["launch_date"] = input.LaunchDate
	}

	return db.Table("startup").Where("id = ?", startupID).Updates(fieldMap).Error
}

// ListParticipatedStartups  list participated startups
func ListParticipatedStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_team_member_rel ON startup_team_member_rel.comer_id = ? AND startup_id = startup.id AND startup.comer_id != startup_team_member_rel.comer_id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Follows").Find(startups).Error
	return
}

// ListBeMemberStartups  list I am a member of startups
func ListBeMemberStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_team_member_rel ON startup_team_member_rel.comer_id = ? AND startup_id = startup.id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	return
}

// StartupFollowIsExist check startup and comer is existed
func StartupFollowIsExist(db *gorm.DB, startupID, comerID uint64) (isExist bool, err error) {
	var count int64
	err = db.Table("startup_follow_rel").Where("startup_id = ? AND comer_id = ?", startupID, comerID).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExist = false
	} else {
		isExist = true
	}
	return
}
