package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/isnastish/nibble/pkg/validator"
)

func (s *Server) signupRoute(respWriter http.ResponseWriter, req *http.Request) {
	/*
		"Content-Type": application/json
		"Body"
		{
			first_name: "Alexey",
			last_name: "Yevtushenko",
			password: isnastish@234,
			email: isnastish@gmail.com
		}
	*/

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(respWriter, "failed to read response body", http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	// get user's data
	var userData UserData
	if err = json.Unmarshal(body, &userData); err != nil {
		http.Error(respWriter, "failed to extract user data", http.StatusInternalServerError)
		return
	}

	// Separate IP address from the port
	ipInfo, err := s.ipResolverClient.Resolve(strings.Split(req.RemoteAddr, ":")[0])
	if err != nil {
		http.Error(respWriter)
		return
	}

	if !validator.ValidateUserPassword(userData.Password) {
		http.Error(respWriter, "password validation failed", http.StatusInternalServerError)
		return
	}

	if !validator.ValidateUserEmailAddress(userData.Email) {
		http.Error(respWriter, "email validation failed", http.StatusInternalServerError)
		return
	}

	if err := s.db.AddUser(userData.FirstName, userData.LastName, userData.Password, userData.Email, ipInfo); err != nil {
		http.Error(respWriter, fmt.Sprintf("failed to add user, error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respWriter.WriteHeader(http.StatusOK)
}

func (s *Server) loginRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello from login route")
}
