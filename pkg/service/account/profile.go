package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	"errors"

	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// GetComerProfile get current comer profile
func GetComerProfile(comerID uint64, response *model.ComerProfileResponse) (err error) {
	//get comer profile
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return err
	}
	if profile.ID == 0 {
		return router.ErrNotFound.WithMsg("user profile does not exists")
	}
	//get comer profile skill relations
	var tagRelList []tag.TagTargetRel
	if err = tag.GetTagRelList(mysql.DB, comerID, tag.ComerSkillTag, &tagRelList); err != nil {
		log.Warn(err)
		return err
	}
	//get skills
	skills := make([]tag.Tag, 0)
	if len(tagRelList) > 0 {
		skillIds := make([]uint64, 0)
		for _, skillRel := range tagRelList {
			skillIds = append(skillIds, skillRel.TagID)
		}
		if err = tag.GetTagListByIDs(mysql.DB, skillIds, &skills); err != nil {
			return err
		}
	}

	*response = model.ComerProfileResponse{
		ComerProfile: profile,
		Skills:       skills,
	}

	return
}

// CreateComerProfile  create a new profil for comer
// current comer should not be exists now
func CreateComerProfile(comerID uint64, post *model.CreateProfileRequest) (err error) {
	//get comer profile
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return
	}
	if profile.ID != 0 {
		return router.ErrBadRequest.WithMsg("user profile already exists")
	}
	var tagRelList []tag.TagTargetRel
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		Website:  post.Website,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		//create skill
		for _, skillName := range post.SKills {
			var isIndex bool
			if len(skillName) > 2 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:    skillName,
				IsIndex: isIndex,
			}
			if err = tag.FirstOrCreateTag(tx, &skill); err != nil {
				return err
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkillTag,
				TargetID: comerID,
			})
		}
		//batch create comer skill relation
		if err = tag.BatchCreateTagRel(tx, tagRelList); err != nil {
			log.Warn(er)
			return
		}
		//create comer profile
		if err = model.CreateComerProfile(tx, &profile); err != nil {
			log.Warn(er)
			return
		}
		return nil
	})
	return err
}

// UpdateComerProfile update the comer profile
// if profile is not exists then will return the not exits error
func UpdateComerProfile(comerID uint64, post *model.UpdateProfileRequest) (err error) {
	//get comer profile
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if profile.ID == 0 {
		return router.ErrBadRequest.WithMsg("user profile does not exists")
	}
	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		Website:  post.Website,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error { //create skill
		for _, skillName := range post.SKills {
			var isIndex bool
			if len(skillName) > 1 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:    skillName,
				IsIndex: isIndex,
			}
			if err = tag.FirstOrCreateTag(tx, &skill); err != nil {
				return err
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkillTag,
				TargetID: comerID,
			})
			tagIds = append(tagIds, skill.ID)
		}
		//delete not used skills
		if err = tag.DeleteTagRel(tx, comerID, tag.ComerSkillTag, tagIds); err != nil {
			return err
		}
		//batch create comer skill rel
		if err = tag.BatchCreateTagRel(tx, tagRelList); err != nil {
			return err
		}
		//create profile
		if err = model.UpdateComerProfile(tx, &profile); err != nil {
			return err
		}
		return nil
	})
	return
}
