package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var dispatchCmd = &cobra.Command{
	Use:   "dispatch <URL>",
	Short: "Sends dispatch request.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenID := authenticationMethod(auth)
		body := fetchEnvToDispatch()
		postURL(args[0], tokenID, body)
	},
}

//nolint:gochecknoinits
func init() {
	authOptions = "List of supported authentication methods values: [cognito]"
	dispatchCmd.PersistentFlags().StringVarP(&auth, "auth", "a", "cognito", authOptions)
}

// Fetches environment variables to feed to dispatcher
func fetchEnvToDispatch() []byte {
	var errMsg []string
	errMsg = errMsg[:0]

	filePath, ok := os.LookupEnv("FILE_PATH")
	if !ok {
		errMsg = append(errMsg, "FILE_PATH")
	}

	repo, ok := os.LookupEnv("REPO")
	if !ok {
		errMsg = append(errMsg, "REPO")
	}

	branch, ok := os.LookupEnv("BRANCH")
	if !ok {
		errMsg = append(errMsg, "BRANCH")
	}

	org, ok := os.LookupEnv("ORG")
	if !ok {
		errMsg = append(errMsg, "ORG")
	}

	if len(errMsg) > 0 {
		log.Printf("The following Environment Variables need to be set: %v", errMsg)
		os.Exit(1)
	}

	jsonBody := &Body{
		Path:   filePath,
		Repo:   repo,
		Branch: branch,
		Org:    org,
	}

	jsonRequest, err := json.Marshal(jsonBody)
	if err != nil {
		log.Fatalf("Unable to Marshal the JSON body with err: %v", err)
	}

	return jsonRequest
}
