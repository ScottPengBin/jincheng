package integral

import (
	"github.com/google/wire"
	"jincheng/internal/service/integral"
)

var Provider = wire.NewSet(NewController,integral.Provider)

// Controller 积分
type Controller struct {
	service *integral.Service
}

func NewController(service *integral.Service)  *Controller{
	return &Controller{
		service: service,
	}
}
