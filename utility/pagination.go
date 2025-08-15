package utility

import (
	"github.com/gogf/gf/v2/util/gpage"
)

// PaginationParams 分页参数
type PaginationParams struct {
	Page     int `json:"page"`     // 当前页码
	PageSize int `json:"pageSize"` // 每页大小
}

// PaginationInfo 分页信息
type PaginationInfo struct {
	CurrentPage int `json:"currentPage"` // 当前页码
	PageSize    int `json:"pageSize"`    // 每页大小
	Total       int `json:"total"`       // 总记录数
	TotalPage   int `json:"totalPage"`   // 总页数
}

// PaginationOptions 分页选项
type PaginationOptions struct {
	DefaultPage     int `json:"defaultPage"`     // 默认页码
	DefaultPageSize int `json:"defaultPageSize"` // 默认每页大小
	MaxPageSize     int `json:"maxPageSize"`     // 最大每页大小
}

// DefaultPaginationOptions 默认分页选项
var DefaultPaginationOptions = PaginationOptions{
	DefaultPage:     1,
	DefaultPageSize: 10,
	MaxPageSize:     100,
}

// ValidateAndFixPagination 验证并修复分页参数
func ValidateAndFixPagination(params *PaginationParams, options *PaginationOptions) *PaginationParams {
	if options == nil {
		options = &DefaultPaginationOptions
	}

	// 设置默认分页参数
	page := params.Page
	if page <= 0 {
		page = options.DefaultPage
	}

	pageSize := params.PageSize
	if pageSize <= 0 {
		pageSize = options.DefaultPageSize
	}

	// 限制分页大小，防止查询过多数据
	if pageSize > options.MaxPageSize {
		pageSize = options.MaxPageSize
	}

	return &PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

// CalculatePaginationInfo 计算分页信息
func CalculatePaginationInfo(total int, params *PaginationParams) *PaginationInfo {
	if total == 0 {
		return &PaginationInfo{
			CurrentPage: params.Page,
			PageSize:    params.PageSize,
			Total:       0,
			TotalPage:   0,
		}
	}

	// 使用 gpage 创建分页对象
	pager := gpage.New(total, params.PageSize, params.Page, "")

	return &PaginationInfo{
		CurrentPage: pager.CurrentPage,
		PageSize:    params.PageSize,
		Total:       pager.TotalSize,
		TotalPage:   pager.TotalPage,
	}
}

// CalculateOffset 计算偏移量
func CalculateOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}

// CreatePaginationParams 创建分页参数（便捷方法）
func CreatePaginationParams(page, pageSize int) *PaginationParams {
	return &PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

// CreatePaginationParamsWithOptions 使用选项创建分页参数（便捷方法）
func CreatePaginationParamsWithOptions(page, pageSize int, options *PaginationOptions) *PaginationParams {
	params := CreatePaginationParams(page, pageSize)
	return ValidateAndFixPagination(params, options)
}
