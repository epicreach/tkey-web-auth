package auth

import (
	"bytes"
	. "chalmers/tkey-group22/internal/structs"
	"chalmers/tkey-group22/internal/tkey"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const regurl = "http://localhost:8080/api/register"

func Register() error {

	username := GetUsername()

	pubkey, err := tkey.GetTkeyPubKey()
	if err != nil {
		return err
	}

	sendRequest(pubkey, username)

	return nil
}

func sendRequest(pubkey ed25519.PublicKey, username string) {
	c := &http.Client{}

	data := RegisterRequest{Username: username, Pubkey: []byte(pubkey)}

	reqBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Post(regurl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		fmt.Printf("User '%s' has been successfully created!\n", username)
	case http.StatusConflict:
		fmt.Printf("User '%s' already exists!\n", username)
	case http.StatusBadRequest:
		fmt.Printf("Invalid request body for user '%s'\n", username)
	case http.StatusInternalServerError:
		fmt.Printf("Unable to save user data for user '%s'\n", username)
	default:
		fmt.Printf("Unexpected error: %s\n", res.Status)
	}

}
