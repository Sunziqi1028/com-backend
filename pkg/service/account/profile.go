package account

import (
	model "ceres/pkg/model/account"
)

// GetComerProfile get current comer profile
// if profile is not exists then will let the router return 404
func GetComerProfile(uin uint64) (response *model.ComerProfileResponse, err error) {
	//profile, err := model.GetComerProfile(mysql.DB, uin)
	//if err != nil {
	//	return
	//}
	//// current comer is no profile saved then the router should return some code for frontend
	//if profile.ID == 0 {
	//	return
	//}
	//skillIds := make([]uint64, 0)
	//for _, v := range strings.Split(profile.Skills, ",") {
	//	id, _ := strconv.ParseInt(v, 10, 64)
	//	skillIds = append(skillIds, uint64(id))
	//}
	//tags, err := model.GetSkillList(mysql.DB, skillIds)
	//var skills []string
	//for _, v := range tags {
	//	skills = append(skills, v.Name)
	//}
	//response = &model.ComerProfileResponse{
	//	Skills:      skills,
	//	About:       profile.About,
	//	Description: profile.Description,
	//	Email:       profile.Email,
	//}
	return
}

// CreateComerProfile  create a new profil for comer
// current comer should not be exists now
func CreateComerProfile(uin uint64, post *model.CreateProfileRequest) (err error) {
	//err = mysql.DB.Transaction(func(tx *gorm.DB) error {
	//	profile, err := model.GetComerProfile(tx, uin)
	//	if err != nil {
	//		return err
	//	}
	//	if profile.ID == 0 {
	//		now := time.Now()
	//		profile.About = post.About
	//		profile.Description = post.Description
	//		profile.Email = post.Email
	//		profile.Identifier = utility.ProfileSequence.Next()
	//		profile.CreateAt = now
	//		profile.UpdateAt = now
	//		var skillIds []string
	//		for _, id := range post.SKills {
	//			skillIds = append(skillIds, strconv.FormatInt(int64(id), 10))
	//		}
	//		profile.Skills = strings.Join(skillIds, ",")
	//		err = model.CreateComerProfile(tx, &profile)
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//	return errors.New("comer profile is exists now")
	//})
	return
}

// UpdateComerProfile update the comer profile
// if profile is not exists then will return the not exits error
func UpdateComerProfile(uin uint64, post *model.UpdateProfileRequest) (err error) {
	//err = mysql.DB.Transaction(func(tx *gorm.DB) error {
	//	profile, err := model.GetComerProfileByIdentifier(tx, post.Identifier)
	//	if err != nil {
	//		return err
	//	}
	//	if profile.ID == 0 {
	//		return errors.New("comer profile is not exists")
	//	}
	//
	//	var skillIds []string
	//	for _, id := range post.SKills {
	//		skillIds = append(skillIds, strconv.FormatInt(int64(id), 10))
	//	}
	//	skills := strings.Join(skillIds, ",")
	//	if profile.Skills != skills {
	//		profile.Skills = skills
	//	}
	//	if profile.Email != post.Email && post.Email != "" {
	//		profile.Email = post.Email
	//	}
	//	if profile.Description != post.Description && post.Description != "" {
	//		profile.Description = post.Description
	//	}
	//	if profile.About != post.About && post.About != "" {
	//		profile.About = post.About
	//	}
	//	profile.UpdateAt = time.Now()
	//	err = model.UpdateComerProfile(tx, &profile)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	return
}
