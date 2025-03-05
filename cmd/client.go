package main

import (
	"chalmers/tkey-group22/internal/auth"
	"chalmers/tkey-group22/internal/util"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Define a flag to choose between cmd-client and web-client
	// Run web-client by default
	mode := flag.String("mode", "web", "Choose the mode to run: cmd or web")
	flag.Parse()

	// Start the appropriate client based on the flag value
	switch *mode {
	case "cmd":
		fmt.Println("Starting command-line client...")
		startCmdClient()
	default:
		fmt.Println("Starting web client...")
		startWebClient()
	}
}

func startCmdClient() {

	// Gets mode from user inputs and runs selected mode. Loops until program is told to exit.
	for {
		mode := util.SelectMode()

		switch mode {
		case 1:
			// Perform register
			util.CallRegister()
		case 2:
			// Perform login
			util.CallLogin()
		case 3:
			// Stop program
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func startWebClient() {
	http.Handle("/api/register", enableCors(http.HandlerFunc(registerHandler)))
	http.Handle("/api/login", enableCors(http.HandlerFunc(loginHandler)))

	fmt.Println("Client running on http://localhost:6060")
	http.ListenAndServe(":6060", nil)
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:8080" || origin == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//origin := r.Header.Get("Origin")
	//origin = replaceOriginPort(origin)

	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	username := requestBody["username"]
	err := auth.Login("http://localhost:8080", username)
	if err != nil {
		http.Error(w, "Failed to log in", http.StatusBadRequest)
		return
	}

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	origin = replaceOriginPort(origin)

	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := requestBody["username"]
	err := auth.Register(origin, username)
	if err != nil {
		http.Error(w, "Failed to register", http.StatusBadRequest)
	}
}

// TODO: Auto-detect which port application is running on
// Change port of request to 8080
func replaceOriginPort(origin string) string {
	parts := strings.Split(origin, ":")
	if len(parts) > 1 {
		parts[len(parts)-1] = "8080"
		origin = strings.Join(parts, ":")
	}
	return origin
}
