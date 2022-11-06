package vault

import (
	hvault "github.com/hashicorp/vault/api"
	"os"
)

var (
	MountPath = os.Getenv("VAULT_MOUNT_PATH")
	token     = os.Getenv("VAULT_TOKEN")
	address   = os.Getenv("VAULT_ADDR")
)

func InitClient() (*hvault.Client, error) {
	config := hvault.DefaultConfig()
	// change default address (localhost:8200) to address from environment variables
	config.Address = address
	client, err := hvault.NewClient(config)
	client.SetToken(token)
	return client, err
}
