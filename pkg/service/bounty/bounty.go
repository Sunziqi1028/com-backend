/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/29 10:04
 */

package bounty

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/bounty"
	"ceres/pkg/model/startup"
	service "ceres/pkg/service/startup"
	"ceres/pkg/utility/tool"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GetStartupsByComerID(comerID uint64) (*model.GetStartupsResponse, error) {
	var startups *model.GetStartupsResponse
	startupsResponse, err := model.GetComerStartups(mysql.DB, comerID, startups)
	if err != nil {
		return nil, err
	}
	return startupsResponse, nil
}

// CreateComerBounty create bounty
func CreateComerBounty(request *model.BountyRequest) error {
	var startupInfo startup.GetStartupResponse
	err := service.GetStartup(request.StartupID, &startupInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("get deposit total supply err:%s", err))
	}
	if int64(request.Deposit.TokenAmount) > startupInfo.TotalSupply {
		return errors.New("check the deposit amount")
	}

	tx := mysql.DB.Begin() // begin Transaction

	paymentMode, totalRewardToken := handlePayDetail(request.PayDetail)

	bountyID, err := createBounty(tx, paymentMode, totalRewardToken, request)
	if err != nil {
		return err
	}
	if bountyID == 0 {
		return errors.New("")
	}
	err = createDeposit(tx, bountyID, request)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createTransaction(tx, bountyID, request)
	if err != nil {
		tx.Rollback()
		return err
	}

	errorsLog := createPaymentTerms(tx, bountyID, request)
	if len(errorsLog) > 0 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("create payment_terms err:%v", errorsLog))
	}

	err = creatPaymentPeriod(tx, bountyID, request)
	if err != nil {
		tx.Rollback()
		return err
	}

	errorsLog = createContact(tx, bountyID, request)
	if len(errorsLog) > 0 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("create contact address err:%v", errorsLog))
	}

	err = createPostUpdate(tx, bountyID, request)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func createBounty(tx *gorm.DB, paymentMode, totalRewardToken int, request *model.BountyRequest) (uint64, error) {
	bounty := &model.Bounty{
		StartupID:          request.StartupID,
		ComerID:            request.ComerID,
		ChainID:            request.ChainID,
		TxHash:             request.TxHash,
		Title:              request.Title,
		ApplyCutoffDate:    tool.ParseTimeString2Time(request.ExpiresIn),
		DiscussionLink:     request.BountyDetail.DiscussionLink,
		DepositTokenSymbol: request.Deposit.TokenSymbol,
		ApplicantDeposit:   request.ApplicantsDeposit,
		FounderDeposit:     request.Deposit.TokenAmount,
		Description:        request.Description,
		PaymentMode:        paymentMode,
		Status:             0,
		TotalRewardToken:   totalRewardToken,
	}

	bountyID, err := model.CreateBounty(tx, bounty)
	if err != nil || bountyID == 0 {
		return 0, errors.New(fmt.Sprintf("created bounty err:%s", err))
	}
	return bountyID, nil
}

func createTransaction(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	transaction := &model.Transaction{
		ChainID:    request.ChainID,
		TxHash:     request.TxHash,
		TimeStamp:  time.Now(),
		Status:     0,
		SourceType: 2,
		SourceID:   int64(bountyID),
	}
	if err := model.CreateTransaction(tx, transaction); err != nil {
		return err
	}
	return nil
}

func createDeposit(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	deposit := &model.BountyDeposit{
		ChainID:     request.ChainID,
		TxHash:      request.TxHash,
		Status:      0,
		BountyID:    bountyID,
		ComerID:     request.ComerID,
		Access:      2,
		TokenSymbol: request.Deposit.TokenSymbol,
		TokenAmount: request.Deposit.TokenAmount,
		TimeStamp:   time.Now(),
	}
	err := model.CreateDeposit(tx, deposit)
	if err != nil {
		return err
	}
	return nil
}

func createContact(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) []string {
	var errorLog []string
	for _, contact := range request.Contacts {
		var (
			contactType    int
			contactAddress string
		)
		if len(contact.Email) > 0 {
			contactType = 1
			contactAddress = contact.Email
		} else if len(contact.Discord) > 0 {
			contactType = 2
			contactAddress = contact.Discord
		} else {
			contactType = 3
			contactAddress = contact.Telegram
		}
		contactModel := &model.BountyContact{
			BountyID:       bountyID,
			ContactType:    contactType,
			ContactAddress: contactAddress,
		}
		err := model.CreateContact(tx, contactModel)
		if err != nil {
			errorLog = append(errorLog, fmt.Sprintf("create contactAddress:%s", contactAddress))
			continue
		}
	}
	return errorLog
}

func createPaymentTerms(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) []string {
	paymentMode, _ := handlePayDetail(request.PayDetail)
	var errorLog []string
	if paymentMode == 1 {
		for _, stage := range request.PayDetail.Stages {
			paymentTerms := &model.BountyPaymentTerms{
				BountyID:     bountyID,
				PaymentMode:  paymentMode,
				Token1Symbol: stage.Token1Symbol,
				Token1Amount: stage.Token1Amount,
				Token2Symbol: stage.Token2Symbol,
				Token2Amount: stage.Token2Amount,
				Terms:        stage.Terms,
				SeqNum:       stage.SeqNum,
				Status:       1,
			}
			err := model.CreatePaymentTerms(tx, paymentTerms)
			if err != nil {
				errorLog = append(errorLog, fmt.Sprintf("creat stage %v err", stage))
				continue
			}
		}
	}

	return errorLog
}

func creatPaymentPeriod(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	var periodType int
	switch request.Period.PeriodType {
	case "days":
		periodType = 1
	case "weeks":
		periodType = 2
	case "months":
		periodType = 3
	}
	periodAmount := int64(request.Period.Token1Amount + request.Period.Token2Amount)
	paymentPeriod := &model.BountyPaymentPeriod{
		BountyID:     bountyID,
		PeriodType:   periodType,
		PeriodAmount: periodAmount,
		HoursPerDay:  request.Period.HoursPerDay,
		Token1Symbol: request.Period.Token1Symbol,
		Token1Amount: request.Period.Token1Amount,
		Token2Symbol: request.Period.Token2Symbol,
		Token2Amount: request.Period.Token2Amount,
		Target:       request.Period.Target,
	}
	err := model.CreatePaymentPeriod(tx, paymentPeriod)
	if err != nil {
		return err
	}
	return nil
}

func createPostUpdate(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	postUpdate := &model.PostUpdate{
		SourceType: 1, //1 bounty
		SourceID:   bountyID,
		ComerID:    request.ComerID,
		Content:    request.Description,
		TimeStamp:  time.Now(),
	}
	err := model.CreatePostUpdate(tx, postUpdate)
	if err != nil {
		return err
	}
	return nil
}

func CreateApplicantsSkills(tx *gorm.DB, bountyID uint64, tagID uint64) {

}

func handlePayDetail(request model.PayDetail) (paymentMode, totalRewardToken int) {
	if len(request.Stages) > 0 {
		paymentMode = 1
		for _, stage := range request.Stages {
			totalRewardToken = stage.Token1Amount + stage.Token2Amount
		}
		return paymentMode, totalRewardToken
	} else {
		paymentMode = 2
		totalRewardToken = request.Period.Token1Amount + request.Period.Token2Amount
		return paymentMode, totalRewardToken
	}
	return 0, 0
}
