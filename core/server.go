package core

import (
	"gin-basic-framework/global"
	"gin-basic-framework/initialize"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) server {

	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//return newServer
}

func StartServer() {
	router := initialize.Router()
	address := global.GLOBAL_CONFIG.Server.AddressProvider()
	server := InitServer(address, router)
	server.ListenAndServe()
}
