package gis

import (
	"fmt"
	"io/ioutil"
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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", string(body))
	images := []Image{}
	return images, nil
}
