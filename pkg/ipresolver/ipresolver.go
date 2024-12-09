package ipresolver

import (
	"fmt"
	"os"
)

type Client struct {
}

func NewClient() (*Client, error) {
	IPFLARE_API_KEY, set := os.LookupEnv("IPFLARE_API_KEY")
	if !set || IPFLARE_API_KEY == "" {
		return nil, fmt.Errorf("IPFLARE_API_KEY is not set")
	}

	// get API key from environment variable,
	// if failed, return an error
	// then the client should be assigned to API server
	// for making queries to external service

	return &Client{}, nil
}

// TODO: This could return a struct with a location and region,
// or something similiar, but for now let's stick with a string
// The function should make an http call to external ip service
func (c *Client) Resolve(ipAddr string) (string, error) {
	return "unknown", nil
}

// IPFLARE_API_KEY='d4815a7185da6aae.69f941c643a3f41f751fcc9ef59dcfcfed08a00fb57907b4e750a4a1cdbffc3a'
