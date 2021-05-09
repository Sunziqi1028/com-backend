package account

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

/// Account Database models and operations

/// Comer the comer model of comunion inner account
type Comer struct {
	ID       uint64    `gorm:"id"`
	UIN      uint64    `gorm:"uin"`
	Address  string    `gorm:"address"`
	ComerID  string    `gorm:"comer_id"`
	Nick     string    `gorm:"nick"`
	Avatar   string    `gorm:"avatar"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

func (Comer) TableName() string {
	return "comer_tbl"
}

/// Account the account model of outer account
type Account struct {
	ID       uint64    `gorm:"id"`
	UIN      uint64    `gorm:"uin"`
	OIN      string    `gorm:"oin"`
	IsMain   bool      `gorm:"main"`
	Nick     string    `gorm:"nick"`
	Avatar   string    `gorm:"avatar"`
	Category int       `gorm:"categor"`
	Type     int       `gorm:"type"`
	IsLinked bool      `gorm:"linked"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

func (Account) TableName() string {
	return "account_tbl"
}

/// Profile the comer profile model
type Profile struct {
	ID          uint64    `gorm:"id"`
	UIN         uint64    `gorm:"uin"`
	Remark      string    `gorm:"remark"`
	Identifier  uint64    `gorm:"identifier"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	Email       string    `gorm:"email"`
	Skills      string    `gorm:"skills"`
	Version     int       `gorm:"version"`
	CreateAt    time.Time `gorm:"create_at"`
	UpdateAt    time.Time `gorm:"update_at"`
}

func (Profile) TableName() string {
	return "comer_profile_tbl"
}

/// CreateComerWithAccount
/// using the outer acccount to create a comer
func CreateComerWithAccount(db *gorm.DB, comer *Comer, account *Account) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		r := tx.Create(comer)
		e := r.Error
		if e != nil {
			return e
		}
		r = tx.Create(account)
		e = r.Error
		if e != nil {
			return e
		}
		return nil
	})
	return
}

/// DeleteComer
/// delete the comer
func DeleteComer(db *gorm.DB, comer *Comer) {
	db.Delete(comer)
}

/// UpdateComer
/// update the comer
func UpdateComer(db *gorm.DB, comer *Comer) (err error) {
	r := db.Save(comer)
	err = r.Error
	return
}

/// LinkComerWithAccount
/// link a new account to an existed comer
func LinkComerWithAccount(db *gorm.DB, uin uint64, account *Account) (err error) {
	if account.UIN != uin {
		err = errors.New("illegal comer UIN to link") // double check but this logic also implement in the router module
		return
	}
	r := db.Save(account)
	err = r.Error
	return
}

/// UnlinkComerAndAccount
/// unlink the account with oin ref to the comer with uin
func UnlinkComerAndAccount(db *gorm.DB, uin uint64, account *Account) {
	account.IsLinked = false
	db.Where("uin = ?", uin).Save(account)
}

/// ListAllAccountsOfComer
/// list all accounts of this comer with uin
func ListAllAccountsOfComer(db *gorm.DB, uin uint64) (list []Account, err error) {
	res := db.Find(&list)
	err = res.Error
	return
}
