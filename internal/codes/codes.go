package codes

import "github.com/pkg/errors"

var (
	// ErrDbException 数据库异常
	ErrDbException = errors.New("100100")
	// ErrDbTransactionException 数据库事务异常
	ErrDbTransactionException = errors.New("100100")

	// ErrTypeConvertException 类型转换异常
	ErrTypeConvertException = errors.New("100110")

	// ErrHavingExist 用户已经存在
	ErrHavingExist = errors.New("200100")
)
