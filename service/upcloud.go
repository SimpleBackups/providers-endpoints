package service

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getUpcloudStorageRegions(doc *goquery.Document) map[string]string {
	var regionMap map[string]string = make(map[string]string)

	doc.Find(".accordion").First().Find(".accordion-item").Each(func(i int, item *goquery.Selection) {
		// Check if item has an <li> tag with "Object Storage" text
		if strings.Contains(item.Find("li").Text(), "Object Storage") {

			regionCode := strings.ToLower(item.Find("button h3").First().Text())
			regionSplit := strings.Split(item.Find("button .location").First().Text(), ", ")
			regionName := regionSplit[1] + " - " + regionSplit[0] + " - " + regionCode

			regionMap[regionCode] = regionName
		}
	})

	return regionMap
}
func getUpcloudComputeRegions(doc *goquery.Document) map[string]string {

	var regionMap map[string]string = make(map[string]string)

	doc.Find(".accordion").First().Find(".accordion-item").Each(func(i int, item *goquery.Selection) {
		// Check if item has an <li> tag with "Object Storage" text
		if strings.Contains(item.Find("li").Text(), "Cloud Servers") {

			regionCode := strings.ToLower(item.Find("button h3").First().Text())
			regionSplit := strings.Split(item.Find("button .location").First().Text(), ", ")
			regionName := regionSplit[1] + " - " + regionSplit[0] + " - " + regionCode

			regionMap[regionCode] = regionName
		}
	})

	return regionMap
}

func GetUpcloudRegions() Regions {
	doc, err := get("https://upcloud.com/data-centres")
	if err != nil {
		return Regions{}
	}
	storageRegions := getUpcloudStorageRegions(doc)
	computeRegions := getUpcloudComputeRegions(doc)
	return Regions{
		Storage: storageRegions,
		Compute: computeRegions,
	}
}
