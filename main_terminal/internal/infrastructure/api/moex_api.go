package api

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/carlmjohnson/requests"
	"main_terminal/internal/models"
)

type MoexAPIRequest struct {
	URL    string
	Params map[string][]string
}

type MoexAPIResponse struct {
	Data struct {
		Rows struct {
			Row []struct {
				Tradeno   string `xml:"TRADENO,attr" json:"tradeno"`
				Tradetime string `xml:"TRADETIME,attr" json:"tradetime"`
				Price     string `xml:"PRICE,attr" json:"price"`
			} `xml:"row"`
		} `xml:"rows"`
	} `xml:"data"`
}

type MoexAPI struct {
	stockURL    string
	currencyURL string
}

func NewMoexAPI(stockURL string, currencyURL string) *MoexAPI {
	return &MoexAPI{
		stockURL:    stockURL,
		currencyURL: currencyURL,
	}
}

func (a *MoexAPI) GetStocks(ctx context.Context, company *models.StockCompany, params map[string][]string) (*models.StockCompany, error) {
	res := new(MoexAPIResponse)
	req := MoexAPIRequest{
		URL:    a.stockURL + fmt.Sprintf("/%s/trades.xml", company.Company),
		Params: params,
	}

	err := requests.URL(req.URL).Params(req.Params).ToDeserializer(xml.Unmarshal, &res).Fetch(ctx)

	if err != nil {
		return nil, err
	}

	resJSON, err := json.Marshal(res.Data.Rows.Row)

	if err != nil {
		return nil, err
	}

	var stocksToResult []models.Stock
	err = json.Unmarshal(resJSON, &stocksToResult)

	if err != nil {
		return nil, err
	}

	companyToResult := &models.StockCompany{
		Company: company.Company,
		Stocks:  stocksToResult,
	}

	return companyToResult, nil
}
