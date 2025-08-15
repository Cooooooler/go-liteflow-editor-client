# 分页工具包 (Pagination Utility)

这是一个用于处理分页逻辑的 Go 工具包，提供了分页参数验证、分页信息计算、偏移量计算等功能。

## 功能特性

- ✅ 分页参数验证和修复
- ✅ 分页信息计算
- ✅ 偏移量计算
- ✅ 自定义分页选项
- ✅ 边界情况处理
- ✅ 便捷方法支持

## 快速开始

### 基本使用

```go
import "go-liteflow-editor-client/utility"

// 创建分页参数
params := utility.CreatePaginationParams(2, 20)

// 验证和修复分页参数
validatedParams := utility.ValidateAndFixPagination(params, nil)

// 计算分页信息
paginationInfo := utility.CalculatePaginationInfo(100, validatedParams)

// 计算偏移量
offset := utility.CalculateOffset(validatedParams.Page, validatedParams.PageSize)
```

### 在控制器中使用

```go
func (c *Controller) GetList(ctx context.Context, req *GetListReq) (*GetListRes, error) {
    // 使用分页工具包验证和修复分页参数
    paginationParams := utility.ValidateAndFixPagination(
        &utility.PaginationParams{
            Page:     req.Page,
            PageSize: req.PageSize,
        },
        nil, // 使用默认选项
    )

    // 构建查询模型
    model := dao.YourModel.Ctx(ctx)

    // 获取总数
    total, err := model.Count()
    if err != nil {
        return nil, err
    }

    // 计算偏移量
    offset := utility.CalculateOffset(paginationParams.Page, paginationParams.PageSize)

    // 分页查询
    var items []YourEntity
    err = model.Page(offset, paginationParams.PageSize).Scan(&items)
    if err != nil {
        return nil, err
    }

    // 计算分页信息
    paginationInfo := utility.CalculatePaginationInfo(total, paginationParams)

    // 返回结果
    return &GetListRes{
        Data: items,
        PageInfo: paginationInfo,
    }, nil
}
```

## API 参考

### 结构体

#### PaginationParams
分页参数结构体

```go
type PaginationParams struct {
    Page     int `json:"page"`     // 当前页码
    PageSize int `json:"pageSize"` // 每页大小
}
```

#### PaginationInfo
分页信息结构体

```go
type PaginationInfo struct {
    CurrentPage int `json:"currentPage"` // 当前页码
    PageSize    int `json:"pageSize"`    // 每页大小
    Total       int `json:"total"`       // 总记录数
    TotalPage   int `json:"totalPage"`   // 总页数
}
```

#### PaginationOptions
分页选项结构体

```go
type PaginationOptions struct {
    DefaultPage     int `json:"defaultPage"`     // 默认页码
    DefaultPageSize int `json:"defaultPageSize"` // 默认每页大小
    MaxPageSize     int `json:"maxPageSize"`     // 最大每页大小
}
```

### 函数

#### ValidateAndFixPagination
验证并修复分页参数

```go
func ValidateAndFixPagination(params *PaginationParams, options *PaginationOptions) *PaginationParams
```

**参数:**
- `params`: 原始分页参数
- `options`: 分页选项，如果为 nil 则使用默认选项

**返回值:**
- 验证和修复后的分页参数

**功能:**
- 设置默认页码和每页大小
- 限制最大每页大小
- 处理无效参数（负数、零值等）

#### CalculatePaginationInfo
计算分页信息

```go
func CalculatePaginationInfo(total int, params *PaginationParams) *PaginationInfo
```

**参数:**
- `total`: 总记录数
- `params`: 分页参数

**返回值:**
- 完整的分页信息

#### CalculateOffset
计算偏移量

```go
func CalculateOffset(page, pageSize int) int
```

**参数:**
- `page`: 页码
- `pageSize`: 每页大小

**返回值:**
- 数据库查询的偏移量

#### CreatePaginationParams
创建分页参数

```go
func CreatePaginationParams(page, pageSize int) *PaginationParams
```

#### CreatePaginationParamsWithOptions
使用选项创建分页参数

```go
func CreatePaginationParamsWithOptions(page, pageSize int, options *PaginationOptions) *PaginationParams
```

## 默认配置

```go
var DefaultPaginationOptions = PaginationOptions{
    DefaultPage:     1,   // 默认页码
    DefaultPageSize: 10,  // 默认每页大小
    MaxPageSize:     100, // 最大每页大小
}
```

## 使用场景

### 1. 基本分页查询
适用于大多数列表查询场景，使用默认配置。

### 2. 自定义分页配置
适用于需要特殊分页规则的场景，如管理后台、报表等。

### 3. 批量数据处理
适用于需要处理大量数据的场景，通过限制每页大小防止内存溢出。

## 最佳实践

1. **始终验证分页参数**: 使用 `ValidateAndFixPagination` 确保参数有效
2. **设置合理的分页大小**: 避免单次查询过多数据
3. **处理边界情况**: 注意处理空数据、无效参数等情况
4. **统一分页逻辑**: 在项目中统一使用此工具包，保持代码一致性

## 示例代码

查看 `pagination_example.go` 文件获取更多使用示例。

## 注意事项

- 页码从 1 开始计数
- 每页大小必须大于 0
- 总页数会根据总记录数和每页大小自动计算
- 偏移量计算基于 0 开始的索引 