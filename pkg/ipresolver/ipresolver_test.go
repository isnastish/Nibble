package ipresolver

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

var testIpAddresses = []*IpInfo{
	{
		Ip:          "34.21.9.50",
		City:        "Washington",
		Country:     "United States",
		CountryCode: "US",
	},
	{
		Ip:          "34.106.208.213",
		City:        "Salt Lake City",
		Country:     "United States",
		CountryCode: "US",
	},
	{
		Ip:          "34.130.107.20",
		City:        "Toronto",
		Country:     "Canada",
		CountryCode: "CA",
	},
	{
		Ip:          "34.39.131.22",
		City:        "Sao Paulo",
		Country:     "Brazil",
		CountryCode: "BR",
	},
	{
		Ip:          "34.240.49.81",
		City:        "Dublin",
		Country:     "Ireland",
		CountryCode: "IE",
	},
	// TODO: Add country codes once the plan is updated
	// since currently we get QUOTA_EXCEEDED error from API server
	{
		Ip:      "35.242.177.6",
		City:    "London",
		Country: "United Kingdom",
	},
	{
		Ip:      "13.36.154.207",
		City:    "Paris",
		Country: "France",
	},
	{
		Ip:      "34.91.238.70",
		City:    "Amsterdam",
		Country: "Netherlands",
	},
	{
		Ip:      "34.159.56.80",
		City:    "Frankfurt",
		Country: "Germany",
	},
}

func TestResolveIpAddress(t *testing.T) {
	os.Setenv("IPFLARE_API_KEY", "d4815a7185da6aae.69f941c643a3f41f751fcc9ef59dcfcfed08a00fb57907b4e750a4a1cdbffc3a")

	ipResolverClient, err := NewClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < len(testIpAddresses); i++ {
		ipAddr := testIpAddresses[i].Ip

		ipInfo, err := ipResolverClient.Resolve(ipAddr)
		if err != nil {
			t.Errorf(err.Error())
		}

		fmt.Printf("Country: %s\n", ipInfo.Country)
		fmt.Printf("City: %s\n", ipInfo.City)
		fmt.Printf("Country code: %s\n", ipInfo.CountryCode)
	}
}
