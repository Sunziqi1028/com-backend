package account

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Comer the comer model of comunion inner account
type Comer struct {
	ID        uint64    `gorm:"column:id"`
	Address   string    `gorm:"column:address"`
	Nick      string    `gorm:"column:nick"`
	City      string    `gorm:"column:city"`
	Avatar    string    `gorm:"column:avatar"`
	Blog      string    `gorm:"column:blog"`
	Intro     string    `gorm:"column:intro"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName Comer table name for gorm
func (Comer) TableName() string {
	return "comer"
}

type AccountType int

const (
	GithubOauth   AccountType = 1
	TwitterOauth  AccountType = 2
	FacebookOauth AccountType = 3
	LikedinOauth  AccountType = 4
	GoogleOauth   AccountType = 5
)

type AccountCategory int

const (
	OauthAccount AccountCategory = 1
)

// Account the account model of outer account
type Account struct {
	ID        uint64          `gorm:"column:id"`
	ComerID   uint64          `gorm:"column:comer_id"`
	OIN       string          `gorm:"column:oin"`
	IsPrimary bool            `gorm:"column:is_primary"`
	Nick      string          `gorm:"column:nick"`
	Avatar    string          `gorm:"column:avatar"`
	Category  AccountCategory `gorm:"column:category"`
	Type      AccountType     `gorm:"column:type"`
	IsLinked  bool            `gorm:"column:linked"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
}

// TableName the Account table name for gorm
func (Account) TableName() string {
	return "account"
}

// Profile the comer profile model
type Profile struct {
	ID          uint64    `gorm:"column:id"`
	UIN         uint64    `gorm:"column:uin"`
	Remark      string    `gorm:"column:remark"`
	Identifier  uint64    `gorm:"column:identifier"`
	Name        string    `gorm:"column:name"`
	About       string    `gorm:"column:about"`
	Description string    `gorm:"column:description"`
	Email       string    `gorm:"column:email"`
	Skills      string    `gorm:"column:skills"`
	Version     int       `gorm:"column:version"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UpdateAt    time.Time `gorm:"column:update_at"`
}

// TableName the Profile table name for gorm
func (Profile) TableName() string {
	return "comer_profile"
}

// ProfileSkillTag profile skill tag model
type ProfileSkillTag struct {
	ID       uint64    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Vaild    bool      `gorm:"column:vaild"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

// TableName the ProfileSkillTag table name for gorm
func (ProfileSkillTag) TableName() string {
	return "comer_profile_skill_tag_tbl"
}

// GetComerByAddress  get comer entity by comer's address
func GetComerByAddress(db *gorm.DB, address string) (comer Comer, err error) {
	if err = db.Where("address = ?", address).Find(&comer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}

// GetComerByID  get comer entity by comer's ID
func GetComerByID(db *gorm.DB, comerID uint64) (comer Comer, err error) {
	if err = db.Where("id = ?", comerID).Find(&comer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}

// CreateComer create a comer
func CreateComer(db *gorm.DB, comer *Comer) (err error) {
	r := db.Save(comer)
	return r.Error
}

//UpdateComerAddress update the comer address
func UpdateComerAddress(db *gorm.DB, comerID uint64, address string) (err error) {
	return db.Model(&Comer{ID: comerID}).Update("address", address).Error
}

// CreateComerWithAccount  using the outer acccount to create a comer
func CreateComerWithAccount(db *gorm.DB, comer *Comer, account *Account) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		r := tx.Save(comer)
		e := r.Error
		if e != nil {
			return e
		}
		r = tx.Save(account)
		e = r.Error
		if e != nil {
			return e
		}
		return nil
	})

	return
}

// DeleteComer  delete the comer
func DeleteComer(db *gorm.DB, comer *Comer) {
	db.Delete(comer)
}

// UpdateComer update the comer
func UpdateComer(db *gorm.DB, comer *Comer) (err error) {
	r := db.Save(comer)
	err = r.Error

	return
}

// GetAccountByOIN get the outer account by OIN
func GetAccountByOIN(db *gorm.DB, oin string) (account Account, err error) {
	db = db.Where("oin = ?", oin).First(&account)
	err = db.Error

	return
}

// GetAccountByIdentifier get account by identifier
func GetAccountByIdentifier(db *gorm.DB, identifier uint64) (account Account, err error) {
	db = db.Where("identifier = ?", identifier).First(&account)
	err = db.Error

	return
}

// LinkComerWithAccount  link a new account to an existed comer
func LinkComerWithAccount(db *gorm.DB, uin uint64, account *Account) (err error) {
	//if account.UIN != uin {
	//	err = errors.New("illegal comer UIN to link") // double check but this logic also implement in the router module
	//	return
	//}
	//r := db.Save(account)
	//err = r.Error

	return
}

// UnlinkComerAccount unlink one account of comer
func UnlinkComerAccount(db *gorm.DB, account *Account) (err error) {
	//account.IsLinked = false
	//account.UIN = 0
	//db = db.Save(account)
	//err = db.Error
	return
}

// ListAllAccountsOfComer  list all accounts of this comer with uin
func ListAllAccountsOfComer(db *gorm.DB, uin uint64) (list []Account, err error) {
	res := db.Where("uin = ?", uin).Find(&list)
	err = res.Error

	return
}

// GetComerByAccountUIN  get comer by account uin
func GetComerByAccountUIN(db *gorm.DB, uin uint64) (comer Comer, err error) {
	db = db.Where("uin = ?", uin).First(&comer)
	err = db.Error

	return
}

// GetComerByAccountOIN  get comer entity by the account oin
func GetComerByAccountOIN(db *gorm.DB, accountType AccountType, oin string) (account Account, err error) {
	if err = db.Where("oin = ? and type = ?", oin, accountType).Find(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}

// GetComerProfile by the uin
// FIXME: should change the function name
func GetComerProfile(db *gorm.DB, uin uint64) (profile Profile, err error) {
	db = db.Where("uin = ?", uin).First(&profile)
	err = db.Error

	return
}
