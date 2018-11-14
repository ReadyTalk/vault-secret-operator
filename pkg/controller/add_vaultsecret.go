package controller

import (
	"github.com/readytalk/vault-secret-operator/pkg/controller/vaultsecret"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, vaultsecret.Add)
}
