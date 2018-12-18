package gis

import (
	"net/http"
	"net/url"
)

type Image struct {
	Description string
	Destination string
	Source      string
}

const GOOGLE_URL = "https://www.google.com/search?&tbm=isch&q="

func Search(query string) ([]Image, error) {
	url := GOOGLE_URL + url.QueryEscape(query)
	_, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	images := []Image{}
	return images, nil
}
