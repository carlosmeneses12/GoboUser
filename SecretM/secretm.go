package SecretM

import (
	"enencoding/json"
	"fmt"

	awsgo "github.com/GoboUser/awsGo"
	"github.com/GoboUser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	secret := models.SecretRDSJson{}
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(awsgo.Ctx, input)
	if err != nil {
		fmt.Println("Error al obtener el secreto")
		return secret, err
	}

	err = json.Unmarshal(result.SecretBinary, &secret)
	if err != nil {
		fmt.Println("Error al deserializar el secreto")
		return secret, err
	}

	return secret, nil
}
