package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create Cognito service client
	svc := cognitoidentityprovider.New(sess)

	result, err := svc.ListUsers(nil)
	if err != nil {
		exitErrorf("Unable to list identities, %v", err)
	}

	fmt.Println("Identities:")

	for _, b := range result.Identities {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.UserPoolID(b.CreationDate))
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
