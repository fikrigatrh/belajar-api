package application

import (
	"final_project/config/env"
	"final_project/router"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

// TODO CREATE FUNC init
func init()  {
	//load env
	err := gotenv.Load()
	if err != nil {
		return
	}
}

// StartApp TODO CREATE FUNC START APP
func StartApp()  {
	// TODO call function router in folder router
	addr := env.Config.ServiceHost + ":" + env.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}