package config

import (
	C "ceres/pkg/config"

	"github.com/gotomicro/ego/core/econf"
)

// Init the config structure from the .toml
func Init() error {
	// Logger = elog.DefaultLogger

	C.Github = &C.GithubOauth{}
	C.Facebook = &C.FacebookOauth{}
	C.Minio = &C.MinioConfig{}
	C.Seq = &C.Sequence{}
	C.JWT = &C.JWTConfig{}

	err := econf.UnmarshalKey("ceres.oauth.github", C.Github)
	if err != nil {
		return err
	}

	err = econf.UnmarshalKey("ceres.oauth.facebook", C.Facebook)
	if err != nil {
		return err
	}

	err = econf.UnmarshalKey("ceres.minio", C.Minio)
	if err != nil {
		return err
	}

	err = econf.UnmarshalKey("ceres.sequence", C.Seq)
	if err != nil {
		return err
	}

	err = econf.UnmarshalKey("ceres.jwt", C.JWT)
	if err != nil {
		return err
	}

	return nil
}
