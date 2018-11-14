package vault

import (
  "os"
  log "github.com/Sirupsen/logrus"
  VaultApi "github.com/hashicorp/vault/api"
)


var VaultClient *VaultApi.Client
var Vault *VaultApi.Logical

func init() {

  // Configure new Vault Client
  vaultAddr := os.Getenv("VAULT_ADDR")
  vaultToken := os.Getenv("VAULT_TOKEN")

  conf := &VaultApi.Config{Address: vaultAddr}
  VaultClient, _ = VaultApi.NewClient(conf)
  VaultClient.SetToken(vaultToken)

  // Define a Logical Vault client (to read/write values)
  Vault = VaultClient.Logical()
}

func GetSecret(path string) *VaultApi.Secret {

  // Read the secret from Vault
  log.Info("Fetching secret: ", path)
  secret, err := Vault.Read(path)
  if err != nil {
    log.Fatal("Error fetching secret: ", err.Error())
  }

  // If we got back an empty response, fail
  if secret == nil {
    log.Fatal("Could not find secret ", path)
  }

  return secret
}
