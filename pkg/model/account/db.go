package account

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

/// Account Database models and operations

// constraints of the category and account
const (
	EthAccount   = 1
	OauthAccount = 2

	GithubOauth   = 1
	MetamaskEth   = 2
	TwitterOauth  = 3
	FacbookOauth  = 4
	LinkedInOauth = 5
	ImtokenEth    = 6
)

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

/// profile skill tag
type ProfileSkillTag struct {
	ID       uint64    `gorm:"id"`
	Name     string    `gorm:"name"`
	Vaild    bool      `gorm:"vaild"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

func (ProfileSkillTag) TableName() string {
	return "comer_profile_skill_tag_tbl"
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

/// GetComerByAccoutOIN
/// get comer entity by the account oin
func GetComerByAccoutOIN(db *gorm.DB, oin string) (comer Comer, err error) {
	account := &Account{}
	db = db.Where("oin = ?", oin).Find(account)
	err = db.Error
	if err != nil {
		return
	}
	uin := account.UIN
	db = db.Where("uin = ?", uin).Find(&comer)
	err = db.Error
	if err != nil {
		return
	}
	return
}

/// GetComerProfile by the uin
func GetComerProfile(db *gorm.DB, uin uint64) (profile Profile, err error) {

	return
}

/// GetSkillList by the ids
func GetSkillList(db *gorm.DB, ids []uint64) (skills []ProfileSkillTag, err error) {

	return
}
