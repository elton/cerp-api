package controllers

import (
	"net/http"

	"github.com/elton/cerp-api/api/responses"
	"github.com/elton/cerp-api/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllShops returns a list of all the shops.
func (s *Server) GetAllShops(c *fiber.Ctx) error {
	shop := models.Shop{}

	shopsGotton, err := shop.GetAllShops()
	if err != nil {
		responses.ResultJSON(c, http.StatusInternalServerError, nil, err)
		return err
	}
	responses.ResultJSON(c, http.StatusOK, shopsGotton, nil)
	return nil
}

// GetAmountByShop returns the sales amount of each shop.
func (s *Server) GetAmountByShop(c *fiber.Ctx) error {
	shop := models.Shop{}
	start := c.Params("start")
	end := c.Params("end")
	amountsGotton, err := shop.GetAmountByShop(start, end)

	if err != nil {
		responses.ResultJSON(c, http.StatusInternalServerError, nil, err)
		return err
	}
	responses.ResultJSON(c, http.StatusOK, amountsGotton, nil)
	return nil
}
