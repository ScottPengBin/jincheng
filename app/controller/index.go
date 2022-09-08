package controller

import (
	"github.com/google/wire"
	"jincheng/app/pers_amount_total"
)

var Provider = wire.NewSet(pers_amount_total.NewController)

