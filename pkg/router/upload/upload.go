package upload

import (
	minio "ceres/pkg/initialization/minio"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	utility "ceres/pkg/utility/minio"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"

	"github.com/gotomicro/ego/core/elog"
)

// GetPresignedURLForUpload  to generate the presigned url with the file url
// and the frontend using this URL to upload
func GetPresignedURLForUpload(ctx *router.Context) {
	name := ctx.Query("file_name")
	if name == "" {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"name is missing",
		)
	}
	uin := ctx.Keys[middleware.ComerUinContextKey]
	if uin == nil {
		//FIXME: should change to the HTTP status to redirect the FE re login
		ctx.ERROR(
			router.ErrForbidden,
			"should login again",
		)
	}
	suffix := filepath.Ext(name)
	prefix := name[0 : len(name)-len(suffix)]
	res := md5.Sum([]byte(prefix))
	prefix = hex.EncodeToString(res[:])
	signed, err := utility.PreSignUpload(minio.Client, prefix+"."+suffix)
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
