package api

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/handlers"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/config"

	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.RouterGroup, h handlers.Handler, cfg config.Config) {
	r.Use(h.CheckUrl)

	auth := r.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/log-in", h.Login)
		auth.GET("/user-me", h.GetUserMe)
		auth.PUT("/user-me", h.UpdateUserMe)
		auth.DELETE("/user-me", h.DeleteUserMe)

		auth.POST("/user/create", h.CreateUser)
		auth.GET("/user/:id", h.GetUserById)
		auth.GET("/user/list", h.GetAllUsers)
		auth.PUT("/user/:id", h.UpdateUser)
		auth.DELETE("/user/:id", h.DeleteUser)
	}

	product := r.Group("/product")
	{
		product.POST("/item", h.CreateArticle)
		product.GET("/item/:id", h.GetArticleById)
		product.GET("/item/list", h.GetAllArticles)
		product.DELETE("/item/:id", h.DeleteArticle)
		product.PUT("/item/:id", h.UpdateArticle)
	}
}
