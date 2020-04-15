package main

// import (
// 	"doc-system/app/models"
// 	"doc-system/config"
// 	"doc-system/router"

// 	"github.com/fvbock/endless"
// 	"github.com/gin-gonic/gin"
// 	"github.com/lexkong/log"
// )

// func main() {
// 	var port = ":" + config.GetStr("port")
// 	// log
// 	log.InitWithFile("log.yaml")

// 	// DB
// 	db := config.InitDB()
// 	models.AutoMigrate(db)

// 	// gin mode
// 	gin.SetMode(gin.ReleaseMode)

// 	g := gin.New()

// 	// Routes
// 	router.InitRoutes(g)

// 	server := endless.NewServer(port, g)
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Infof("Server err: %v", err)
// 	}
// }
