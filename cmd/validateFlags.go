package cmd

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/OpsInc/enroller-client/internal/authentication"
	"github.com/OpsInc/enroller-client/internal/cloud"
)

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
		log.Fatalf("Unable to send POST to url: %v because of err: %v", url, err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)
}
