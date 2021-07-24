package pkg

import (
	"ceres/pkg/utility/auth"
	"testing"
)

func TestAuthGithub(t *testing.T) {
	github := auth.NewGithubClient("a7baa7cf66921570b604", "2c63d784110ab1a461b7baf81fd38a1241260b28")
	_, err := github.GetUserProfile()
	if err != nil {
		t.Error(err)
	}
}
