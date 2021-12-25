package locations_provider

//package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/herrera-ignacio/go-testing/src/api/domain/locations"
	"github.com/herrera-ignacio/go-testing/src/api/utils/errors"
	"net/http"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(client *resty.Client, countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Printf(urlGetCountry, countryId)
	response, err := client.R().Get(fmt.Sprintf(urlGetCountry, countryId))

	fmt.Println(response)

	if response == nil || err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response on get country %s", countryId),
		}
	}

	if response.StatusCode() == 400 {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("invalid country id %s", countryId),
		}
	}

	if response.StatusCode() == 404 {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("country not found %s", countryId),
		}
	}

	if response.StatusCode() >= 300 {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: "unhandled error",
		}
	}

	var result locations.Country
	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country %s", countryId),
		}
	}

	return &result, nil
}

//func main() {
//	country, _ := GetCountry("AR")
//	fmt.Println(country.Name)
//}
