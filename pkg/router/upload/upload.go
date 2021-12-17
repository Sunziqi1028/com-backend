package upload

import (
	minio "ceres/pkg/initialization/minio"
	"ceres/pkg/router"
	utility "ceres/pkg/utility/minio"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"

	"github.com/qiniu/x/log"
)

// GetPresignedURLForUpload  to generate the presigned url with the file url
// and the frontend using this URL to upload
func GetPresignedURLForUpload(ctx *router.Context) {
	name := ctx.Query("file_name")
	if name == "" {
		err := router.ErrBadRequest.WithMsg("name is missing")
		ctx.HandleError(err)
		return
	}
	suffix := filepath.Ext(name)
	prefix := name[0 : len(name)-len(suffix)]
	res := md5.Sum([]byte(prefix))
	prefix = hex.EncodeToString(res[:])
	signed, err := utility.PreSignUpload(minio.Client, prefix+"."+suffix)
	if err != nil {
		log.Warn("signed the MINIO url failed %v", err)
		ctx.HandleError(err)
		return
	}

	ctx.OK(signed)
}
