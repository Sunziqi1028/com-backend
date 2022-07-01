package bounty

import (
	"ceres/pkg/model/tag"
	"gorm.io/gorm"
)

// TODO: bounty model

func GetComerStartups(db *gorm.DB, comerID uint64, startups []*GetStartupsResponse) ([]*GetStartupsResponse, error) {
	err := db.Table("startup").Select("id, name").Where("comer_id = ? and is_deleted = 0", comerID).Order("convert(name using gbk)").Find(&startups).Error
	if err != nil {
		return nil, err
	}
	return startups, nil
}

func CreateBounty(db *gorm.DB, bounty *Bounty) (uint64, error) {
	if err := db.Create(&bounty).Error; err != nil {
		return 0, err
	}
	return bounty.ID, nil
}

func CreateTransaction(db *gorm.DB, transaction *Transaction) error {
	return db.Create(&transaction).Error
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

func CreatePostUpdate(db *gorm.DB, postUpdate *PostUpdate) error {
	return db.Create(&postUpdate).Error
}

func UpdateBountyDepositContract(db *gorm.DB, bountyID uint64, depositContract string) error {
	return db.Model(&Bounty{}).Where("id = ?", bountyID).Update("deposit_contract", depositContract).Error
}

func UpdateTransactionStatus(db *gorm.DB, bountyID uint64, status uint64) error {
	return db.Model(&Transaction{}).Where("bounty_id = ?", bountyID).Update("status", int(status)).Error
}

func GetAndUpdateTagID(db *gorm.DB, name string) (tagID uint64, err error) {
	err = db.Table("tag").Select("id").Where("name = ? and catagory = comerSkill", name).Find(&tagID).Error
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

func GetTransaction(db *gorm.DB) (transactionResponse []*GetTransactions, err error) {
	err = db.Table("transcation").Select("chain_id, tx_hash, source_id").Where("status = 0").Find(&transactionResponse).Error
	if err != nil {
		return nil, err
	}
	return transactionResponse, nil
}
