package apisix

import (
	"github.com/goexl/apisix"
	"github.com/pangum/http"
	"github.com/pangum/logging"
)

// Creator 客户端创建者
type Creator struct {
	*apisix.Creator
}

func newCreator(logger *logging.Logger, http *http.Client) (creator *Creator, err error) {
	creator = new(Creator)
	creator.Creator = apisix.New().Logger(logger.Logger).HttpClient(http.Client)

	return
}
