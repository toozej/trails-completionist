package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
	"github.com/toozej/trails-completionist/internal/types"
)

// Fetch the web page content
func fetchHTMLContent(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if viper.GetBool("debug") {
		log.Printf("Response object: %v\n", resp)
	}
	return resp, err
}

// Parse the HTML document
func parseHTMLContent(resp *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if viper.GetBool("debug") {
		log.Printf("goquery Document object: %v\n", doc)
	}
	defer resp.Body.Close()
	return doc, err
}

// Extract trail information
func extractTrailInfo(doc *goquery.Document) ([]types.Trail, error) {
	var trails []types.Trail

	// TODO figure out why doc.Find() isn't getting anything
	doc.Find("sc-dPWrhe digRqT").Each(func(i int, s *goquery.Selection) {
		name := s.Find("sc-cjibBx cKMyqg").Text()
		park := s.Find("sc-gYbzsP hksFOR").Text()
		typeAndLength := s.Find("sc-cCjUiG gazdbj").Text()
		url, _ := s.Find("sc-hhOBVt IBdfr").Attr("href")

		trail := types.Trail{
			Name:          name,
			Park:          park,
			TypeAndLength: typeAndLength,
			URL:           url,
		}

		trails = append(trails, trail)
	})
	return trails, nil
}

func ParseTrailsFromHTML(url string) ([]types.Trail, error) {
	resp, err := fetchHTMLContent(url)
	if err != nil {
		return []types.Trail{}, err
	}

	doc, err := parseHTMLContent(resp)
	if err != nil {
		return []types.Trail{}, err
	}

	trails, err := extractTrailInfo(doc)
	if err != nil {
		return []types.Trail{}, err
	}

	return trails, nil
}
