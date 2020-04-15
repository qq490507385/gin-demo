package main

import (
	"doc-system/config"
	"doc-system/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func main() {
	var port = ":" + config.GetStr("port")
	// log
	log.InitWithFile("log.yaml")

	// DB
	// db := config.InitDB()
	// models.AutoMigrate(db)

	// gin mode
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	// Routes
	router.InitRoutes(g)

	fmt.Println(port)

	// localhost
	// log.Infof("Start to listening the incoming requests on https address: %s", config.GetStr("port"))
	log.Info(http.ListenAndServe(port, g).Error())

	// production
	// server := endless.NewServer(port, g)
	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Infof("Server err: %v", err)
	// }
}
