package constant

// 启用状态常量
const (
	// StatusEnabled 启用状态
	StatusEnabled = 1
	// StatusDisabled 禁用状态
	StatusDisabled = 0
)

// 启用状态描述
var StatusMap = map[int]string{
	StatusEnabled:  "启用",
	StatusDisabled: "禁用",
}

// IsEnabled 检查是否启用
func IsEnabled(status int) bool {
	return status == StatusEnabled
}

// IsDisabled 检查是否禁用
func IsDisabled(status int) bool {
	return status == StatusDisabled
}
