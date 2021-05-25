package upload

import (
	minio "ceres/pkg/initialization/minio"
	"ceres/pkg/router"
	utility "ceres/pkg/utility/minio"

	"github.com/gotomicro/ego/core/elog"
)

/// GetPresignedURLForUpload
/// to generate the presigned url with the file url
/// and the frontend using this URL to upload
func GetPresignedURLForUpload(ctx *router.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"url is missing",
		)
	}
	signed, err := utility.PreSignUpload(minio.Client, url)
	if err != nil {
		elog.Errorf("signed the MINIO url failed %v", err)
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(signed)
}
