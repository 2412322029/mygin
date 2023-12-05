package main

import (
	"mygin/pkg/cfg"
	"mygin/pkg/log"
	"mygin/pkg/router"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.InitLogger(cfg.GetConfig().Log.Path, cfg.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", cfg.GetConfig()))
	log.Logger.Info("start server")
	router := router.NewRouter()
	router.Use(gin.Logger())
	s := &http.Server{
		Addr:           cfg.GetConfig().Server.Addr + ":" + strconv.Itoa(cfg.GetConfig().Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Logger.Error(err.Error())
	}
}
