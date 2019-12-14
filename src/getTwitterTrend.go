package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	anaconda.SetConsumerKey("339ttP5v0cWVKlBmkIeyfSjmk")
	anaconda.SetConsumerSecret("MZJXV2SbTiNxeD0yf7dhortxX69IdP78ADlP4kPxrN9jdFssiV")
	api := anaconda.NewTwitterApi("955371524403298304-sKHG03JGLUoIJ13m9LgwtLWyacT3K4f", "BjtC6S7ger7RiOm8gGjo0MbUUn6UuYUajOvmnleCUK0FN")

	id := int64(23424856)
	v := url.Values{}

	trends, err := api.GetTrendsByPlace(id, v)
	if err != nil {
		panic(err)
	}

	fmt.Println(trends)

}