package account

import (
	"github.com/thoohv5/template/internal/enum"
	"github.com/thoohv5/template/internal/service/user/account/base"
	"github.com/thoohv5/template/internal/service/user/account/standard"
	"github.com/thoohv5/template/internal/service/user/account/user"
)

// Adapter 账号适配
func Adapter(accountType enum.AccountType, opts ...base.Option) standard.IAccount {
	var account standard.IAccount
	switch accountType {
	case enum.User:
		account = user.New(opts...)
	case enum.Phone:

	}
	return account
}
