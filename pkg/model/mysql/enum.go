// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_200720
package mysql

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"ceres/pkg/invoker"
	"ceres/pkg/model/transport"
	
)

type Enum struct {
	
	Id int `gorm:"AUTO_INCREMENT;comment:'id'"` id
	
	GroupKey string `gorm:"not null;comment:'组名key，唯一标识符'"` 唯一标识符
	
	GroupTitle string `gorm:"not null;comment:'名称'"` 名称
	
	Key int `gorm:"not null;comment:'ant键'"` 值
	
	Title string `gorm:"not null;comment:'ant名称'"` 名称
	
	Ctime int64 `gorm:"not null;comment:'创建时间'"` 创建时间
	
	Utime int64 `gorm:"not null;comment:'更新时间'"` 更新时间
	
	Dtime int64 `gorm:"not null;comment:'删除时间'"` 删除时间
	
}

type Enums []*Enum

// TableName 设置表明
func (t Enum) TableName() string {
	return "enum"
}

// EnumCreate 创建一条记录
func EnumCreate(db *gorm.DB,data *Enum) (err error) {
	if err = db.Create(data).Error; err != nil {
		invoker.Logger.Error("create enum error", zap.Error(err))
		return
	}
	return
}

// EnumUpdate 根据主键更新一条记录
func EnumUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	
	if err = db.Model(Enum{}).Where(sql, binds...).Updates(ups).Error; err != nil {
		invoker.Logger.Error("enum update error", zap.Error(err))
		return
	}
	return
}

// EnumUpdateX Update的扩展方法，根据Cond更新一条或多条记录
func EnumUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Model(Enum{}).Where(sql, binds...).Updates(ups).Error; err != nil {
		invoker.Logger.Error("enum update error", zap.Error(err))
		return
	}
	return
}

// EnumDelete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func EnumDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	
			if err = db.Model(Enum{}).Where(sql, binds...).Delete(&Enum{}).Error; err != nil {
					invoker.Logger.Error("enum delete error", zap.Error(err))
					return
			}
	
	return
}

// EnumDeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func EnumDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)
  
    if err = db.Model(Enum{}).Where(sql, binds...).Delete(&Enum{}).Error; err != nil {
        invoker.Logger.Error("enum delete error", zap.Error(err))
        return
    }
  
	return
}

// EnumInfo 根据PRI查询单条记录
func EnumInfo(db *gorm.DB, paramId int) (resp Enum, err error) {
    var sql = "`id`= ?"
	var binds = []interface{}{paramId}

	if err = db.Model(Enum{}).Where(sql, binds...).First(&resp).Error; err != nil {
		invoker.Logger.Error("enum info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func EnumInfoX(db *gorm.DB, conds Conds) (resp Enum, err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Model(Enum{}).Where(sql, binds...).First(&resp).Error; err != nil {
		invoker.Logger.Error("enum info error", zap.Error(err))
		return
	}
	return
}

// EnumList 查询list，extra[0]为sorts
func EnumList(conds Conds, extra ...string) (resp []*Enum, err error) {
  
	sql, binds := BuildQuery(conds)
	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = invoker.Db.Model(Enum{}).Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		invoker.Logger.Error("enum info error", zap.Error(err))
		return
	}
	return
}

// EnumListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func EnumListMap(conds Conds) (resp map[int]*Enum, err error) {
  
	sql, binds := BuildQuery(conds)
	mysqlSlice := make([]*Enum, 0)
	resp = make(map[int]*Enum, 0)
	if err = invoker.Db.Model(Enum{}).Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		invoker.Logger.Error("enum info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// EnumListPage 根据分页条件查询list
func EnumListPage( conds Conds, reqList *transport.ReqPage) (total int, respList Enums) {
  respList = make(Enums,0)
  
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := invoker.Db.Model(Enum{}).Where(sql, binds...)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}

