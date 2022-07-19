package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
	"fmt"
	"gorm.io/gorm"
	"math"
)

// TODO: bounty model

const (
	AccessIn                       = 1
	AccessOut                      = 2
	PaymentModeStage               = 1
	PaymentModePeriod              = 2
	BountyPaymentTermsStatusUnpaid = 1
	BountyPaymentTermsStatusPaid   = 2
	BountyPaymentTermsPeriodSeqNum = 1
	BountyStatusReadyToWork        = 1
	BountyStatusWordStarted        = 2
	BountyStatusCompleted          = 3
	BountyStatusExpired            = 4
	ApplicantStatusApplied         = 1
	ApplicantStatusApproved        = 2
	ApplicantStatusSubmitted       = 3
	ApplicantStatusRevoked         = 4
	ApplicantStatusRejected        = 5
	ApplicantStatusQuited          = 6
)

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

func UpdateBountyDepositStatus(db *gorm.DB, bountyID uint64, status uint64) error {
	return db.Model(&BountyDeposit{}).Where("bounty_id = ?", bountyID).Update("status", status).Error
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
func PageSelectOnChainBounties(db *gorm.DB, pagination model.Pagination) (*model.Pagination, error) {
	var bounties []*Bounty

	cntSql := fmt.Sprintf("select count(b.id) from bounty b left join bounty_deposit bd on b.id=bd.bounty_id and b.comer_id=bd.comer_id where bd.status=1")
	var cnt int64
	if err := db.Raw(cntSql).Scan(&cnt).Error; err != nil {
		return &pagination, err
	}
	pageSql := fmt.Sprintf("select b.* from bounty b left join bounty_deposit bd on b.id=bd.bounty_id and b.comer_id=bd.comer_id where bd.status=1 order by %s limit %d,%d", pagination.Sort, pagination.GetOffset(), pagination.Limit)

	if err := db.Raw(pageSql).Scan(&bounties).Error; err != nil {
		return &pagination, err
	}
	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return &pagination, nil
}

func PageSelectBountiesByStartupId(db *gorm.DB, pagination model.Pagination, startupId uint64) (*model.Pagination, error) {
	var bounties []*Bounty

	cntSql := fmt.Sprintf("select count(b.id) from bounty b left join bounty_deposit bd on b.id=bd.bounty_id and b.comer_id=bd.comer_id where bd.status=1 and b.startup_id=%d", startupId)
	var cnt int64
	if err := db.Raw(cntSql).Scan(&cnt).Error; err != nil {
		return &pagination, err
	}
	pageSql := fmt.Sprintf("select b.* from bounty b left join bounty_deposit bd on b.id=bd.bounty_id and b.comer_id=bd.comer_id where bd.status=1 and b.startup_id=%d order by %s limit %d,%d", startupId, pagination.Sort, pagination.GetOffset(), pagination.Limit)

	if err := db.Raw(pageSql).Scan(&bounties).Error; err != nil {
		return &pagination, err
	}
	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
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
	var countSql = fmt.Sprintf("select count(b.id) from bounty b left join bounty_applicant ba on b.id = ba.bounty_id where ba.comer_id=%d and ba.status not in (4,5)", comerId)
	var cnt int64
	if err := db.Raw(countSql).Scan(&cnt).Error; err != nil {
		return &pagination, err
	}
	var sql = fmt.Sprintf("select b.* from bounty b left join bounty_applicant ba on b.id = ba.bounty_id where ba.comer_id=%d and ba.status not in (4,5) order by b.created_at desc limit %d,%d", comerId, pagination.GetOffset(), pagination.GetLimit())
	if err := db.Raw(sql).Scan(&bounties).Error; err != nil {
		return &pagination, err
	}

	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return &pagination, nil
}

func GetDetailByBountyID(db *gorm.DB, bountyID uint64) (*DetailResponse, error) {
	var detailResponse DetailResponse
	var sql = fmt.Sprintf("select title, status, discussion_link, apply_cutoff_date, applicant_deposit, description, created_at from bounty where id = %d", bountyID)
	err := db.Raw(sql).Scan(&detailResponse).Error
	if err != nil {
		return nil, err
	}
	var tagIds []uint64
	err = db.Table("tag_target_rel").Select("tag_id").Where("target_id = ?", bountyID).Find(&tagIds).Error
	if err != nil {
		return nil, err
	}
	var skillNames []string
	var skillName string
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ? and category = bounty", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	var contacts []Contact
	db.Table("bounty_contact").Select("contact_type, contact_address").Where("bounty_id = ?", bountyID).Find(&contacts)
	detailResponse.ApplicantSkills = skillNames
	detailResponse.Contacts = contacts
	return &detailResponse, nil
}

func GetPaymentByBountyID(db *gorm.DB, bountyID uint64) (*PaymentResponse, error) {
	var paymentResponse PaymentResponse
	var comerID uint64
	err := db.Table("bounty").Select("comer_id").Where("id = ? and status != 0", bountyID).Find(&comerID).Error
	if err != nil {
		return nil, err
	}
	err = db.Table("bounty").Select("payment_mode, founder_deposit").Where("id = ? and status != 0", bountyID).Find(&paymentResponse).Error
	if err != nil {
		return nil, err
	}
	if paymentResponse.PaymentMode == PaymentModeStage {
		var StagePayments []StageTerm
		db.Table("bounty_payment_terms").Select("seq_num, status, token1_symbol,token1_amount, token2_symbol, token2_amount, terms").Order("seq_num asc").Find(&StagePayments)

		paymentResponse.StageTerms = StagePayments
		var sql = fmt.Sprintf("SELECT SUM(token_amount) FROM bounty_deposit WHERE bounty_id = %d and status != 2 and comer_id != %d", bountyID, comerID)
		db.Raw(sql).Scan(&paymentResponse.ApplicantsTotalDeposit)
		return &paymentResponse, nil
	}
	if paymentResponse.PaymentMode == PaymentModePeriod {
		var periodMode []PeriodMode
		db.Table("bounty_payment_terms").Select("seq_num, status, token1_symbol,token1_amount, token2_symbol, token2_amount").Order("seq_num asc").Find(&periodMode)
		var terms string
		db.Table("bounty_payment_terms").Select("terms").Order("seq_num asc").Find(&terms)

		paymentResponse.PeriodTerms.PeriodModes = periodMode
		paymentResponse.PeriodTerms.Terms = terms
		var sql = fmt.Sprintf("SELECT SUM(token_amount) FROM bounty_deposit WHERE bounty_id = %d and status != 2 and comer_id != %d", bountyID, comerID)
		db.Raw(sql).Scan(&paymentResponse.ApplicantsTotalDeposit)
		return &paymentResponse, nil
	}
	return nil, nil
}

func UpdateBountyStatusByID(db *gorm.DB, bountyID uint64, isDeleted int) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("is_deleted", isDeleted).Error
}

