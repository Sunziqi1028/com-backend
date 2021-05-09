// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_202936
// @Deprecated： dead code
package api

import (
	"ceres/pkg/invoker"
	"ceres/pkg/model/dto"
	"ceres/pkg/model/mysql"
	"ceres/pkg/model/transport"
	"ceres/pkg/router/core"

	"github.com/spf13/cast"
)

// EnumInfo 查询单条记录
func EnumInfo(c *core.Context) {
	id := cast.ToInt(c.Param("id"))
	if id == 0 {
		c.JSONE(1, "请求错误", nil)
		return
	}
	info, _ := mysql.EnumInfo(invoker.Db, id)
	c.JSONOK(info)
}

// EnumList 查询多条带分页记录
func EnumList(c *core.Context) {
	req := &transport.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误", err)
		return
	}

	query := mysql.Conds{}
	if v := c.Query("id"); v != "" {
		query["id"] = v
	}
	if v := c.Query("group_key"); v != "" {
		query["group_key"] = v
	}
	if v := c.Query("group_title"); v != "" {
		query["group_title"] = v
	}
	if v := c.Query("key"); v != "" {
		query["key"] = v
	}
	if v := c.Query("title"); v != "" {
		query["title"] = v
	}
	if v := c.Query("ctime"); v != "" {
		query["ctime"] = v
	}
	if v := c.Query("utime"); v != "" {
		query["utime"] = v
	}
	if v := c.Query("dtime"); v != "" {
		query["dtime"] = v
	}

	total, list := mysql.EnumListPage(query, req)
	c.JSONPage(list, core.Pagination{
		Current: req.Current, PageSize: req.PageSize, Total: total,
	})
}

// EnumCreate 创建记录
func EnumCreate(c *core.Context) {
	req := &dto.EnumCreate{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误: "+err.Error(), err)
		return
	}

	create := &mysql.Enum{

		Id: req.Id,

		GroupKey: req.GroupKey,

		GroupTitle: req.GroupTitle,

		Key: req.Key,

		Title: req.Title,

		Ctime: req.Ctime,

		Utime: req.Utime,

		Dtime: req.Dtime,
	}

	err := mysql.EnumCreate(invoker.Db, create)
	if err != nil {
		c.JSONE(1, "创建失败", err)
		return
	}
	c.JSONOK(req)
}

// EnumUpdate 更新指定记录
func EnumUpdate(c *core.Context) {
	req := &dto.EnumUpdate{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误: "+err.Error(), err)
		return
	}

	id := cast.ToInt(c.Param("id"))
	if id == 0 {
		c.JSONE(1, "请求错误", nil)
		return
	}

	err := mysql.EnumUpdate(invoker.Db, id, mysql.Ups{
		"id":          req.Id,
		"group_key":   req.GroupKey,
		"group_title": req.GroupTitle,
		"key":         req.Key,
		"title":       req.Title,
		"ctime":       req.Ctime,
		"utime":       req.Utime,
		"dtime":       req.Dtime,
	})
	if err != nil {
		c.JSONE(1, "更新失败", err)
		return
	}
	c.JSONOK()
}

// EnumDelete 删除指定记录
func EnumDelete(c *core.Context) {
	id := cast.ToInt(c.Param("id"))
	if id == 0 {
		c.JSONE(1, "请求id错误", nil)
		return
	}

	err := mysql.EnumDelete(invoker.Db, id)
	if err != nil {
		c.JSONE(1, "删除失败", err)
		return
	}
	c.JSONOK()
}
