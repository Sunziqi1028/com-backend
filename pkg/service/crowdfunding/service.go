package crowdfunding

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/crowdfunding"
	"ceres/pkg/model/startup"
	model "ceres/pkg/model/transaction"
	"ceres/pkg/router"
	serviceTransaction "ceres/pkg/service/transaction"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
	"time"
)

func CreateCrowdfunding(request crowdfunding.CreateCrowdfundingRequest) error {
	var st startup.Startup
	if err := startup.GetStartup(mysql.DB, request.StartupId, &st); err != nil {
		return err
	}
	if st.ComerID != request.ComerId {
		return router.ErrBadRequest.WithMsg("Current Comer is not founder of startup")
	}
	crowdfundingList, er := crowdfunding.SelectOnGoingByStartupId(mysql.DB, request.StartupId)
	if er != nil {
		return er
	}
	// double check
	if len(crowdfundingList) > 0 {
		return router.ErrBadRequest.WithMsg("Startup has not ended crowdfunding")
	}
	funding := crowdfunding.Crowdfunding{
		ChainInfo:   request.ChainInfo,
		SellInfo:    request.SellInfo,
		BuyInfo:     request.BuyInfo,
		StartupId:   request.StartupId,
		ComerId:     request.ComerId,
		RaiseGoal:   request.RaiseGoal,
		TeamWallet:  request.TeamWallet,
		SwapPercent: request.SwapPercent,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		Poster:      request.Poster,
		Youtube:     request.Youtube,
		Detail:      request.Detail,
		Description: request.Description,
		Status:      0,
	}
	funding.RaiseBalance = funding.RaiseGoal
	funding.SellTokenBalance = funding.SellTokenDeposit
	if err := mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		log.Infof("#### TO_BE_CREATED_CROWDFUNDING:: %s", funding.Json())
		if er = crowdfunding.CreateCrowdfunding(mysql.DB, &funding); er != nil {
			return
		}
		x := int(crowdfunding.Pending)
		transaction := &model.Transaction{
			ChainID:    funding.ChainId,
			TxHash:     funding.TxHash,
			TimeStamp:  time.Now(),
			Status:     x,
			SourceType: serviceTransaction.CrowdfundingContractCreated,
			RetryTimes: 1,
			SourceID:   int64(funding.ID),
		}
		if er = model.CreateTransaction(tx, transaction); er != nil {
			return
		}
		return nil
	}); err != nil {
		return err
	}
	// query contract
	address, status := serviceTransaction.GetContractAddress(request.ChainId, request.TxHash)
	log.Infof("#### CONTRACT_ADDRESS_AND_ON_CHAIN_STATUS_OF_CREATED_CROWDFUNDING:: %s, %d\n", address, status)
	// update transaction status by corresponding source id
	if err := model.UpdateTransactionStatus(mysql.DB, funding.ID, status); err != nil {
		return err
	}
	// UpComing -> Live 怎么变？
	var fs crowdfunding.CrowdfundingStatus
	switch status {
	case 0:
		fs = crowdfunding.Upcoming
	case 1:
		fs = crowdfunding.Live
	case 2:
		fs = crowdfunding.OnChainFailure
	}
	if err := crowdfunding.UpdateCrowdfundingContractAddressAndStatus(mysql.DB, funding.ID, address, fs); err != nil {
		return err
	}
	return nil
}

func SelectNonFundingStartups(comerId uint64) (startups []crowdfunding.CrowdfundableStartup, err error) {
	return crowdfunding.SelectStartupsWithNonCrowdfundingOnGoing(mysql.DB, comerId)
}
