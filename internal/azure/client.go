package azure

import (
	"errors"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

type Client struct {
	AZContainerAppsClient *armappcontainers.ContainerAppsClient
}

func NewClient() (*Client, error) {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return nil, errors.New("AZURE_SUBSCRIPTION_ID must be set")
	}
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return nil, err
	}
	clientFactory, err := armappcontainers.NewClientFactory(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}
	return &Client{AZContainerAppsClient: clientFactory.NewContainerAppsClient()}, nil
}
