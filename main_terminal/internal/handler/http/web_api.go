package http_custom

import (
	"github.com/gin-gonic/gin"
	"main_terminal/internal/models"
	"net/http"
)

type webApiBodyIncoming struct {
	Company string `json:"company"`
}

func (h *GinHttp) stocksShortGet(c *gin.Context) {
	var body *webApiBodyIncoming

	err := c.BindJSON(&body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	compStock := models.StockCompany{
		Company: body.Company,
	}

	h.service.GetStockList(c)
}
