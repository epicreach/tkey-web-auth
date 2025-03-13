package auth

import (
	"bytes"
	. "chalmers/tkey-group22/client/internal/structs"
	"chalmers/tkey-group22/client/internal/tkey"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

// Sends a request to the server to login a user
// This requires that the app has the /api/login and /api/verify endpoints
// It returns an error if the login process fails
//
// Parameters:
// - appurl: The URL of the application server
// - username: The username of the user to login
//
// Returns:
// - A cookie pointer
// - An error if the login process fails
func Login(appurl string, username string) (*http.Cookie, error) {

	//Creates a cookie jar to send the cookie across functions programmatically
	jar, _ := cookiejar.New(nil)
	c := &http.Client{Jar: jar}

	// Fetches the generated challenge from the server
	challengeResponse, err := getChallenge(appurl, username)
	if err != nil {
		return nil, err
	}

	var cookie http.Cookie

	// TODO: Implement signature verification
	// if !verifySignature(challengeResponse) {
	// 	return fmt.Errorf("signature verification failed")
	// }

	// Signs the challenge
	signedChallenge, err := signChallenge(username, challengeResponse)
	if err != nil {
		return nil, err
	}

	// TODO: Make more customizable
	endpoint := "/api/verify"

	body, err := json.Marshal(signedChallenge)
	if err != nil {
		return nil, err
	}

	// Sends the signed challenge to the server in the format of a VerifyRequest
	resp, err := c.Post(appurl+endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//Takes the (first) cookie from the response
	cookies := resp.Cookies()
	if len(cookies) > 0 {
		cookie = *cookies[0]
	} else {
		return nil, fmt.Errorf("No cookies received")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Printf("User '%s' has been successfully logged in!\n", username)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("invalid signature")
	case http.StatusNotFound:
		return nil, fmt.Errorf("no active challenge found for the user")
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("unable to read user data")
	default:
		return nil, fmt.Errorf("unexpected error: %s", resp.Status)
	}

	return &cookie, nil
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
func signChallenge(username string, challenge *LoginResponse) (*VerifyRequest, error) {
	fmt.Printf("Touch the TKey to continue...\n")
	sig, err := tkey.Sign([]byte(challenge.Challenge))
	if err != nil {
		return nil, err
	}

	return &VerifyRequest{
		Username:  username,
		Signature: sig,
	}, nil
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
func getChallenge(appurl string, user string) (*LoginResponse, error) {
	endpoint := "/api/login"

	c := &http.Client{}

	body, err := json.Marshal(LoginRequest{Username: user})
	if err != nil {
		return nil, err
	}

	// Get the signature and message from the endpoint
	resp, err := c.Post(appurl+endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// Continue processing
	case http.StatusNotFound:
		return nil, fmt.Errorf("user not found")
	case http.StatusBadRequest:
		return nil, fmt.Errorf("invalid request body or missing username")
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("unable to read user data")
	default:
		return nil, fmt.Errorf("unexpected error: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body")
	}

	// Decode the response body into a LoginResponse struct
	var res *LoginResponse
	err = json.Unmarshal(respBody, &res)

	if err != nil {
		return nil, fmt.Errorf("error decoding challenge response")
	}

	return res, nil
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
