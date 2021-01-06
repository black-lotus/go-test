package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/black-lotus/go-test/domain/locations"
	"github.com/black-lotus/go-test/services"
	"github.com/black-lotus/go-test/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country locations.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
