package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	awsgo "github.com/GoboUser/awsGo"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.IniciarAws()

	if !ValidoParametros() {
		fmt.Println("No se enviaron los parametros necesarios")
		err := errors.New("No se enviaron los parametros necesarios")
		return event, err
	}

}

func ValidoParametros() bool {
	var traeParametro bool
	_, tratraeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
