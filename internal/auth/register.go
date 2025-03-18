package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"chalmers/tkey-group22/client/internal/tkey"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/http"
)

// Register registers a new user with the given username and label at the specified app URL
// This requires that the app has the /api/register endpoint
// It returns an error if the registration process fails
//
// Parameters:
//   - appurl: The base URL of the application where the user will be registered
//   - username: The username of the user to be registered
//   - label: The label for the public key
//
// Returns:
//   - error: An error if the registration process fails, otherwise nil
//   - string: A string containing the error message from the application. Will be empty string if no error occured.
func Register(appurl string, username string, label string) (*http.Response, error) {

	pubkey, err := tkey.GetTkeyPubKey()
	if err != nil {
		return nil, err
	}

	regurl := appurl + "/api/register"
	res, err := sendRequest(regurl, pubkey, username, label)

	if err != nil {
		return res, err
	}

	return res, nil
}

// sendRequest sends a registration request to the specified application URL with the provided public key, username, and label
// It returns an error if the request fails or if the server responds with a status code indicating an error. It also returns a response body if an error occurs.
//
// Parameters:
// - appurl: The URL of the application to which the registration request is sent
// - pubkey: The public key of the user being registered
// - username: The username of the user being registered
// - label: The label for the public key
//
// Returns:
// - An error if the request fails or if the server responds with an error status code
// - A string containing the body of the response in case of error.

func sendRequest(appurl string, pubkey ed25519.PublicKey, username string, label string) (*http.Response, error) {
	c := &http.Client{}

	data := RegisterRequest{Username: username, Pubkey: []byte(pubkey), Label: label}
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res, err := c.Post(appurl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return res, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		fmt.Printf("User '%s' has been successfully created!\n", username)
		return res, nil
	case http.StatusConflict:
		return res, fmt.Errorf("user '%s' already exists", username)
	case http.StatusBadRequest:
		return res, fmt.Errorf("invalid request body for user '%s'", username)
	case http.StatusInternalServerError:
		return res, fmt.Errorf("unable to save user data for user '%s'", username)
	default:
		return res, fmt.Errorf("unexpected error: %s", res.Status)
	}
}
