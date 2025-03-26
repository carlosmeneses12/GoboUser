package db

import (
	"os"

	"github.com/GoboUser/SecretM/secretm"
	"github.com/GoboUser/models"
)

var err error
var SecretModel models.SecretRDSJson

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
