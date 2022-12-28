package enums

import "fmt"

// ArticleFieldSort 文章model所支持排序的字段
type ArticleFieldSort int

const (
	ArticleFieldSortUnknown ArticleFieldSort = iota
	// ArticleFieldSortId 文章id
	ArticleFieldSortId
	// ArticleFieldSortTitle 文章标题
	ArticleFieldSortTitle
)

// SortOrder 排序顺序
type SortOrder int

const (
	SortOrderUnknown SortOrder = iota
	// SortOrderReverse 倒序
	SortOrderReverse
	// SortOrderBackward 正序
	SortOrderBackward
)

// ToSortBson 根据排序的枚举值转化成mongo 排序语句
func ToSortBson(filed ArticleFieldSort, order SortOrder) string {
	if filed == ArticleFieldSortUnknown || order == SortOrderUnknown {
		return ""
	}
	var (
		filedStr string
		orderStr string
	)
	switch filed {
	case ArticleFieldSortId:
		filedStr = "_id"
	case ArticleFieldSortTitle:
		filedStr = "title"
	}
	switch order {
	case SortOrderReverse:
		orderStr = "-"
	}
	return fmt.Sprintf("%s%s", orderStr, filedStr)
}
