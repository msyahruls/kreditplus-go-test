package main

import (
	"log"

	"github.com/msyahruls/kreditplus-go-test/internal/config"
	"github.com/msyahruls/kreditplus-go-test/internal/handler"
	"github.com/msyahruls/kreditplus-go-test/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/msyahruls/kreditplus-go-test/docs"
)

func main() {
	db := config.ConnectDB()
	router := config.SetupRouter()

	router.Use(middleware.ErrorFormatterMiddleware())  // error formatter middleware
	router.Use(middleware.SecurityHeadersMiddleware()) // security headers middleware
	router.Use(middleware.RateLimiterMiddleware())     // rate limiter middleware

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler.InitRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
