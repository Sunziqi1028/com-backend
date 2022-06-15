package account

import (
	model "ceres/pkg/model/account"
	"testing"
)

func TestGetComerProfile(t *testing.T) {
	type args struct {
		comerID  uint64
		response *model.ComerProfileResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetComerProfile(tt.args.comerID, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("GetComerProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