func UpdatePaidStatusByBountyID(db *gorm.DB, bountyID uint64, request *PaidStatusRequest) error {
	return db.Table("bounty_payment_terms").Where("bounty_id = ? and seq_num = ?", bountyID, request.SeqNum).Update("status", request.Status).Error
}

func CreateApplicants(db *gorm.DB, request *BountyApplicant) error {
	return db.Create(&request).Error
}

func GetActivitiesByBountyID(db *gorm.DB, bountyID uint64) ([]*ActivitiesResponse, error) {
	var comerIDs []uint64
	var comerInfo ComerInfo
	var activitiesResponse ActivitiesResponse
	var activitiesTotal []*ActivitiesResponse
	err := db.Table("post_update").Select("comer_id").Where("bounty_id = ?", bountyID).Find(&comerIDs).Error
	if err != nil {
		return nil, err
	}
	for _, comerID := range comerIDs {
		db.Table("comer_profile").Select("name, avatar").Where("comer_id = ?", comerID).Find(&comerInfo)
		activitiesResponse.ComerInfo = comerInfo
		db.Table("post_update").Select("content, created_at, source_type").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Find(&activitiesResponse.ActivitiesContent)
		activitiesResponse.ComerID = comerID
		activitiesTotal = append(activitiesTotal, &activitiesResponse)
	}
	return activitiesTotal, nil
}

func GetApplicants(db *gorm.DB, bountyID uint64) (*BountyApplicantsResponse, error) {
	var comerID uint64
	err := db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&comerID).Error
	if err != nil {
		return nil, err
	}
	var applicantComerIDs []uint64
	err = db.Table("bounty_deposit").Select("comer_id").Where("bounty_id = ? and comer_id != ?", bountyID, comerID).Find(&applicantComerIDs).Error
	if err != nil {
		return nil, err
	}
	var comerInfo ComerInfo
	var bountyApplicant BountyApplicant
	var applicant Applicant
	var applicantResponse BountyApplicantsResponse
	for _, applicantComerID := range applicantComerIDs {
		db.Table("comer_profile").Select("name, avatar").Where("comer_id = ?", applicantComerID).Find(&comerInfo)
		db.Table("bounty_applicant").Select("description, apply_at").Where("comer_id = ? and bounty_id = ?", applicantComerID, bountyID).Find(&bountyApplicant)
		applicant.Name = comerInfo.ComerName
		applicant.Image = comerInfo.ComerImage
		applicant.Description = bountyApplicant.Description
		applicant.ApplyAt = bountyApplicant.ApplyAt
		applicant.ComerID = applicantComerID
		applicantResponse.Applicants = append(applicantResponse.Applicants, applicant)
		return &applicantResponse, nil
	}
	return nil, nil
}

