package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/isnastish/nibble/pkg/log"
	"github.com/isnastish/nibble/pkg/validator"
)

func (s *Server) signupRoute(respWriter http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Logger.Error("Failed to read request body: %s", err.Error())
		http.Error(respWriter, "failed to read response body", http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	// extract user's data
	var userData UserData
	if err = json.Unmarshal(body, &userData); err != nil {
		log.Logger.Error("Failed to unmarshal request body: %s", err.Error())
		http.Error(respWriter, fmt.Sprintf("failed to extract user data, error: %s", err), http.StatusInternalServerError)
		return
	}

	if !validator.ValidateUserPassword(userData.Password) {
		log.Logger.Error("Failed to validate user's password: %s", userData.Password)
		http.Error(respWriter, "password validation failed", http.StatusBadRequest)
		return
	}

	if !validator.ValidateUserEmailAddress(userData.Email) {
		log.Logger.Error("Failed to validate user's email address: %s", userData.Email)
		http.Error(respWriter, "email validation failed", http.StatusBadRequest)
		return
	}

	// NOTE: If X-Forwarded-For header contains an ip address,
	// use that instead. Otherwise determine ip address
	// from the request itself. For more information about x-forwarded-for read mozilla docs:
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, err = net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(respWriter, fmt.Sprintf("failed to extract IP address: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	exists, err := s.db.HasUser(userData.Email)
	if err != nil {
		log.Logger.Error("Failed to check if user: %s exists in a database", userData.Email)
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(respWriter, fmt.Sprintf("user with email: %s already exists, try again", userData.Email), http.StatusBadRequest)
		return
	}

	ipInfo, err := s.ipResolverClient.GetGeolocationData(ip)
	if err != nil {
		log.Logger.Error("Failed to get geolocation data: %s", err.Error())
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.db.AddUser(userData.FirstName, userData.LastName, userData.Password, userData.Email, ipInfo); err != nil {
		log.Logger.Error("Failed to add user %s to the database", userData.Email)
		http.Error(respWriter, fmt.Sprintf("failed to add user, error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respWriter.Write([]byte(fmt.Sprintf("Successfully added user: %s", userData.Email)))
}
