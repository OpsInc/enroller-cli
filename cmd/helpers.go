package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/OpsInc/enroller-client/internal/authentication"
	"github.com/OpsInc/enroller-client/internal/cloud"
)

//nolint:gochecknoglobals
var (
	auth        string
	authOptions string
	gitKind     string
)

type Data struct {
	Git      string `json:"git"`
	Path     string `json:"path"`
	Repo     string `json:"repo"`
	Branch   string `json:"branch"`
	Org      string `json:"org"`
	PrNumber string `json:"prNumber,omitempty"`
}

// postURL sends a JSON POST request.
// It uses a Bearer auth Header.
func postURL(url string, tokenID string, body []byte) {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Unable to configure POST to url: %v because of err: %v", url, err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Bearer", tokenID)

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		log.Printf("Unable to send POST to url: %v because of err: %v", url, err)
	}
	defer resp.Body.Close()

	readBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading response body failed with error: %v", err)
	}

	os.Stdout.Write([]byte(string(readBody) + "\n"))
}

// Flag: auth
// authenticationMethod defines the authentication method
// and fetchs the auth token.
//
//nolint:revive
func authenticationMethod(auth string) string {
	var errMsg []string
	errMsg = errMsg[:0]

	if strings.ToLower(auth) == "cognito" {
		return cognitoMethod(errMsg)
	} else {
		log.Printf("Unknown value \"%v\" for flag --auth|-a.\n%v", auth, authOptions)
		os.Exit(1)
	}

	return ""
}

func cognitoMethod(errMsg []string) string {
	username, ok := os.LookupEnv("COGNITO_USER")
	if !ok {
		errMsg = append(errMsg, "COGNITO_USER")
	}

	password, ok := os.LookupEnv("COGNITO_PASSWORD")
	if !ok {
		errMsg = append(errMsg, "COGNITO_PASSWORD")
	}

	appClientID, ok := os.LookupEnv("COGNITO_CLIENT_ID")
	if !ok {
		errMsg = append(errMsg, "COGNITO_CLIENT_ID")
	}

	if len(errMsg) > 0 {
		log.Printf("The following Environment Variables need to be set: %v", errMsg)
		os.Exit(1)
	}

	awsCfg := cloud.AWSConnectionExternal("ca-central-1")
	tokenID := authentication.CognitoUserSignin(awsCfg, username, password, appClientID)

	// os.Stdout.Write([]byte(tokenID))

	return tokenID
}
