package tag

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//GetTagList get tag list tag ids
func GetTagList(db *gorm.DB, input ListRequest, tags *[]Tag) (count int64, err error) {
	db = db.Where("is_index = ? AND is_deleted = false", input.IsIndex)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Category != "" {
		db = db.Where("category = ?", input.Category)
	}
	err = db.Find(tags).Count(&count).Limit(input.Limit).Offset(input.Offset).Error
	return
}

//GetTagListByIDs get tag list tag ids
func GetTagListByIDs(db *gorm.DB, tagIDs []uint64, tags *[]Tag) error {
	return db.Find(tags, tagIDs).Error
}

//GetTagRelList get tag-target relation list
func GetTagRelList(db *gorm.DB, targetID uint64, target Target, comerSkillRel *[]TagTargetRel) error {
	return db.Where("target_id = ? AND target = ?", targetID, target).Find(comerSkillRel).Error
}

//FirstOrCreateTag first or create tags
func FirstOrCreateTag(db *gorm.DB, tag *Tag) error {
	return db.Where("name = ?", tag.Name).FirstOrCreate(&tag).Error
}

//DeleteTagRel delete comer skill relation where not used
func DeleteTagRel(db *gorm.DB, comerID uint64, target Target, skillIds []uint64) error {
	return db.Delete(&TagTargetRel{}, "target_id = ? AND target = ? AND tag_id NOT IN ?", comerID, target, skillIds).Error
}

//BatchCreateTagRel delete comer skill relation where not used
func BatchCreateTagRel(db *gorm.DB, tagTargetRelList []TagTargetRel) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tagTargetRelList).Error
}
