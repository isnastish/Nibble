package ipresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

type Client struct {
	// http client
	*http.Client
	// ipflare API key retrieved from env variable
	ipflareApiKey string
}

func NewClient() (*Client, error) {
	IPFLARE_API_KEY, set := os.LookupEnv("IPFLARE_API_KEY")
	if !set || IPFLARE_API_KEY == "" {
		return nil, fmt.Errorf("IPFLARE_API_KEY is not set")
	}

	return &Client{
		Client:        &http.Client{},
		ipflareApiKey: IPFLARE_API_KEY,
	}, nil
}

func (c *Client) Resolve(ipAddr string) (*IpInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.ipflare.io/%s", ipAddr), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a request: %s", err.Error())
	}

	req.Header.Add("X-API-Key", c.ipflareApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode)

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

	return &ipInfo, nil
}
