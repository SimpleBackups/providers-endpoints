package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getLightsailComputeRegions() map[string]string {
	url := "https://docs.aws.amazon.com/lightsail/latest/userguide/understanding-regions-and-availability-zones-in-amazon-lightsail.html"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("Error loading HTML:", err)
		return nil
	}

	var regionMap map[string]string = make(map[string]string)

	doc.Find(".listitem").Each(func(i int, listItem *goquery.Selection) {
		value := listItem.Find("p").Text()

		// there are two pairs of parentheses in the value, extract the regionCode from the second pair
		regionCode := value[strings.LastIndex(value, "(")+1 : strings.LastIndex(value, ")")]
		regionName := value

		regionMap[regionCode] = regionName
	})

	return regionMap
}

func getLightsailRegions() Regions {
	return Regions{
		Compute: getLightsailComputeRegions(),
	}
}