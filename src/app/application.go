package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sebagalan/bookstore_oauth-api/src/http"
	db_respository "github.com/sebagalan/bookstore_oauth-api/src/repository"
	"github.com/sebagalan/bookstore_oauth-api/src/services"
)

var (
	router = gin.Default()
)

// StartAplication ...
func StartAplication() {

	atHandler := http.NewHandler(services.NewAccessTokenServices(db_respository.NewDbRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.PUT("/oauth/access_token/:access_token_id", atHandler.Update)

	router.Run(":10080")
}
