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
	"ceres/pkg/model/tag"
	"ceres/pkg/service/postupdate"
	"ceres/pkg/service/transaction"
	"ceres/pkg/utility/tool"
	"errors"
	"fmt"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
	"time"
)

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
)

// CreateComerBounty create bounty
func CreateComerBounty(request *model.BountyRequest) error {

	err := mysql.DB.Transaction(func(tx *gorm.DB) (ere error) {
		paymentMode, totalRewardToken := handlePayDetail(request.PayDetail)

		bountyID, err := createBounty(tx, paymentMode, totalRewardToken, request)
		if err != nil {
			log.Warn(err)
			return
		}
		if bountyID == 0 {
			return errors.New(fmt.Sprintf("create bounty err: %d", bountyID))
		}

		getContract(request.ChainID, request.TxHash, bountyID)

		err = transaction.CreateTransaction(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return
		}

		err = postupdate.CreatePostUpdate(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return
		}

		err = createDeposit(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return
		}

		errorsLog := createPaymentTerms(tx, bountyID, request)
		if len(errorsLog) > 0 {
			return errors.New(fmt.Sprintf("create payment_terms err:%v", errorsLog))
		}

		err = creatPaymentPeriod(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return
		}

		errorsLog = createContact(tx, bountyID, request)
		if len(errorsLog) > 0 {
			return errors.New(fmt.Sprintf("create contact address err:%v", errorsLog))
		}

		err = createApplicantsSkills(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return
		}

		return nil
	})

	return err
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
		Status:             BountyStatusReadyToWork,
		TotalRewardToken:   totalRewardToken,
	}

	bountyID, err := model.CreateBounty(tx, bounty)
	if err != nil || bountyID == 0 {
		return 0, errors.New(fmt.Sprintf("created bounty err:%s", err))
	}
	return bountyID, nil
}

func createDeposit(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	deposit := &model.BountyDeposit{
		ChainID:     request.ChainID,
		TxHash:      request.TxHash,
		Status:      transaction.Pending,
		BountyID:    bountyID,
		ComerID:     request.ComerID,
		Access:      AccessIn,
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
		contactModel := &model.BountyContact{
			BountyID:       bountyID,
			ContactType:    contact.ContactType,
			ContactAddress: contact.ContactAddress,
		}
		err := model.CreateContact(tx, contactModel)
		if err != nil {
			errorLog = append(errorLog, fmt.Sprintf("create contactAddress:%s err:%v", contact.ContactAddress, err))
			continue
		}
	}
	return errorLog
}

func createPaymentTerms(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) []string {
	paymentMode, _ := handlePayDetail(request.PayDetail)
	var errorLog []string
	if paymentMode == PaymentModeStage {
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
				Status:       BountyPaymentTermsStatusUnpaid,
			}
			err := model.CreatePaymentTerms(tx, paymentTerms)
			if err != nil {
				errorLog = append(errorLog, fmt.Sprintf("create stage %v err:%v", stage, err))
				continue
			}
		}
	} else {
		paymentTerms := &model.BountyPaymentTerms{
			BountyID:     bountyID,
			PaymentMode:  paymentMode,
			Token1Symbol: request.Period.Token1Symbol,
			Token1Amount: request.Period.Token1Amount,
			Token2Symbol: request.Period.Token2Symbol,
			Token2Amount: request.Period.Token2Amount,
			Terms:        request.Period.Target,
			SeqNum:       BountyPaymentTermsPeriodSeqNum,
			Status:       BountyPaymentTermsStatusUnpaid,
		}
		err := model.CreatePaymentTerms(tx, paymentTerms)
		if err != nil {
			errorLog = append(errorLog, fmt.Sprintf("create period err:%v", err))
			return errorLog
		}
	}

	return errorLog
}

func creatPaymentPeriod(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	paymentMode, _ := handlePayDetail(request.PayDetail)
	if paymentMode == PaymentModePeriod {
		periodAmount := int64(request.Period.Token1Amount + request.Period.Token2Amount)
		paymentPeriod := &model.BountyPaymentPeriod{
			BountyID:     bountyID,
			PeriodType:   request.Period.PeriodType,
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
	return nil
}

func createApplicantsSkills(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	for _, applicantsSkill := range request.ApplicantsSkills {
		tagID, err := model.GetAndUpdateTagID(tx, applicantsSkill)
		if err != nil {
			return err
		}
		tarTargetRel := &tag.TagTargetRel{
			TargetID: bountyID,
			TagID:    tagID,
			Target:   tag.Bounty,
		}
		err = model.CreateTagTargetRel(tx, tarTargetRel)
		if err != nil {
			return err
		}
	}
	return nil
}

func getContract(chainID uint64, txHash string, bountyID uint64) {
	var contractChan = make(chan *model.ContractInfoResponse, 1)
	go func() {
		contractAddress, status := transaction.GetContractAddress(chainID, txHash)
		contractInfo := &model.ContractInfoResponse{
			ContractAddress: contractAddress,
			Status:          status,
		}
		select {
		case contractChan <- contractInfo:
			for contract := range contractChan {
				transaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, bountyID, contract.Status, contract.ContractAddress)
			}
		case <-time.After(5 * time.Second):
			fmt.Println("get contract address time over!")
		}
	}()
	return
}

func handlePayDetail(request model.PayDetail) (paymentMode, totalRewardToken int) {
	if len(request.Stages) > 0 {
		paymentMode = PaymentModeStage
		for _, stage := range request.Stages {
			totalRewardToken = stage.Token1Amount + stage.Token2Amount
		}
		return paymentMode, totalRewardToken
	} else {
		paymentMode = PaymentModePeriod
		totalRewardToken = request.Period.Token1Amount + request.Period.Token2Amount
		return paymentMode, totalRewardToken
	}
}
