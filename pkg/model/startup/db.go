package startup

import "gorm.io/gorm"

// GetStartup  get startup
func GetStartup(db *gorm.DB, startupID uint64, startup *Startup) error {
	return db.Where("is_deleted = false AND id = ?", startupID).Preload("Wallets").Find(&startup).Error
}

// ListStartups  list startups
func ListStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false")
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	} else {
		db = db.Where("is_set = true")
	}
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != "" {
		db = db.Where("mode = ?", input.Mode)
	}
	err = db.Preload("Wallets").Order("created_at DESC").Find(startups).Count(&total).Limit(input.Limit).Offset(input.Offset).Error
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
	if input.Mode != "" {
		db = db.Where("mode = ?", input.Mode)
	}
	err = db.Preload("Wallets").Order("created_at DESC").Find(startups).Count(&total).Limit(input.Limit).Offset(input.Offset).Error
	return
}
