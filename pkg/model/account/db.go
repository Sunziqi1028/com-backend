package account

import (
	"ceres/pkg/model"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

// GetComerByAddress  get comer entity by comer's address
func GetComerByAddress(db *gorm.DB, address string, comer *Comer) error {
	return db.Where("address = ?", address).Find(comer).Error
}

// GetComerByID  get comer entity by comer's ID
func GetComerByID(db *gorm.DB, comerID uint64, comer *Comer) (err error) {
	return db.Where("id = ?", comerID).Find(comer).Error
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
	return db.Where("type = ? AND oin = ?", accountType, oin).Find(comerAccount).Error
}

func ListAccount(db *gorm.DB, comerID uint64, accountList *[]ComerAccount) (err error) {
	return db.Where("comer_id = ? ", comerID).Find(accountList).Error
}

func CreateAccount(db *gorm.DB, comerAccount *ComerAccount) (err error) {
	return db.Create(comerAccount).Error
}

func DeleteAccount(db *gorm.DB, comerID, accountID uint64) error {
	return db.Where("comer_id = ? AND id = ?", comerID, accountID).Delete(&ComerAccount{}).Error
}

//GetComerProfile update the comer address
func GetComerProfile(db *gorm.DB, comerID uint64, profile *ComerProfile) (err error) {
	return db.Where("comer_id = ?", comerID).Find(profile).Error
}

//CreateComerProfile update the comer address
func CreateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Create(&comerProfile).Error
}

//UpdateComerProfile update the comer address
func UpdateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Where("comer_id = ?", comerProfile.ComerID).Updates(comerProfile).Error
}

//GetSkillByIds first or create skills
func GetSkillByIds(db *gorm.DB, skillIDs []uint64, skills *[]Skill) error {
	return db.Find(skills, skillIDs).Error
}

//GetSkillRelListByComerID first or create skills
func GetSkillRelListByComerID(db *gorm.DB, comerID uint64, comerSkillRel *[]ComerSkillRel) error {
	return db.Where("comer_id = ?", comerID).Find(comerSkillRel).Error
}

//FirstOrCreateSkill first or create skills
func FirstOrCreateSkill(db *gorm.DB, skill *Skill) error {
	return db.Where(skill).FirstOrCreate(&skill).Error
}

//DeleteComerSkillRelByNotIds delete comer skill relation where not used
func DeleteComerSkillRelByNotIds(db *gorm.DB, comerID uint64, skillIds []uint64) error {
	return db.Delete(&ComerSkillRel{}, "comer_id = ? AND skill_id NOT IN ?", comerID, skillIds).Error
}

func BatchCreateComerSkillRel(db *gorm.DB, ComerSkillRelList []ComerSkillRel) error {
	db = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&ComerSkillRelList)
	return db.Error
}
