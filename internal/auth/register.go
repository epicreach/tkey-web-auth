package auth

import (
	"bytes"
	"chalmers/tkey-group22/internal/tkey"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const regurl = "http://localhost:8080/api/register"

func Register() error {

	username := getUsername()

	pubkey, err := tkey.GetTkeyPubKey()
	if err != nil {
		return err
	}

	sendRequest(pubkey, username)

	return nil
}

func getUsername() string {
	var username string
	fmt.Print("Please enter username: ")
	fmt.Scan(&username)
	return username
}

func convertToJSON(data map[string]string) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

func createRequest(jsonData []byte, url string) *http.Request {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func sendRequest(pubkey []byte, username string) {

	pubkeyStr := string(pubkey[:])
	data := map[string]string{"username": username, "pubkey": pubkeyStr}

	jsonData := convertToJSON(data)

	req := createRequest(jsonData, regurl)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200" {
		fmt.Printf("Could not create user! Error: %s", resp.Status)
		log.Fatal()
	} else {
		fmt.Printf("User '%s' has been successfully created!", username)
	}

}
