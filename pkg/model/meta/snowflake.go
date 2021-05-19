package meta

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Snowflake struct {
	ID       uint64    `gorm:"id"`
	IP       string    `gorm:"ip"`
	CreateAt time.Time `gorm:"create_at"`
}

func (Snowflake) TableName() string {
	return "meta_snowflake_ip_tbl"
}

/// GetMachineID
/// use the database auto increament to generate the machine ID
func GetMachineID(db *gorm.DB, IP string) (id uint16) {
	model := &Snowflake{
		IP: IP,
	}
	db.Where("ip = ?", IP).Save(model)
	id = uint16(model.ID)
	return
}
