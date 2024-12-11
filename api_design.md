# Api design

# Database schema

For persisting user's data as well as its geolocation information, [Postgres](https://www.postgresql.org/) database was chosen. The database schema is very simple and contains only one table for storing all the data.

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

`PostgresDB` is the core structure responsible for handling interaction with postgres database. The struct only contains a connection pool for acquairing connections.

```go
// Struct representing Postgres database controller.
type PostgresDB struct {
	connPool *pgxpool.Pool
}
```

**NOTE:** On the server shutdown, a public `Close` method should be invoked to release all the resources and make sure that no connections could be acquired further.

The database controller exposes two public methods. `AddUser`, which adds a new user together with its geolocation data into a database, and `HasUser` which cheks if a user with a specified `email` address already exists.
These are the signatures for earlier mentioned methods:

```go
func (db *PostgresDB) AddUser(
    firstName,
    lastName,
    password, 
    email string, 
    ipInfo *ipresolver.IpInfo) error {}

func (db *PostgresDB) HasUser(email string) (bool, error) {}

func (db *PostgresDB) Close() error {}
```

# IP resolver

For retrieving user's geolocation data based on its `IP` address I have chosen to use [ipflare](https://www.ipflare.io/) external service. An API key was purchased with `30,000` requests per month, which is more than enough for this assignment.

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

# User data validation

On each request, user's email and password are validated using the following functions.

```go
// Validate user's password. The password should contain at least 10,
// and at most 32 characters. It's a simplified version of how
// password validation could be achieved.
// Returns true if passowrd is valid, false otherwise.
func ValidateUserPassword(pwd string) bool {
    // ...
}

// Validate user's email address. This is a bare minimum validation
// which should never be used in a production, but is sufficient for our purposes.
// Only checks if provided strings contains `@` symbol,
// if so, returns true, false otherwise.
// Since writing a complete email validator is not part of this assignment.
func ValidateUserEmailAddress(email string) bool {
    // ...
}
```

The validation is oversimplified and not meant to be used in a production, especially email address validator.
