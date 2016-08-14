package server

import (
	"github.com/elramir/config"
	"github.com/elramir/model"
	"github.com/fvbock/endless"
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
