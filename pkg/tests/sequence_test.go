package pkg

import (
	"ceres/pkg/utility/net"
	"ceres/pkg/utility/sequence"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"ceres/pkg/model/account"
	"time"

	"github.com/google/uuid"
)

func TestSeq(t *testing.T) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	machineID %= 32
	if err != nil {
		return
	}
	flake := sequence.NewSnowflake(1525705533000, uint64(machineID))

	for i := 0; i <= 4097; i++ {
		fmt.Printf("%d\n", flake.Next())
		fmt.Println(flake)
	}
}

func TestLogin(t *testing.T) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	machineID %= 32
	if err != nil {
		return
	}
	// Create snowflake sequences
	AccountSequnece := sequence.NewSnowflake(1, uint64(machineID))
	fmt.Println(AccountSequnece)
	// try to find comer
	// comer, err := account.GetComerByAccountOIN(mysql.DB, oauth.GetUserID())
	// if err != nil {
	// 	elog.Errorf("Comunion Oauth login faild, because of %v", err)
	// 	return
	// }

	comer := account.Comer{}
	comer.ID = 0

	if comer.ID == 0 {
		// create comer with account
		now := time.Now()
		comer.UIN = AccountSequnece.Next()
		fmt.Println(comer.UIN)
		comer.ComerID = strings.Replace(uuid.Must(uuid.NewV4(), nil).String(), "-", "", -1)
		comer.CreateAt = now
		comer.UpdateAt = now
		if comer.Avatar == "" {
			comer.Avatar = comer.ComerID
		}
	}

	// sign with jwt using the comer UIN

	// token := jwt.Sign(comer.UIN)
	// fmt.Println(token)

	response := &account.ComerLoginResponse{
		ComerID: comer.ComerID,
		Address: "127.0.0.1",
		Nick:    "Adennan",
		Avatar:  comer.Avatar,
		Token:   "token",
		UIN:     comer.UIN,
	}
	fmt.Println(response)
}
