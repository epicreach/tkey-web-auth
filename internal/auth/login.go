package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"chalmers/tkey-group22/client/internal/tkey"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Programatically returns a signed challenge. Expects an appurl to request the
// challenge from and a username to associate with that challenge.
// It returns an error if unable to get challenge or sign the challenge.
// It will also return a message related to the error if applicable.
//
// Parameters:
// - appurl: The URL of the application server
// - username: The username of the user to login
//
// Returns:
// - A signed challenge
// - An username
// - A error message string (if applicable)
// - An error if the login process fails
func GetAndSign(appurl string, username string) (string, []byte, string, error) {

	// Fetches the generated challenge from the server
	challengeResponse, errMsg, err := getChallenge(appurl, username)

	if err != nil {
		fmt.Println("Error getting challenge")
		return "", nil, errMsg, err
	}

	// TODO: Implement signature verification
	// if !verifySignature(challengeResponse) {
	// 	return fmt.Errorf("signature verification failed")
	// }

	// Signs the challenge
	user, signedChallenge, err := signChallenge(username, challengeResponse)
	if err != nil {
		return "", nil, err.Error(), err
	}
	return user, signedChallenge, "", nil
}

// An internal function that signs the challenge using the tkey
//
// Parameters:
// - username: The username of the user to sign the challenge for
// - challenge: The challenge to sign
//
// Returns:
// - A VerifyRequest struct containing the username and signature
// - An error if the signing process fails
func signChallenge(username string, challenge *LoginResponse) (string, []byte, error) {
	fmt.Printf("Touch the TKey to continue...\n")
	sig, err := tkey.Sign([]byte(challenge.Challenge))
	if err != nil {
		return "", nil, err
	}

	return username, sig, nil
}

// An internal function that fetches the challenge from the server
//
// Parameters:
// - appurl: The URL of the application server
// - user: The username of the user to fetch the challenge for
//
// Returns:
// - A LoginResponse struct containing the challenge and signature
// - An error if the request fails
func getChallenge(appurl string, user string) (*LoginResponse, string, error) {
	endpoint := "/api/login"

	c := &http.Client{}

	body, err := json.Marshal(LoginRequest{Username: user})
	if err != nil {
		return nil, "", err
	}

	// Get the signature and message from the endpoint
	resp, err := c.Post(appurl+endpoint, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	respBodyStr := string(respBody)

	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// Continue processing
	case http.StatusNotFound:
		return nil, respBodyStr, fmt.Errorf("user not found")
	case http.StatusBadRequest:
		return nil, respBodyStr, fmt.Errorf("invalid request body or missing username")
	case http.StatusInternalServerError:
		return nil, respBodyStr, fmt.Errorf("unable to read user data")
	default:
		return nil, respBodyStr, fmt.Errorf("unexpected error: %s", resp.Status)
	}

	// Decode the response body into a LoginResponse struct
	var res *LoginResponse
	err = json.Unmarshal(respBody, &res)

	if err != nil {
		return nil, "", fmt.Errorf("error decoding challenge response")
	}

	return res, "", nil
}

// TODO: Uncomment once server verification is implemented
// func verifySignature(res *LoginResponse) bool {
// 	pubkey, err := getPublicKey()
// 	if err != nil {
// 		return false
// 	}
// 	return ed25519.Verify(pubkey, []byte(res.Challenge), []byte(res.Signature))
// }

// func getPublicKey() ([]byte, error) {
// 	baseUrl := "http://localhost:8080"
// 	endpoint := "/api/public"

// 	c := &http.Client{}

// 	resp, err := c.Get(baseUrl + endpoint)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("error in response when getting public key")
// 	}

// 	var pubkey []byte

// 	err = json.NewDecoder(resp.Body).Decode(&pubkey)

// 	return pubkey, err
// }
