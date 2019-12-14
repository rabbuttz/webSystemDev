package getTwitterTrend

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func GetTrends(locate_id int64) (trnd_name.Name []anaconda.Trend) {

	anaconda.SetConsumerKey("339ttP5v0cWVKlBmkIeyfSjmk")
	anaconda.SetConsumerSecret("MZJXV2SbTiNxeD0yf7dhortxX69IdP78ADlP4kPxrN9jdFssiV")
	api := anaconda.NewTwitterApi("955371524403298304-sKHG03JGLUoIJ13m9LgwtLWyacT3K4f", "BjtC6S7ger7RiOm8gGjo0MbUUn6UuYUajOvmnleCUK0FN")

	id := int64(locate_id)
	v := url.Values{}

	trends, err := api.GetTrendsByPlace(id, v)
	if err != nil {
		panic(err)
	}

	trnd := trends.Trends
	trnd_name := trnd.Name

	/*
	for _, trnd := range trends.Trends {
		return trnd
	}
	*/

	return trnd_name
}
