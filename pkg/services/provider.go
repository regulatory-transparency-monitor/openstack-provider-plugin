package services

import (
	"fmt"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/client"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/repository"
)

// OpenStackPlugin is a struct that holds the Keystone and Nova services
type OpenStackPlugin struct {
	Keystone repository.KeystoneRepository
	Nova     repository.NovaRepository
}

type ScanResult struct {
    Source string
    Data   []interface{}  // or CommonDataInterface
}

func (provider *OpenStackPlugin) Initialize() error {
	httpClient := client.NewClient("https://api.pub1.infomaniak.cloud/")

	provider.Keystone = &api.KeystoneService{
		BaseURL: "identity/v3/",
		Client:  httpClient,
		// provide url and credentials here
	}
	provider.Nova = &api.NovaService{
		BaseURL: "compute/v2.1/",
		Client:  httpClient,
	}

	return nil

}

// Sample scan function that returns a list of projects and servers
func (provider *OpenStackPlugin) Scan() ([]interface{}, error) {

	// returns a auth token and project ID
	id, err := provider.Keystone.Authenticate()
	if err != nil {
		return nil, err
	}
	projectModel, err := provider.Keystone.GetProjectDetailsByID(id)
	if err != nil {
		return nil, err
	}

	serverListModel, err := provider.Nova.GetServerListByProjectID(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Project Details:", projectModel)
	fmt.Println("Server List:", serverListModel)

	return []interface{}{projectModel, serverListModel}, err
}

