package authentication

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func CognitoUserSignin(awsCfg aws.Config, username string, password string, appClientID string) string {
	//nolint:exhaustruct
	// Ignoring unsused fields: AnalyticsMetadata, ClientMetadata, UserContextData
	InitiateAuthInput := &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(appClientID),
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,

		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
	}

	cognitoProvider := cognitoidentityprovider.NewFromConfig(awsCfg)

	out, err := cognitoProvider.InitiateAuth(context.TODO(), InitiateAuthInput)
	if err != nil {
		log.Fatal("Cognito signin failed because of error: ", err)
	}

	// fmt.Println(*out.AuthenticationResult.IdToken)

	return *out.AuthenticationResult.IdToken
}
