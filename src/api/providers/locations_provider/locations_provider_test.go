package locations_provider

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLocationsProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Locations Suite")
}

func mockClient(fakeUrl string, fixture string, statusCode int) *resty.Client {
	httpmock.Reset()
	client := resty.New()

	httpmock.ActivateNonDefault(client.GetClient())
	responder := httpmock.NewStringResponder(statusCode, fixture)
	httpmock.RegisterResponder(http.MethodGet, fakeUrl, responder)

	return client
}

func TestGetCountryNotFound(t *testing.T) {
	fakeUrl := "https://api.mercadolibre.com/countries/AR"
	fixture := `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []`

	country, apiErr := GetCountry(mockClient(fakeUrl, fixture, http.StatusNotFound), "AR")

	assert.Nil(t, country)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "country not found AR", apiErr.Message)
}

func TestGetCountryRestClientError(t *testing.T) {
	fakeUrl := "https://api.mercadolibre.com/countries/AR"

	country, apiErr := GetCountry(mockClient(fakeUrl, "", http.StatusInternalServerError), "AR")

	assert.Nil(t, country)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusInternalServerError, apiErr.Status)
	assert.EqualValues(t, "unhandled error", apiErr.Message)

}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	fakeUrl := "https://api.mercadolibre.com/countries/AR"
	fixture := `{"id": 123,"name": "Argentina","time_zone": "GMT-03:00", "unknown_key": "I should not be here"}`

	country, apiErr := GetCountry(mockClient(fakeUrl, fixture, http.StatusOK), "AR")

	assert.Nil(t, country)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusInternalServerError, apiErr.Status)
	assert.EqualValues(t, "error when trying to unmarshal country AR", apiErr.Message)
}

func TestGetCountryNoError(t *testing.T) {
	fakeUrl := "https://api.mercadolibre.com/countries/AR"
	fixture := `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`

	country, apiErr := GetCountry(mockClient(fakeUrl, fixture, http.StatusOK), "AR")

	assert.Nil(t, apiErr)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}
