# Api design
The core component of the whole system is a `Server` struct.
It holds a database contoller for persisting user data in an external storage, and an `IP` resolver client for retrieving user's geolocation based on its ip address. 
```go
type Server struct {
	// http server
	*http.Server
	// ip resolver client
	ipResolverClient *ipresolver.Client
	// database connector
	db *db.PostgresDB
	// settings
	port int
}
```
It exposes a single API endpoint `http://localhost:3030/signup` with a `POST` method. The endpoint accepts user's data in the following form. 
```json
{
 "first_name": "Alexey",
 "last_name": "Yevtusnenko",
 "password": "234324@sd22ad",
 "email": "isnastish@gmail.com"
}
```
If you wish to specify an `IP` address in the request itself, pass the following header: `X-Forwarded-For: "your-IP-address"`. Otherwise, it will be deducted automatically from the `URL`.

**NOTE:**  When running the service locally, the default ip address, deducted from the `URL` would be `127.0.0.1`, which will cause `RESERVED_IP_ADDRESS` when tried to retrieve its geolocation data from an `ipflare` service. In order to avoid that situation and test the application thoroughly, use the earlier mentioned header and pass any ip address you want. 
Here are some examples: 
```json
{
    "34.21.9.50": { "City": "Washington", "Country": "United States" },
    "34.130.107.20": { "City": "Toronto", "Country": "Canada" },
    "34.39.131.22": { "City": "Sao Paulo", "Country": "Brazil" }
}
```
**NOTE:** For a complete usage example with `curl` take a look at [readme document](/README.md).

The server exposes two public methods, `Serve` which boots up an http server and listens for incoming requests on the specified port, the default is `3030`, and `Shutdown` method, which should be invoked on shutdown of the application. It closes the database connection and does the graceful shutdown. 
```go
func (s *Server) Serve() error {}

func (s *Server) Shutdown() error {}
```

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
func ValidateUserPassword(pwd string) bool {}
func ValidateUserEmailAddress(email string) bool {}
```

The validation is oversimplified and not meant to be used in a production, especially email address validator.
