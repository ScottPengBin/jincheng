package controller

import (
	"github.com/google/wire"
	"jincheng/app/controller/pers_amount_total"
	"jincheng/app/controller/user"
)

var Provider = wire.NewSet(pers_amount_total.Provider,user.Provider)