func GetFounderByBountyID(db *gorm.DB, bountyID uint64) (*FounderResponse, error) {
	var comerID uint64
	var comerInfo account.ComerProfile
	var tagIds []uint64
	var skillNames []string
	var skillName string
	var founderInfo FounderResponse
	db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&comerID)
	db.Table("comer_profile").Select("name, avatar, time_zone").Where("comer_id = ?", comerID).Find(&comerInfo)
	db.Table("tag_target_rel").Select("tag_id").Where("target_id = ?", bountyID).Find(&comerInfo)
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ?", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	founderInfo.ComerID = comerID
	founderInfo.Name = comerInfo.Name
	founderInfo.Image = comerInfo.Avatar
	founderInfo.TimeZone = comerInfo.TimeZone
	founderInfo.ApplicantsSkills = skillNames
	founderInfo.Location = comerInfo.Location
	founderInfo.Email = comerInfo.Email
	return &founderInfo, nil
}

func GetApprovedApplicantByBountyID(db *gorm.DB, bountyID uint64) (*ApprovedResponse, error) {
	var comerID uint64
	err := db.Table("bounty_applicant").Select("comer_id").Where("bounty_id = ? and status = 2", bountyID).Find(&comerID).Error
	if err != nil {
		return nil, err
	}
	var comerInfo account.ComerProfile
	var tagIds []uint64
	var skillNames []string
	var skillName string
	var approvedInfo ApprovedResponse
	db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&comerID)
	db.Table("comer_profile").Select("name, avatar, time_zone").Where("comer_id = ?", comerID).Find(&comerInfo)
	db.Table("tag_target_rel").Select("tag_id").Where("target_id = ?", bountyID).Find(&comerInfo)
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ?", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	approvedInfo.ComerID = comerID
	approvedInfo.Name = comerInfo.Name
	approvedInfo.Image = comerInfo.Avatar
	approvedInfo.ApplicantsSkills = skillNames
	return &approvedInfo, nil
}

func GetDepositRecordsByBountyID(db *gorm.DB, bountyID uint64) (*DepositRecordsResponse, error) {
	var comerIDs []uint64
	var depositRecorids DepositRecordsResponse
	var depositRecorid DepositRecord
	err := db.Table("bounty_deposit").Select("comer_id").Where("bounty_id = ? ", bountyID).Find(&comerIDs).Error
	if err != nil {
		return nil, err
	}
	for _, comerID := range comerIDs {
		db.Table("bounty_deposit").Select("token_amount, access, created_at").Where("comer_id = ?", comerID).Find(&depositRecorid)
		depositRecorid.ComerID = comerID
		depositRecorids.DepositRecords = append(depositRecorids.DepositRecords, depositRecorid)
	}
	return &depositRecorids, nil
}

func UpdateApplicantApprovedStatus(db *gorm.DB, bountyID, comerID uint64, status int) (err error) {
	return db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Update("status", status).Error
}

func GetStartupByBountyID(db *gorm.DB, bountyID uint64) (*StartupListResponse, error) {
	var startupID uint64
	err := db.Table("bounty").Select("startup_id").Where("id = ?", bountyID).Find(&startupID).Error
	if err != nil {
		return nil, err
	}
	var startupListResponse StartupListResponse
	err = db.Table("startup").Select("name, mode, logo, chain_id, tx_hash, contract_audit, website, discord, twitter, telegram, docs, mission").Where("id = ?", startupID).Find(&startupListResponse).Error
	if err != nil {
		return nil, err
	}
	var tagIDs []uint64
	err = db.Table("tag_target_rel").Select("tag_id").Where("target_id = ?", startupID).Find(&tagIDs).Error
	if err != nil {
		return nil, err
	}
	var tagName string
	var tagNames []string
	for _, tagID := range tagIDs {
		err = db.Table("tag").Select("name").Where("id = ?", tagID).Find(&tagName).Error
		tagNames = append(tagNames, tagName)
	}
	startupListResponse.Tag = tagNames
	return &startupListResponse, nil
}
