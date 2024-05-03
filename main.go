package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type ProviderRegions struct {
	Provider string
	Regions  Regions
}

type Regions struct {
	Storage map[string]string `json:"storage"`
	Compute map[string]string `json:"compute"`
}

func getRegions() map[string]Regions {
	var wg sync.WaitGroup
	providerRegions := make(chan ProviderRegions)

	providers := []struct {
		name string
		fn   func() Regions
	}{
		{"AWS", service.getAmazonRegions},
		{"LIGHTSAIL", service.getLightsailRegions},
		{"DIGITALOCEAN", service.getDigitalOceanRegions},
		{"UPCLOUD", service.getUpcloudRegions},
		{"EXOSCALE", service.getExoscaleRegions},
		{"WASABI", service.getWasabiRegions},
		{"GOOGLE_CLOUD", service.getGoogleCloudRegions},
		{"BACKBLAZE", service.getBackblazeRegions},
		{"LINODE", service.getLinodeRegions},
		{"OUTSCALE", service.getOutscaleRegions},
		{"STORJ", service.getStorjRegions},
		{"VULTR", service.getVultrRegions},
	}

	for _, provider := range providers {
		wg.Add(1)
		go func(provider struct {
			name string
			fn   func() Regions
		}) {
			defer wg.Done()
			providerRegions <- ProviderRegions{Provider: provider.name, Regions: provider.fn()}
		}(provider)
	}

	go func() {
		wg.Wait()
		close(providerRegions)
	}()

	regions := make(map[string]Regions)
	for pr := range providerRegions {
		regions[pr.Provider] = pr.Regions
	}

	return regions
}

func main() {
	regions := getRegions()
	regionsJson, err := json.Marshal(regions)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(regionsJson))
}
