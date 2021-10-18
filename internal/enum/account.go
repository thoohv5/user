package enum

// AccountType 账号类型
type AccountType = int32

const (
	// User 账号密码
	User AccountType = iota
	// Phone 手机账号
	Phone
)
