package http_custom

import (
	"github.com/gin-gonic/gin"
	"main_terminal/internal/models"
	"net/http"
)

type webApiBodyIncoming struct {
	Companies []string            `json:"companies"`
	Params    map[string][]string `json:"params"`
}

type webApiBodyOutcoming struct {
	CompaniesWithStocks []*models.StockCompany `json:"companiesWithStocks"`
}

func (h *GinHttp) stocksShortGet(c *gin.Context) {
	var bodyIn *webApiBodyIncoming

	err := c.BindJSON(&bodyIn)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var bodyOut *webApiBodyOutcoming

	for _, company := range bodyIn.Companies {
		compStock := &models.StockCompany{
			Company: company,
		}

		compStock, err = h.service.GetStock(c, compStock, bodyIn.Params)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		bodyOut.CompaniesWithStocks = append(bodyOut.CompaniesWithStocks, compStock)

	}

	c.JSON(http.StatusOK, bodyOut)
}
