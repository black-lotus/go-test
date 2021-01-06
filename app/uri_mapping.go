package app

import (
	"github.com/black-lotus/go-test/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
