package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

var apps armappcontainers.ContainerApp

type ContainerApp struct {
    ResourceGroup string
    Name string
    Config armappcontainers.ContainerApp
}

func (c *Client) Deploy(ctx context.Context, ca ContainerApp) (armappcontainers.ContainerAppsClientCreateOrUpdateResponse, error) {
    poller, err := c.AZContainerAppsClient.BeginCreateOrUpdate(
        ctx,
        ca.ResourceGroup, // TODO: replace with real RG name
        ca.Name, // TODO: replace with real container apps name
        ca.Config, // TODO: replae type armappcontainers.CotainerApp
        nil, // Option
    )
    if err != nil {
        panic(err)
    }
    return poller.PollUntilDone(ctx, nil)
}
