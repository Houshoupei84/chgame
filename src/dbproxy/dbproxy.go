//go:binary-only-package-my

package dbproxy

import "exportor/defines"

func NewDbProxy() defines.IDbProxy {
	return newDBProxyServer()
}