package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"ceres/pkg/utility/tool"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// GetComerProfile get current comer profile
// if profile is not exists then will let the router return 404
func GetComerProfile(uin uint64) (response *model.ComerProfileResponse, err error) {
	profile, err := model.GetComerProfile(mysql.DB, uin)
	if err != nil {
		return
	}
	// current comer is no profile saved then the router should return some code for frontend
	if profile.ID == 0 {
		return
	}

	response = &model.ComerProfileResponse{
		Name:       profile.Name,
		Location:   profile.Location,
		Website:    profile.Website,
		Bio:        profile.Bio,
		Skills:     getSkills(uin),
		Socials:    getSocials(uin),
		Wallets:    getWallets(uin),
	}
	return
}

func getSkills(comerId uint64) []string {
	var skills []string
	skillRels, err := model.GetSkillRels(mysql.DB, comerId)
	if err != nil {
		return skills
	}
	var skillIds []uint64
	for _, v := range skillRels {
		skillIds = append(skillIds, v.SkillID)
	}

	skillModels, err2 := model.GetSkillList(mysql.DB, skillIds)
	if err2 != nil {
		return skills
	}
	for _, v := range skillModels {
		skills = append(skills, v.Name)
	}
	return skills
}

func getSocials(comerId uint64) []model.SocialEntity {
	var socials []model.SocialEntity
	socialRels, err := model.GetSocialRels(mysql.DB, comerId)
	if err != nil {
		return socials
	}
	var socialIds []uint64
	for _, v := range socialRels {
		socialIds = append(socialIds, v.SocialID)
	}

	socialModels, err2 := model.GetSocialList(mysql.DB, socialIds)
	if err2 != nil {
		return socials
	}
	for _, v := range socialModels {
		entity := new(model.SocialEntity)
		entity.Type = v.Type
		entity.Account = v.Account
		socials = append(socials, *entity)
	}
	return socials
}

func getWallets(comerId uint64) []string {
	var wallets []string
	walletRels, err := model.GetWalletRels(mysql.DB, comerId)
	if err != nil {
		return wallets
	}
	var walletIds []uint64
	for _, v := range walletRels {
		walletIds = append(walletIds, v.WalletID)
	}

	walletModels, err2 := model.GetWalletList(mysql.DB, walletIds)
	if err2 != nil {
		return wallets
	}
	for _, v := range walletModels {
		wallets = append(wallets, v.Address)
	}
	return wallets
}

// CreateComerProfile  create a new profil for comer
// current comer should not be exists now
func CreateComerProfile(uin uint64, post *model.CreateProfileRequest) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		profile, err := model.GetComerProfile(tx, uin)
		if err != nil {
			return err
		}
		if profile.ID == 0 {
			now := time.Now()
			profile.Name = post.Name
			profile.Location = post.Location
			profile.Website = post.Website
			profile.Bio = post.Bio
			profile.CreateAt = now
			profile.UpdateAt = now
			if err = model.CreateComerProfile(tx, &profile); err != nil {
				return err
			}

			// bind skill
			var skillNames []string
			skillModels, _ := model.GetSkillListByNames(tx, post.Skills)
			for _, skillModel := range skillModels {
				rel := model.ComerSkillRel{}
				rel.ComerID = uin
				rel.SkillID = skillModel.ID
				model.CreateSkillRel(tx, &rel)
			}

			for _, model := range skillModels {
				skillNames = append(skillNames, model.Name)
			}
			skillDiffNames := tool.SliceDiff(post.Skills, skillNames)
			for _, name := range skillDiffNames{
				skill := model.Skill{}
				skill.Name = name
				model.CreateSkill(tx, &skill)

				rel := model.ComerSkillRel{}
				rel.ComerID = uin
				rel.SkillID = skill.ID
				model.CreateSkillRel(tx, &rel)
			}

			// bind skill
			if err = bindSkillToComer(tx, uin, post.Skills); err != nil {
				return err
			}
			// bind social
			if err = bindSocialToComer(tx, uin, post.Socials); err != nil {
				return err
			}

			// bind wallet
			if err = bindWalletToComer(tx, uin, post.Wallets); err != nil {
				return err
			}
			return nil
		}
		return errors.New("comer profile is exists now")
	})
	return
}

// UpdateComerProfile update the comer profile
// if profile is not exists then will return the not exits error
func UpdateComerProfile(uin uint64, post *model.UpdateProfileRequest) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		profile, err := model.GetComerProfile(tx, uin)
		if err != nil {
			return err
		}
		if profile.ID == 0 {
			return errors.New("comer profile is not exists")
		}

		if profile.Name != post.Name {
			profile.Name = post.Name
		}
		if profile.Location != post.Location && post.Location != "" {
			profile.Location = post.Location
		}
		if profile.Website != post.Website && post.Website != "" {
			profile.Website = post.Website
		}
		if profile.Bio != post.Bio && post.Bio != "" {
			profile.Bio = post.Bio
		}
		profile.UpdateAt = time.Now()
		err = model.UpdateComerProfile(tx, &profile)
		if err != nil {
			return err
		}
		// bind skill
		if err = bindSkillToComer(tx, uin, post.Skills); err != nil {
			return err
		}
		// bind social
		if err = bindSocialToComer(tx, uin, post.Socials); err != nil {
			return err
		}

		// bind wallet
		if err = bindWalletToComer(tx, uin, post.Wallets); err != nil {
			return err
		}
		return nil
	})
	return
}

func bindSkillToComer(db *gorm.DB, comerId uint64, skills []string) error {
	for _, name := range skills{
		skill := model.Skill{}
		skill.Name = name
		if err := model.FirstOrCreateSkill(db, &skill); err != nil {
			return err
		}

		rel := model.ComerSkillRel{}
		rel.ComerID = comerId
		rel.SkillID = skill.ID
		if err :=model.FirstOrCreateSkillRel(db, &rel); err != nil {
			return err
		}
	}
	return nil
}

func bindSocialToComer(db *gorm.DB, comerId uint64, entityList []model.SocialEntity) error {
	for _, socialEntity := range entityList{
		social := model.Social{}
		social.Account = socialEntity.Account
		social.Type = socialEntity.Type
		if err := model.FirstOrCreateSocial(db, &social); err != nil {
			return err
		}

		rel := model.ComerSocialRel{}
		rel.ComerID = comerId
		rel.SocialID = social.ID
		if err :=model.FirstOrCreateSocialRel(db, &rel); err != nil {
			return err
		}
	}
	return nil
}

func bindWalletToComer(db *gorm.DB, comerId uint64, addresss []string) error {
	for _, address := range addresss{
		wallet := model.Wallet{}
		wallet.Address = address
		if err := model.FirstOrCreateWallet(db, &wallet); err != nil {
			return err
		}

		rel := model.ComerWalletRel{}
		rel.ComerID = comerId
		rel.WalletID = wallet.ID
		if err :=model.FirstOrCreateWalletRel(db, &rel); err != nil {
			return err
		}
	}
	return nil
}
