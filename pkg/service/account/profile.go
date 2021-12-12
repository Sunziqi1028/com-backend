package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"errors"

	"gorm.io/gorm"
)

// GetComerProfile get current comer profile
// if profile is not exists then will let the router return 404
func GetComerProfile(comerID uint64) (response *model.ComerProfileResponse, err error) {
	profile, err := model.GetComerProfile(mysql.DB, comerID)
	if err != nil {
		return
	}
	// current comer is no profile saved then the router should return some code for frontend
	if profile.ID == 0 {
		return
	}

	skillRels, err := model.GetSkillRelListByComerID(mysql.DB, comerID)
	if err != nil {
		return nil, err
	}
	skillIds := make([]uint64, 0)
	for _, skillRel := range skillRels {
		skillIds = append(skillIds, skillRel.SkillID)
	}
	skills, err := model.GetSkillByIds(mysql.DB, skillIds)
	if err != nil {
		return nil, err
	}
	response = &model.ComerProfileResponse{
		ComerProfile: profile,
		Skills:       skills,
	}
	return
}

// CreateComerProfile  create a new profil for comer
// current comer should not be exists now
func CreateComerProfile(comerID uint64, post *model.CreateProfileRequest) (err error) {
	profile, err := model.GetComerProfile(mysql.DB, comerID)
	if err != nil {
		return err
	}
	if profile.ID != 0 {
		return errors.New("comer profile has exist")
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
	profile, err := model.GetComerProfile(mysql.DB, comerID)
	if err != nil {
		return err
	}
	if profile.ID == 0 {
		return errors.New("comer does not exist")
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
