package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var (
	auth        string
	authOptions string
)

type ValidateBody struct {
	Path     string `json:"path"`
	Repo     string `json:"repo"`
	Branch   string `json:"branch"`
	Org      string `json:"org"`
	PrNumber string `json:"prNumber"`
}

//nolint:gochecknoglobals
var validateCmd = &cobra.Command{
	Use:   "validate <URL>",
	Short: "Sends validation request.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenID := authenticationMethod(auth)
		body := fetchEnvToJSON()
		postURL(args[0], tokenID, body)
	},
}

//nolint:gochecknoinits
func init() {
	authOptions = "List of supported authentication methods values: [cognito]"
	validateCmd.PersistentFlags().StringVarP(&auth, "auth", "a", "cognito", authOptions)
}

// fetchEnvToJson will fetch all the following requireds environment variables:
// [FILE_PATH, REPO, BRANCH, ORG, PR_NUMBER].
func fetchEnvToJSON() []byte {
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

	prNumber, ok := os.LookupEnv("PR_NUMBER")
	if !ok {
		errMsg = append(errMsg, "PR_NUMBER")
	}

	if len(errMsg) > 0 {
		log.Printf("The following Environment Variables need to be set: %v", errMsg)
		os.Exit(1)
	}

	jsonBody := &ValidateBody{
		Path:     filePath,
		Repo:     repo,
		Branch:   branch,
		Org:      org,
		PrNumber: prNumber,
	}

	jsonRequest, err := json.Marshal(jsonBody)
	if err != nil {
		log.Fatalf("Unable to Marshal the JSON body with err: %v", err)
	}

	return jsonRequest
}
