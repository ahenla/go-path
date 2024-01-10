package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"fe.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body) // the site will return a stream of bytes.
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("Status code received %v", res.StatusCode)
		//creating a new error message using fmt package
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
