package api

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/carlmjohnson/requests"
	"main_terminal/internal/models"
)

type MoexAPIRequest struct {
	URL    string
	Params map[string][]string
}

type MoexAPIResponse struct {
	Document xml.Name `xml:"document"`
	Data     struct {
		Rows []struct {
			Tradeno   string `xml:"TRADENO,attr" json:"tradeno"`
			Tradetime string `xml:"TRADETIME,attr" json:"tradetime"`
			Price     string `xml:"PRICE,attr" json:"price"`
		} `xml:"rows" json:"stocks"`
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

func (a *MoexAPI) GetStocks(ctx context.Context, companies []models.StockCompany, params map[string][]string) ([]*models.StockCompany, error) {
	var resultedCompanies []*models.StockCompany
	for _, comp := range companies {
		var res MoexAPIResponse
		req := MoexAPIRequest{URL: a.stockURL, Params: params}
		err := requests.URL(req.URL).Pathf("/%s/trades.json", comp.Company).Params(req.Params).ToDeserializer(xml.Unmarshal, &res).Fetch(ctx)

		if err != nil {
			return nil, err
		}

		resJSON, err := json.Marshal(res.Data.Rows)

		if err != nil {
			return nil, err
		}

		var stocksToResult []models.Stock
		err = json.Unmarshal(resJSON, &stocksToResult)

		if err != nil {
			return nil, err
		}

		companyToResult := &models.StockCompany{
			Company: comp.Company,
			Stocks:  stocksToResult,
		}

		resultedCompanies = append(resultedCompanies, companyToResult)

	}

	return resultedCompanies, nil
}
