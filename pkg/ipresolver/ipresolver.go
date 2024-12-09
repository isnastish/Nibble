package ipresolver

type Client struct {
}

func NewClient() (*Client, error) {
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
