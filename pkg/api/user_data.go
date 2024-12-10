package api

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`

	// If User provides an API address, use that instead to
	// determine the geolocation, otherwise deduct the ip address
	// from a request
	Ip string `json:"ip,omitempty"`
}
