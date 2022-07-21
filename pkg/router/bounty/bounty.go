/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/25 22:54
 */

package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/model/bounty"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/bounty"
	"ceres/pkg/utility/jwt"
	"fmt"
	"github.com/qiniu/x/log"
	"strconv"
)

// CreateBounty create bounty
func CreateBounty(ctx *router.Context) {
	request := new(bounty.BountyRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	fmt.Println(request, "router/bounty.go line:26") // 注释

	if err := service.CreateComerBounty(request); err != nil {
		ctx.HandleError(err)
		return
	}
	response := "create bounty successful!"

	ctx.OK(response)
}

func parsePagination(ctx *router.Context, pagination *model.Pagination, defaultLimit int) (err error) {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	if page == 0 {
		page = 1
	}
	pagination.Page = page
	_sort := ctx.Query("sort")
	switch _sort {
	case "Created:Recent":
		_sort = "created_at desc"
	case "Created:Oldest":
		_sort = "created_at asc"
	case "Value:Highest":
		_sort = "total_reward_token desc"
	case "Value:Lowest":
		_sort = "total_reward_token asc"
	case "Deposit:Highest":
		_sort = "founder_deposit desc"
	case "Deposit:Lowest":
		_sort = "founder_deposit asc"
	default:
		_sort = "created_at desc"
	}

	pagination.Sort = _sort
	pagination.Limit = defaultLimit
	log.Infof("pagination param is : %v\n", pagination)
	return nil
}

// GetPublicBountyList bounty list displayed in bounty tab
func GetPublicBountyList(ctx *router.Context) {
	var request model.Pagination
	if err := parsePagination(ctx, &request, 10); err != nil {
		ctx.HandleError(err)
		return
	}

	if response, err := service.QueryAllOnChainBounties(request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetBountyListByStartup get bounty list belongs to startup
func GetBountyListByStartup(ctx *router.Context) {
	var request model.Pagination
	if err := parsePagination(ctx, &request, 3); err != nil {
		ctx.HandleError(err)
		return
	}
	startupId, err := strconv.ParseUint(ctx.Param("startupId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if startupId == 0 {
		err := router.ErrBadRequest.WithMsg("Invalid startupId!")
		ctx.HandleError(err)
		return
	}

	if response, err := service.QueryBountiesByStartup(startupId, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetMyPostedBountyList get bounty list posted by me
func GetMyPostedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.Pagination
	if err := parsePagination(ctx, &request, 5); err != nil {
		ctx.HandleError(err)
		return
	}

	if response, err := service.QueryComerPostedBountyList(comerID, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

// GetMyParticipatedBountyList get bounty list
func GetMyParticipatedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.Pagination
	if err := parsePagination(ctx, &request, 8); err != nil {
		ctx.HandleError(err)
		return
	}

	if response, err := service.QueryComerParticipatedBountyList(comerID, request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(response)
	}
}

func GetBountyDetailByID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetBountyDetailByID(bountyID)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("get bounty detail fail")
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetPaymentByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	fmt.Println("router.go line:165:", bountyID) //注释
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetPaymentByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func UpdateBountyStatus(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	request := new(bounty.BountyCloseRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateBountyStatusByID(bountyID, request.IsDeleted)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("close bounty success")
}

func AddDeposit(ctx *router.Context) {
	request := new(bounty.AddDepositRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err := service.AddDeposit(request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("add deposit success")
}

func UpdatePaidStatusByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	request := new(bounty.PaidStatusRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdatePaidStatusByBountyID(bountyID, request)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK("update paid status success")
}

func CreateActivities(ctx *router.Context) {
	request := new(bounty.ActivitiesRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err := service.CreateActivities(request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("activities create success")
}

func CreateApplicants(ctx *router.Context) {
	request := new(bounty.ApplicantsDepositRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err := service.CreateApplicants(request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("create applicants success")
}

func GetActivitiesLists(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetActivitiesByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetAllApplicantsByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetAllApplicantsByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetFounderByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetFounderByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetApprovedApplicantByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetApprovedApplicantByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetDepositRecords(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetDepositRecords(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func UpdateFounderApprovedApplicant(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	request := new(bounty.ApprovedRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateApplicantApprovedStatus(bountyID, request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("approved success")
}

func UpdateFounderUnapprovedApplicant(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	request := new(bounty.ApprovedRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateApplicantApprovedStatus(bountyID, request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("approved success")
}

func GetStartupByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetStartupByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetBountyRoleByComerID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	fmt.Println(bountyID)
	header := ctx.Request.Header
	token := header.Get("x-comunion-authorization")
	fmt.Println(token)
	s, _ := jwt.Verify(token)
	fmt.Println(s)
}
