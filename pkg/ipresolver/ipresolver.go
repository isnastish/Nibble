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
