package ipresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/isnastish/nibble/pkg/log"
)

// Struct containing IP's geolocation data
// retrieved from making a request to external service
// and parsing its response body.
// It could contain an optional error code and an error message.
type IpInfo struct {
	Ip          string `json:"ip"`
	City        string `json:"city"`
	Region      string `json:"region"`
	RegionCode  string `json:"region_code"`
	Country     string `json:"country_name"`
	CountryCode string `json:"country_code"`

	// error handling
	ErrorCode string `json:"code,omitempty"`
	ErrorMsg  string `json:"error,omitempty"`
}

// A user facing client for interacting with an external
// service for resolving geolocation.
type Client struct {
	// http client
	*http.Client
	// ipflare API key retrieved from env variable
	ipflareApiKey string
}

func NewClient() (*Client, error) {
	ipflareApiKey, set := os.LookupEnv("IPFLARE_API_KEY")
	if !set || ipflareApiKey == "" {
		return nil, fmt.Errorf("IPFLARE_API_KEY is not set")
	}

	return &Client{
		Client:        &http.Client{},
		ipflareApiKey: ipflareApiKey,
	}, nil
}

// Make a request to an external service `ipflare`: https://www.ipflare.io/
// Parse its response and retreive geolocation data based on provided
// ip address.
// Return an error, if any, otherwise an instance of `IpInfo` containing the necessary information.
func (c *Client) GetGeolocationData(ipAddr string) (*IpInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.ipflare.io/%s", ipAddr), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a request: %s", err.Error())
	}

	req.Header.Add("X-API-Key", c.ipflareApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %s", err.Error())
	}

	var ipInfo IpInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response data: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s", ipInfo.ErrorCode, ipInfo.ErrorMsg)
	}

	log.Logger.Info("Got geolocation for IP: %s, city: %s, country: %s", ipAddr, ipInfo.City, ipInfo.Country)

	return &ipInfo, nil
}
