package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"github.com/qiniu/x/log"
)

func FollowComer(comerID, targetComerID uint64) (err error) {
	return model.CreateComerFollowRel(mysql.DB, comerID, targetComerID)
}

func UnfollowComer(comerID, targetComerID uint64) (err error) {
	followRel := model.FollowRelation{
		ComerID:       comerID,
		TargetComerID: targetComerID,
	}
	if err = model.DeleteComerFollowRel(mysql.DB, &followRel); err != nil {
		log.Warn(err)
	}
	return
}

func FollowedByComer(comerID, targetComerID uint64) (isFollowed bool, err error) {
	isFollowed, err = model.ComerFollowIsExist(mysql.DB, comerID, targetComerID)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}
