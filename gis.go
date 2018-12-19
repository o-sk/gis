package gis

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	ID          string `json:"id"`
	Cite        string `json:"rh"`
	Description string `json:"pt"`
	Destination string `json:"ru"`
	Extension   string `json:"ity"`
	Source      string `json:"ou"`
	Thumbnail   string `json:"tu"`
}

type DownloadResult struct {
	FileName string
	Image    Image
	Error    error
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

func downloadRequest(directory string, image Image) DownloadResult {
	response, err := http.Get(image.Source)
	if err != nil {
		return DownloadResult{FileName: "", Image: image, Error: err}
	}
	defer response.Body.Close()
	filename := filepath.Join(directory, image.ID)
	file, err := os.Create(filename)
	if err != nil {
		return DownloadResult{FileName: filename, Image: image, Error: err}
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return DownloadResult{FileName: filename, Image: image, Error: nil}
}

func Download(directory, filename string, images []Image) []DownloadResult {
	download := func(done <-chan interface{}, imageStream <-chan Image) <-chan DownloadResult {
		downloadStream := make(chan DownloadResult)
		go func() {
			defer close(downloadStream)
			for image := range imageStream {
				select {
				case <-done:
					return
				case downloadStream <- downloadRequest(directory, image):
				}
			}
		}()
		return downloadStream
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan DownloadResult) <-chan DownloadResult {
		var wg sync.WaitGroup
		multiplexedStream := make(chan DownloadResult)
		mulitiplex := func(c <-chan DownloadResult) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go mulitiplex(c)
		}
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan DownloadResult,
		num int,
	) <-chan DownloadResult {
		takeStream := make(chan DownloadResult)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	baseFileName := filepath.Join(directory, filename)
	imageStream := make(chan Image, len(images))
	defer close(imageStream)
	for _, image := range images {
		imageStream <- image
	}

	done := make(chan interface{})
	defer close(done)

	numChannel := runtime.NumCPU()
	downloaders := make([]<-chan DownloadResult, numChannel)
	for i := 0; i < numChannel; i++ {
		downloaders[i] = download(done, imageStream)
	}
	downloadResults := make([]DownloadResult, len(images))
	count := 1
	for downloadResult := range take(done, fanIn(done, downloaders...), len(images)) {
		if downloadResult.Error == nil {
			var err error
			format := downloadResult.Image.Extension
			if format == "" {
				f, _ := os.Open(downloadResult.Image.ID)
				defer f.Close()
				_, format, err = image.DecodeConfig(f)
				if err != nil {
					downloadResult.Error = err
				}
			}
			fn := baseFileName + strconv.Itoa(count) + "." + format
			if err = os.Rename(downloadResult.FileName, fn); err != nil {
				downloadResult.Error = err
			} else {
				downloadResult.FileName = filename
			}
		}
		downloadResults = append(downloadResults, downloadResult)
		count++
	}
	return downloadResults
}
