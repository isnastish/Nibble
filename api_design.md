# Api design 

# Database schema
```sql
CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL, 
    "first_name" VARCHAR(64) NOT NULL, 
    "last_name" VARCHAR(64) NOT NULL,
    "password" CHAR(64) NOT NULL, 
    "email" VARCHAR(64) NOT NULL UNIQUE,
    "city" VARCHAR(64) NOT NULL,
    "country" VARCHAR(64) NOT NULL,
    PRIMARY KEY("id")
);
```

# IP resolver
For retrieving user's geolocation data based on its `IP` address I have choosen to use [ipflare](https://www.ipflare.io/) external service. An API key was purchased with `30,000` requests per month, which is more than enough for this assignment.

The core structure responsible for communication with an ipflare service is a `Client` holding an `http.Client` instance for making http requests, and an API key string. 
```go
// A user facing client for interacting with an external
// service for resolving geolocation.
type Client struct {
	// http client
	*http.Client
	// ipflare API key retrieved from env variable
	ipflareApiKey string
}
```   
The client has a single public method `GetGeolocationData`, which makes an actual request to `ipflare` server and parser the response into `IpInfo` struct, which contains all the metadata, like country name, city, country code, region, etc., as well as potential errors occured down the road.
The `IpInfo` structure has the following schema:
```go
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
```
If the response status doesn't equal `200`, the `ErrorCode` and `ErrorMsg` fields will contain the necessary information which is further propagated to the client.

# User input validation 