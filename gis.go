package gis

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	Cite        string `json:"rh"`
	Description string `json:"pt"`
	Destination string `json:"ru"`
	Source      string `json:"ou"`
	Thumbnail   string `json:"tu"`
}

const GOOGLE_URL = "https://www.google.com/search?&tbm=isch&q="

func Search(query string) ([]Image, error) {
	url := GOOGLE_URL + url.QueryEscape(query)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	images := make([]Image, 100)
	doc.Find(".rg_bx.rg_di.rg_el.ivg-i").Each(func(i int, s *goquery.Selection) {
		var image Image
		meta := s.Find(".rg_meta.notranslate").Text()
		json.Unmarshal([]byte(meta), &image)
		images[i] = image
	})
	return images, nil
}
