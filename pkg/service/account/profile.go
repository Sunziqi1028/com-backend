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
	if err = model.GetComerProfile(mysql.DB, comerID, &response.ComerProfile); err != nil {
		log.Warn(err)
		return err
	}
	var accounts []model.ComerAccount
	_ = model.GetComerAccountsByComerId(mysql.DB, comerID, &accounts)
	var accountInfos []model.ComerAccountInfo
	if accounts != nil {
		for _, account := range accounts {
			accountInfos = append(accountInfos, model.ComerAccountInfo{
				ComerAccountId:   account.ID,
				ComerAccountType: account.Type,
			})
		}
	}
	response.ComerAccounts = accountInfos
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
	if post.Twitter == nil {
		post.Twitter = new(string)
	}
	if post.Discord == nil {
		post.Discord = new(string)
	}
	if post.Telegram == nil {
		post.Telegram = new(string)
	}
	if post.Medium == nil {
		post.Medium = new(string)
	}
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		TimeZone: *post.TimeZone,
		Website:  post.Website,
		Email:    *post.Email,
		Twitter:  *post.Twitter,
		Discord:  *post.Discord,
		Telegram: *post.Telegram,
		Medium:   *post.Medium,
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
				Name:     skillName,
				IsIndex:  isIndex,
				Category: tag.ComerSkill,
			}
			if er = tag.FirstOrCreateTag(tx, &skill); err != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkill,
				TargetID: comerID,
			})
		}
		//batch create comer skill relation
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			log.Warn(er)
			return
		}
		//create comer profile
		if er = model.CreateComerProfile(tx, &profile); er != nil {
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
		TimeZone: *post.TimeZone,
		Website:  post.Website,
		Email:    *post.Email,
		Twitter:  *post.Twitter,
		Discord:  *post.Discord,
		Telegram: *post.Telegram,
		Medium:   *post.Medium,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) { //create skill
		for _, skillName := range post.SKills {
			var isIndex bool
			if len(skillName) > 1 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:     skillName,
				Category: tag.ComerSkill,
				IsIndex:  isIndex,
			}
			if er = tag.FirstOrCreateTag(tx, &skill); er != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkill,
				TargetID: comerID,
			})
			tagIds = append(tagIds, skill.ID)
		}
		//delete not used skills
		if er = tag.DeleteTagRel(tx, comerID, tag.ComerSkill, tagIds); er != nil {
			return er
		}
		//batch create comer skill rel
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			return er
		}
		//create profile
		if er = model.UpdateComerProfile(tx, &profile); er != nil {
			return er
		}
		return nil
	})
	return
}
