package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"chalmers/tkey-group22/client/internal/tkey"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"log"
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
func Register(appurl string, username string, label string) error {
	pubkey, err := tkey.GetTkeyPubKey()
	if err != nil {
		return err
	}

	regurl := appurl + "/api/register"
	err = sendRequest(regurl, pubkey, username, label)
	if err != nil {
		return err
	}

	return nil
}

// sendRequest sends a registration request to the specified application URL with the provided public key, username, and label
// It returns an error if the request fails or if the server responds with a status code indicating an error
//
// Parameters:
// - appurl: The URL of the application to which the registration request is sent
// - pubkey: The public key of the user being registered
// - username: The username of the user being registered
// - label: The label for the public key
//
// Returns:
// - An error if the request fails or if the server responds with an error status code
func sendRequest(appurl string, pubkey ed25519.PublicKey, username string, label string) error {
	c := &http.Client{}

	data := RegisterRequest{Username: username, Pubkey: []byte(pubkey), Label: label}
	reqBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Post(appurl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		fmt.Printf("User '%s' has been successfully created!\n", username)
		return nil
	case http.StatusConflict:
		return fmt.Errorf("user '%s' already exists", username)
	case http.StatusBadRequest:
		return fmt.Errorf("invalid request body for user '%s'", username)
	case http.StatusInternalServerError:
		return fmt.Errorf("unable to save user data for user '%s'", username)
	default:
		return fmt.Errorf("unexpected error: %s", res.Status)
	}
}
