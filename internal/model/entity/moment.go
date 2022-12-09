// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Moment is the golang structure for table moment.
type Moment struct {
	Id          int64       `json:"id"          description:""`
	Content     string      `json:"content"     description:"动态内容"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	Likes       int         `json:"likes"       description:"点赞数量"`
	IsPublished int         `json:"isPublished" description:"是否公开"`
}