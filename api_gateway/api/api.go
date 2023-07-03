package api

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/handlers"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	//item
	r.POST("/article", h.CreateArticle)
	r.GET("/article/:id", h.GetArticleById)
	r.GET("/article", h.GetAllArticles)
	r.DELETE("/article/:id", h.DeleteArticle)
	r.PUT("/article/:id", h.UpdateArticle)
	r.PATCH("/article/:id", h.PatchArticle)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
