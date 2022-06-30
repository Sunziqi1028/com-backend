package bounty

import (
	"gorm.io/gorm"
)

// TODO: bounty model

func GetComerStartups(db *gorm.DB, comerID uint64, startups *GetStartupsResponse) (*GetStartupsResponse, error) {
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

func CreateTransaction(db *gorm.DB, chainInfo *Transaction) error {
	return db.Create(&chainInfo).Error
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
	return db.Model(&Bounty{}).Where("bounty_id = ?", bountyID).Update("deposit_contract", depositContract).Error
}
