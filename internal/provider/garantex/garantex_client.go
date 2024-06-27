package garantex

import (
	"Gl0ven/kata_projects/rates/internal/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type GarantexApi interface {
	GetRates() (models.Rates, error)
}

type garantexProvider struct {
	client *http.Client
	marketId string
}

var getRatesUrl = "https://garantex.org/api/v2/depth"

func NewGarantexProvider(mId string) GarantexApi {
	client := &http.Client{}
	return &garantexProvider{
		client: client,
		marketId: mId,
	}
}

func(p *garantexProvider) GetRates() (models.Rates, error) {
	query := url.Values{}
	query.Add("market", p.marketId)
	
	body, err := p.httpRequest(http.MethodGet, query, nil)
	if err != nil {
		return models.Rates{}, err
	}
	
	var rates Rates
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return models.Rates{}, err
	}
	
	mRates, err := convertRespToModel(rates)

	return mRates, err
}

func (p *garantexProvider) httpRequest(method string, query url.Values, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, getRatesUrl, body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return nil, errors.New(string(respBody))
	}

	return respBody, err
}

func convertRespToModel(rates Rates) (models.Rates, error) {
	mRates := models.Rates{Timestamp: int(rates.Timestamp)}
	if len(rates.Asks) != 0 {
		ask, err := strconv.ParseFloat(rates.Asks[0].Price, 64)
		if err != nil {
			return models.Rates{}, err
		}
		mRates.AskPrice = ask
	}

	if len(rates.Bids) != 0 {
		bid, err := strconv.ParseFloat(rates.Bids[0].Price, 64)
		if err != nil {
			return models.Rates{}, err
		}
		mRates.BidPrice = bid
	}

	return mRates, nil
}