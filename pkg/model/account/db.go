package account

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Account Database models and operations

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

// Comer the comer model of comunion inner account
type Comer struct {
	ID       uint64    `gorm:"column:id"`
	UIN      uint64    `gorm:"column:uin"`
	Address  string    `gorm:"column:address"`
	ComerID  string    `gorm:"column:comer_id"`
	Nick     string    `gorm:"column:nick"`
	Avatar   string    `gorm:"column:avatar"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

// TableName Comer table name for gorm
func (Comer) TableName() string {
	return "comer_tbl"
}

// Account the account model of outer account
type Account struct {
	ID         uint64    `gorm:"column:id"`
	Identifier uint64    `gorm:"column:identifier"`
	UIN        uint64    `gorm:"column:uin"`
	OIN        string    `gorm:"column:oin"`
	IsMain     bool      `gorm:"column:main"`
	Nick       string    `gorm:"column:nick"`
	Avatar     string    `gorm:"column:avatar"`
	Category   int       `gorm:"column:category"`
	Type       int       `gorm:"column:type"`
	IsLinked   bool      `gorm:"column:linked"`
	CreateAt   time.Time `gorm:"column:create_at"`
	UpdateAt   time.Time `gorm:"column:update_at"`
}

// TableName the Account table name for gorm
func (Account) TableName() string {
	return "account_tbl"
}

// Profile the comer profile model
type Profile struct {
	ID          uint64    `gorm:"primary_key;column:id"`
	ComerID     uint64    `gorm:"column:comer_id"`
	Name        string    `gorm:"column:name"`
	Location    string    `gorm:"column:location"`
	Website     string    `gorm:"column:website"`
	Bio         string    `gorm:"column:bio"`
	IsDelete     int      `gorm:"column:is_delete"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UpdateAt    time.Time `gorm:"column:update_at"`
}

// TableName the Profile table name for gorm
func (Profile) TableName() string {
	return "comer_profile"
}

// Skill model
type Skill struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	Name     string    `gorm:"column:name"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

// TableName the Skill table name for gorm
func (Skill) TableName() string {
	return "comer_skill"
}

// ComerSkillRel model
type ComerSkillRel struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	ComerID  uint64    `gorm:"column:comer_id"`
	SkillID  uint64    `gorm:"column:skill_id"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

func (ComerSkillRel) TableName() string {
	return "comer_skill_rel"
}

// Social model
type Social struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	Type     string    `gorm:"column:type"`
	Account  string    `gorm:"column:account"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

// TableName the Social table name for gorm
func (Social) TableName() string {
	return "comer_social"
}

// ComerSocialRel model
type ComerSocialRel struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	ComerID  uint64    `gorm:"column:comer_id"`
	SocialID uint64    `gorm:"column:social_id"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

func (ComerSocialRel) TableName() string {
	return "comer_social_rel"
}

// Wallet model
type Wallet struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	Address  string    `gorm:"column:address"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

// TableName the Wallet table name for gorm
func (Wallet) TableName() string {
	return "comer_wallet"
}

// ComerWalletRel model
type ComerWalletRel struct {
	ID       uint64    `gorm:"primary_key;column:id"`
	ComerID  uint64    `gorm:"column:comer_id"`
	WalletID uint64    `gorm:"column:wallet_id"`
	IsDelete  int      `gorm:"column:is_delete"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

func (ComerWalletRel) TableName() string {
	return "comer_wallet_rel"
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
	if account.UIN != uin {
		err = errors.New("illegal comer UIN to link") // double check but this logic also implement in the router module
		return
	}
	r := db.Save(account)
	err = r.Error

	return
}

