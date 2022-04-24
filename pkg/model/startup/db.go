package startup

import (
	"ceres/pkg/model/tag"

	"gorm.io/gorm"
)

// GetStartup  get startup
func GetStartup(db *gorm.DB, startupID uint64, startup *Startup) error {
	return db.Where("is_deleted = false AND id = ?", startupID).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Find(&startup).Error
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
	return db.Where("startup_id = ? AND (wallet_name = ? OR wallet_address = ?)", wallet.StartupID, wallet.WalletName, wallet.WalletAddress).FirstOrCreate(&wallet).Error
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
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Find(startups).Error
	return
}

// CreateStartupFollowRel create comer relation for startup and comer
func CreateStartupFollowRel(db *gorm.DB, comerID, startupID uint64) error {
	return db.Create(&FollowRelation{ComerID: comerID, StartupID: startupID}).Error
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
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Find(startups).Error
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
func UpdateStartupBasicSetting(db *gorm.DB, startupID uint64, input *StartupBasicSetting) (err error) {
	return db.Table("startup").Where("id = ?", startupID).Select("kyc", "contract_audit", "website", "discord", "twitter", "telegram", "docs").Updates(input).Error
}

// UpdateStartupFinanceSetting  update startup finance setting
func UpdateStartupFinanceSetting(db *gorm.DB, startupID uint64, input *StartupFinanceSetting) (err error) {
	return db.Table("startup").Where("id = ?", startupID).Select("token_contract_address", "presale_date", "launch_date").Updates(input).Error
}
