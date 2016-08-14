package server

import (
	"github.com/fvbock/endless"
	"github.com/hydnwang/elramir/config"
	"github.com/hydnwang/elramir/model"
)

func init() {
	model.InitDB()
}

func RunHTTPServer() error {
	var err error

	if config.Mode == "debug" {
		RoutersEngine().Run(":" + config.Port)
	} else {
		err = endless.ListenAndServe(":"+config.Port, RoutersEngine())
	}
	return err
}
