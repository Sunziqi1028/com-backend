package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
	"fmt"
	"gorm.io/gorm"
	"math"
)

// TODO: bounty model

func CreateBounty(db *gorm.DB, bounty *Bounty) (uint64, error) {
	if err := db.Create(&bounty).Error; err != nil {
		return 0, err
	}
	return bounty.ID, nil
}

func CreateContact(db *gorm.DB, contact *BountyContact) error {
	return db.Create(&contact).Error
}

func CreateDeposit(db *gorm.DB, deposit *BountyDeposit) error {
	return db.Create(&deposit).Error
}

func CreatePaymentTerms(db *gorm.DB, paymentTerm *BountyPaymentTerms) error {
	return db.Create(&paymentTerm).Error
}

func CreatePaymentPeriod(db *gorm.DB, paymentPeriod *BountyPaymentPeriod) error {
	return db.Create(&paymentPeriod).Error
}

func UpdateBountyDepositContract(db *gorm.DB, bountyID uint64, depositContract string) error {
	return db.Model(&Bounty{}).Where("id = ?", bountyID).Update("deposit_contract", depositContract).Error
}

func GetAndUpdateTagID(db *gorm.DB, name string) (tagID uint64, err error) {
	err = db.Table("tag").Select("id").Where("name = ? and 'category' = 'comerSkill' ", name).Find(&tagID).Error
	if err != nil {
		return 0, err
	}

	if tagID == 0 {
		var isIndex bool
		if len(name) > 2 && name[0:1] == "#" {
			isIndex = true
		}
		skill := tag.Tag{
			Name:     name,
			IsIndex:  isIndex,
			Category: tag.Bounty,
		}
		tag.FirstOrCreateTag(db, &skill)
		return skill.ID, nil
	}
	return tagID, nil
}

func CreateTagTargetRel(db *gorm.DB, tagTargetRel *tag.TagTargetRel) error {
	return db.Create(&tagTargetRel).Error
}

// GetPaymentTermsByBountyId get payment_terms list
func GetPaymentTermsByBountyId(db *gorm.DB, bountyId uint64, termList *[]BountyPaymentTerms) error {
	return db.Model(&BountyPaymentTerms{}).Where("bounty_id = ? ", bountyId).Find(termList).Error
}

func GetPaymentPeriodsByBountyId(db *gorm.DB, bountyId uint64, termList *[]BountyPaymentPeriod) error {
	return db.Model(&BountyPaymentPeriod{}).Where("bounty_id = ? ", bountyId).Find(termList).Error
}

func GetBountyTagNames(db *gorm.DB, bountyId uint64) (tagNames []string, err error) {
	var tagIds []uint64
	if err := db.Model(&tag.TagTargetRel{}).Where("target= ? and target_id = ?", "bounty", bountyId).Select("tag_id").Find(&tagIds).Error; err != nil {
		return nil, err
	}
	if len(tagIds) >= 0 {
		if err := db.Model(&tag.Tag{}).Where("id in ?", tagIds).Select("name").Find(&tagNames).Error; err != nil {
			return nil, err
		}
	}
	return
}

func GetApplicantCountOfBounty(db *gorm.DB, bountyId uint64) (cnt int64, err error) {
	if err := db.Model(&BountyApplicant{}).Where("bounty_id = ?", bountyId).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func GetApplicantByBountyAndComer(db *gorm.DB, bountyId uint64, comerId uint64) (app BountyApplicant, err error) {
	if err := db.Model(&BountyApplicant{}).Where("bounty_id = ? and comer_id = ?", bountyId).Find(&app).Error; err != nil {
		return BountyApplicant{}, err
	}
	return app, nil
}

func GetBountyDepositByBountyAndComer(db *gorm.DB, bountyID uint64, crtComerId uint64) (bd BountyDeposit, err error) {
	if err := db.Model(&BountyDeposit{}).Where("bounty_id = ? and comer_id = ?", bountyID, crtComerId).Find(&bd).Error; err != nil {
		return BountyDeposit{}, err
	}
	return bd, nil
}
func PageSelectBounties(db *gorm.DB, pagination model.Pagination) (*model.Pagination, error) {
	var bounties []*Bounty
	if err := db.Scopes(model.Paginate(&Bounty{}, &pagination, db)).Find(&bounties).Error; err != nil {
		return nil, err
	}
	pagination.Rows = bounties
	return &pagination, nil
}

func PageSelectBountiesByStartupId(db *gorm.DB, pagination model.Pagination, startupId uint64) (*model.Pagination, error) {
	var bounties []*Bounty
	if err := db.Scopes(model.Paginate(&Bounty{}, &pagination, db)).Where("startup_id = ?", startupId).Find(&bounties).Error; err != nil {
		return nil, err
	}
	pagination.Rows = bounties
	return &pagination, nil
}

func PageSelectPostedBounties(db *gorm.DB, pagination model.Pagination, comerId uint64) (*model.Pagination, error) {
	var bounties []*Bounty
	if err := db.Scopes(model.Paginate(&Bounty{}, &pagination, db)).Where("comer_id = ?", comerId).Find(&bounties).Error; err != nil {
		return nil, err
	}
	pagination.Rows = bounties
	return &pagination, nil
}

func PageSelectParticipatedBounties(db *gorm.DB, pagination model.Pagination, comerId uint64) (*model.Pagination, error) {
	var bounties []Bounty
	var countSql = fmt.Sprintf("select count(b.id) from bounty t left join bounty_applicant ba on b.id = ba.bounty_id where ba.comer_id=%d and ba.status not in (4,5)", comerId)
	var cnt int64
	if err := db.Raw(countSql).Scan(&cnt).Error; err != nil {
		return &pagination, err
	}
	var sql = fmt.Sprintf("select b.* from bounty t left join bounty_applicant ba on b.id = ba.bounty_id where ba.comer_id=%d and ba.status not in (4,5) order by b.created_at desc limit %d,%d", comerId, pagination.GetLimit(), pagination.GetOffset())
	if err := db.Raw(sql).Scan(&bounties).Error; err != nil {
		return &pagination, err
	}

	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return &pagination, nil
}
