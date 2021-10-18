// Package server 服务注入
package server

import (
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/server/user"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	user.New,
)
