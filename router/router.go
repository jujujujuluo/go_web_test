package router

import (
	"exchangeapp/controller"
	"exchangeapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("register", controller.Register)
	}

	api := r.Group("/api")
	api.GET("exchangeRates", controller.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("exchangeRates", controller.CreateExchangeRate)
		api.POST("/articles", controller.CreateArticle)
		api.GET("/articles", controller.GetArticles)
		api.GET("/articles/:id", controller.GetArticlesById)

		api.POST("/articles/:id/like", controller.LikeArticle)
		api.GET("/articles/:id/like", controller.GetArticleLikes)
	}
	return r
}
