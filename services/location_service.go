package services

import (
	"github.com/black-lotus/go-test/domain/locations"
	"github.com/black-lotus/go-test/providers/location_provider"
	"github.com/black-lotus/go-test/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return location_provider.GetCountry(countryId)
}
