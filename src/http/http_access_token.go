package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sebagalan/bookstore_oauth-api/src/domain/access_token"
	"github.com/sebagalan/bookstore_oauth-api/src/services"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

// AccessTokenHandler ...
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
}

type accessTokenHandler struct {
	service services.AccessTokenServices
}

//NewHandler ....
func NewHandler(service services.AccessTokenServices) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

//GetByID ..
func (ath *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := ath.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (ath *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		errJson := errors.NewBadRequestError("some params are wrong")
		c.JSON(errJson.Status, errJson)
		return
	}

	err := ath.service.Create(at)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}

func (ath *accessTokenHandler) Update(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		errJson := errors.NewBadRequestError("some params are wrong")
		c.JSON(errJson.Status, errJson)
		return
	}

	err := ath.service.Update(at)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
