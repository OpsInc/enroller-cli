package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var validateCmd = &cobra.Command{
	Use:   "validate <URL>",
	Short: "Sends validation request.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { //nolint:revive
		tokenID := authenticationMethod(auth)
		body := fetchEnvToValidate()
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
func fetchEnvToValidate() []byte {
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

	jsonData := &Data{
		Git:      gitKind,
		Path:     filePath,
		Repo:     repo,
		Branch:   branch,
		Org:      org,
		PrNumber: prNumber,
	}

	jsonRequest, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Unable to Marshal the JSON body with err: %v", err)
	}

	return jsonRequest
}
