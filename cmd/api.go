package main

import (
	"jatis/pkg/config"
	"jatis/pkg/router"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type option struct {
	Addr string
}

func runAPI(opt option) {
	e := echo.New()
	router.InitRouter(e)

	server := http.Server{
		Handler:      e,
		Addr:         config.Cfg.Server.Addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}
