package controllers

import (
	"net/http"

	"github.com/elton/cerp-api/api/responses"
	"github.com/elton/cerp-api/models"
	"github.com/gin-gonic/gin"
)

// GetAllShops returns a list of all the shops.
func (s *Server) GetAllShops(ctx *gin.Context) {
	shop := models.Shop{}

	shopsGotton, err := shop.GetAllShops()
	if err != nil {
		responses.ResultJSON(ctx, http.StatusInternalServerError, nil, err)
	}
	responses.ResultJSON(ctx, http.StatusOK, shopsGotton, nil)
}
