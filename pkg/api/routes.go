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
		http.Error(respWriter, fmt.Sprintf("failed to extract user data, error: %s", err), http.StatusInternalServerError)
		return
	}

	log.Logger.Info("user data: %v", userData)

	if !validator.ValidateUserPassword(userData.Password) {
		http.Error(respWriter, "password validation failed", http.StatusInternalServerError)
		return
	}

	if !validator.ValidateUserEmailAddress(userData.Email) {
		http.Error(respWriter, "email validation failed", http.StatusInternalServerError)
		return
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		http.Error(respWriter, fmt.Sprintf("failed to extract IP address: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	log.Logger.Info("ip: %s", ip)

	ipInfo, err := s.ipResolverClient.Resolve(ip)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Logger.Info("ip info: %v", ipInfo)

	// if err := s.db.AddUser(userData.FirstName, userData.LastName, userData.Password, userData.Email, ipInfo); err != nil {
	// 	http.Error(respWriter, fmt.Sprintf("failed to add user, error: %s", err.Error()), http.StatusInternalServerError)
	// 	return
	// }

	respWriter.WriteHeader(http.StatusOK)
}

func (s *Server) loginRoute(respWriter http.ResponseWriter, req *http.Request) {
	/*
		 "Content-Type": application/json
		 "Body":
		{
		}
	*/
	io.WriteString(respWriter, "Hello from login route")
}
