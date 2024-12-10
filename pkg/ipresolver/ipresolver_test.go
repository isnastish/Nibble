package ipresolver

import (
	"testing"

	_ "github.com/stretchr/testify/assert"
)

var testIpAddresses = []*IpInfo{
	{
		Ip:      "34.21.9.50",
		City:    "Washington",
		Country: "United States",
	},
	{
		Ip:      "34.106.208.213",
		City:    "Salt Lake City",
		Country: "United States",
	},
	{
		Ip:      "34.130.107.20",
		City:    "Toronto",
		Country: "Canada",
	},
	{
		Ip:      "34.39.131.22",
		City:    "Sao Paulo",
		Country: "Brazil",
	},
	{
		Ip:      "34.240.49.81",
		City:    "Dublin",
		Country: "Ireland",
	},
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

		if ipInfo.Country != testIpAddresses[i].Country {
			t.Errorf("invalid geolocation data for IP: %s. Expected: %s, got: %s", ipAddr, testIpAddresses[i].Country, ipInfo.Country)
		}
	}
}
