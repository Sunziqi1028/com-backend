package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
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
	var comerSkillRel []model.ComerSkillRel
	if err = model.GetSkillRelListByComerID(mysql.DB, comerID, &comerSkillRel); err != nil {
		log.Warn(err)
		return err
	}
	//get skills
	skills := make([]model.Skill, 0)
	if len(comerSkillRel) > 0 {
		skillIds := make([]uint64, 0)
		for _, skillRel := range comerSkillRel {
			skillIds = append(skillIds, skillRel.SkillID)
		}
		if err = model.GetSkillByIds(mysql.DB, skillIds, &skills); err != nil {
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
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if profile.ID != 0 {
		return router.ErrBadRequest.WithMsg("user profile already exists")
	}
	var comerSkillRelList []model.ComerSkillRel
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		Website:  post.Website,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//create skill
		for _, skillName := range post.SKills {
			skill := model.Skill{
				Name: skillName,
			}
			if err = model.FirstOrCreateSkill(tx, &skill); err != nil {
				return err
			}
			comerSkillRelList = append(comerSkillRelList, model.ComerSkillRel{
				ComerID: comerID,
				SkillID: skill.ID,
			})
		}
		//batch create comer skill relation
		if err = model.BatchCreateComerSkillRel(tx, comerSkillRelList); err != nil {
			return err
		}
		//create comer profile
		if err = model.CreateComerProfile(tx, &profile); err != nil {
			return err
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
	var skillIds []uint64
	var comerSkillRelList []model.ComerSkillRel
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		Website:  post.Website,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//create skill
		for _, skillName := range post.SKills {
			skill := model.Skill{
				Name: skillName,
			}
			if err = model.FirstOrCreateSkill(tx, &skill); err != nil {
				return err
			}
			comerSkillRelList = append(comerSkillRelList, model.ComerSkillRel{
				ComerID: comerID,
				SkillID: skill.ID,
			})
			skillIds = append(skillIds, skill.ID)
		}
		//delete not used skills
		if err = model.DeleteComerSkillRelByNotIds(tx, comerID, skillIds); err != nil {
			return err
		}
		//batch create comer skill rel
		if err = model.BatchCreateComerSkillRel(tx, comerSkillRelList); err != nil {
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