// UnlinkComerAccount unlink one account of comer
func UnlinkComerAccount(db *gorm.DB, account *Account) (err error) {
	account.IsLinked = false
	account.UIN = 0
	db = db.Save(account)
	err = db.Error
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
func GetComerByAccountOIN(db *gorm.DB, oin string) (comer Comer, err error) {
	account := &Account{}
	if err = db.Where("oin = ?", oin).Find(account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	if err = db.Where("uin = ?", account.UIN).Find(&comer).Error; err != nil {
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
	db = db.Where("comer_id = ?", uin).First(&profile)
	err = db.Error

	return
}


// GetComerProfileByIdentifier by identifier
func GetComerProfileByIdentifier(db *gorm.DB, identifier uint64) (profile Profile, err error) {
	db = db.Where("identifier = ?", identifier).First(&profile)
	err = db.Error

	return
}

// CreateComerProfile create a new comer profile
func CreateComerProfile(db *gorm.DB, profile *Profile) (err error) {
	db = db.Save(profile)
	err = db.Error

	return
}

// UpdateComerProfile update the comer profile
func UpdateComerProfile(db *gorm.DB, profile *Profile) (err error) {
	db = db.Save(profile)
	err = db.Error

	return
}

// GetSkillList by the ids
func GetSkillList(db *gorm.DB, ids []uint64) (skills []Skill, err error) {
	db = db.Where("id in ?", ids).Find(&skills)
	err = db.Error

	return
}

// GetSkillListByNames by the names
func GetSkillListByNames(db *gorm.DB, names []string) (skills []Skill, err error) {
	db = db.Where("name in ?", names).Find(&skills)
	err = db.Error

	return
}

// CreateSkill create a new
func CreateSkill(db *gorm.DB, skill *Skill) (err error) {
	db = db.Save(skill)
	err = db.Error
	return
}

func GetSkillRels(db *gorm.DB, comer_id uint64) (skillRels []ComerSkillRel, err error) {
	db = db.Where("comer_id = ?", comer_id).Find(&skillRels)
	err = db.Error

	return
}

// CreateSkillRel create a new
func CreateSkillRel(db *gorm.DB, skillRel *ComerSkillRel) (err error) {
	db = db.Save(skillRel)
	err = db.Error
	return
}

func FirstOrCreateSkill(db *gorm.DB, skill *Skill) (err error) {
	db = db.Where(skill).FirstOrCreate(&skill)
	err = db.Error
	return
}

// FirstOrCreateSkillRel create a new
func FirstOrCreateSkillRel(db *gorm.DB, skillRel *ComerSkillRel) (err error) {
	db = db.Where(skillRel).FirstOrCreate(&skillRel)
	err = db.Error
	return
}


func GetSocialList(db *gorm.DB, ids []uint64) (socials []Social, err error) {
	db = db.Where("id in ?", ids).Find(&socials)
	err = db.Error

	return
}

func GetSocialListByAccount(db *gorm.DB, accounts []string) (socials []Social, err error) {
	db = db.Where("account in ?", accounts).Find(&socials)
	err = db.Error

	return
}

// CreateSocial create a new
func CreateSocial(db *gorm.DB, social *Social) (err error) {
	db = db.Save(social)
	err = db.Error
	return
}

func GetSocialRels(db *gorm.DB, comer_id uint64) (socialRels []ComerSocialRel, err error) {
	db = db.Where("comer_id = ?", comer_id).Find(&socialRels)
	err = db.Error

	return
}

// CreateSocialRel create a new
func CreateSocialRel(db *gorm.DB, socialRel *ComerSocialRel) (err error) {
	db = db.Save(socialRel)
	err = db.Error
	return
}

func FirstOrCreateSocial(db *gorm.DB, social *Social) (err error) {
	db = db.Where(social).FirstOrCreate(&social)
	err = db.Error
	return
}

// FirstOrCreateSocialRel create a new
func FirstOrCreateSocialRel(db *gorm.DB, socialRel *ComerSocialRel) (err error) {
	db = db.Where(socialRel).FirstOrCreate(&socialRel)
	err = db.Error
	return
}


func GetWalletList(db *gorm.DB, ids []uint64) (wallets []Wallet, err error) {
	db = db.Where("id in ?", ids).Find(&wallets)
	err = db.Error

	return
}

func GetWalletListByAddress(db *gorm.DB, addesss []string) (wallets []Wallet, err error) {
	db = db.Where("address in ?", addesss).Find(&wallets)
	err = db.Error

	return
}

// CreateWallet create a new
func CreateWallet(db *gorm.DB, wallet *Wallet) (err error) {
	db = db.Save(wallet)
	err = db.Error
	return
}

func GetWalletRels(db *gorm.DB, comer_id uint64) (walletRels []ComerWalletRel, err error) {
	db = db.Where("comer_id = ?", comer_id).Find(&walletRels)
	err = db.Error

	return
}

// CreateWalletRel create a new
func CreateWalletRel(db *gorm.DB, walletRel *ComerWalletRel) (err error) {
	db = db.Save(walletRel)
	err = db.Error
	return
}

func FirstOrCreateWallet(db *gorm.DB, wallet *Wallet) (err error) {
	db = db.Where(wallet).FirstOrCreate(&wallet)
	err = db.Error
	return
}

// FirstOrCreateWalletRel create a new
func FirstOrCreateWalletRel(db *gorm.DB, walletRel *ComerWalletRel) (err error) {
	db = db.Where(walletRel).FirstOrCreate(&walletRel)
	err = db.Error
	return
}


