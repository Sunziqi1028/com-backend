package crowdfunding

import (
	"ceres/pkg/model/startup"
	"gorm.io/gorm"
)

func CreateCrowdfunding(db *gorm.DB, c *Crowdfunding) error {
	return db.Create(c).Error
}

func CreateCrowdfundingSwap(db *gorm.DB, cs *CrowdfundingSwap) error {
	return db.Create(cs).Error
}
func SelectOnGoingByStartupId(db *gorm.DB, startupId uint64) (crowdfundingList []Crowdfunding, err error) {
	if err = db.Model(&Crowdfunding{}).Where("startup_id = ? and is_deleted = 0 and status in (0, 1, 2)", startupId).Find(&crowdfundingList).Error; err != nil {
		return
	}
	return crowdfundingList, nil
}
func SelectStartupsWithNonCrowdfundingOnGoing(db *gorm.DB, comerId uint64) (startups []CrowdfundableStartup, err error) {
	var sts []startup.Startup
	var onGoingByComer []uint64
	err = db.Model(&Crowdfunding{}).Select("startup_id").Where("comer_id=? and is_deleted=0 and status in (0, 1, 2)", comerId).Find(&onGoingByComer).Error
	if err != nil {
		return
	}
	err = db.Model(&startup.Startup{}).Where("comer_id = ? and is_deleted = 0 ", comerId).Find(&sts).Error

	if len(sts) > 0 {
		for _, st := range sts {
			can := true
			if len(onGoingByComer) > 0 {
				for _, u := range onGoingByComer {
					if u == st.ID {
						can = false
					}
				}
			}
			startups = append(startups, CrowdfundableStartup{st.ID, st.Name, can, st.TokenContractAddress})
		}
	}
	return
}

func UpdateCrowdfundingContractAddressAndStatus(db *gorm.DB, fundingID uint64, address string, status CrowdfundingStatus) error {
	return db.Model(&Crowdfunding{}).Where("id = ?", fundingID).Updates(map[string]interface{}{"crowdfunding_contract": address, "status": status}).Error
}
