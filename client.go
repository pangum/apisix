package apisix

import (
	"github.com/goexl/apisix"
)

// Client 客户端，主要是为了解决包冲突
type Client struct {
	*apisix.Client
}
