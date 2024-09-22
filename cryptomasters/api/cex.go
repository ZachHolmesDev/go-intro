package api

import (
	"cryptomasters/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 char needed : %v recived", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	response, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err // network error
	}
	var resJson CEXResponse
	if response.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err // something happended trying to read body
		}
		// var cryptoRate datatypes.Rate
		err = json.Unmarshal(bodyBytes, &resJson)
		if err != nil {
			return nil, err
		}

	} else { // status code not 200
		return nil, fmt.Errorf("status code recived: %v", response.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: resJson.Bid}
	return &rate, nil
}
