package cloud

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// AWSConnectionExternal works for instances where you do not need
// an AWS authentication but simply need the configuration.
// For example Cognito InitiateAuth function will not require any AWS auth.
//
//nolint:gocritic
func AWSConnectionExternal(AWSRegion string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("AWS LoadDefaultConfig failed because of error: ", err)
	}

	// Set the Region since we have not authenticated to AWS
	cfg.Region = AWSRegion

	return cfg
}
