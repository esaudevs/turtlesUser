package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/esaudevs/turtlesUser/awsgo"
	"github.com/esaudevs/turtlesUser/models"
	"github.com/esaudevs/turtlesUser/bd"


	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	
	awsgo.InitAWS()

	if !ParamsAreValid() {
		fmt.Println("Error in params, 'SecretName' is missing")
		err := errors.New("Error in params, 'SecretName' is missing")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err
}

func ParamsAreValid() bool {
	var params bool
	_, params = os.LookupEnv("SecretName")
	return params
}