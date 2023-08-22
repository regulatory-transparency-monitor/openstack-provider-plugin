package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/interfaces"
)

// OpenStackPlugin is a struct holding Keystone and Nova service
type OpenStackPlugin struct {
	Keystone interfaces.KeystoneAPI
	Nova     interfaces.NovaAPI
}
type APIContext struct {
	ProjectID string
}

type CombinedResources struct {
	Source string
	Data   []ServiceData
}
type ServiceData struct {
	ServiceSource string
	Data          []interface{}
}

// TODO pass base url and credentials as parameters from the graph builder service
func (provider *OpenStackPlugin) Initialize() error {
	httpClient := httpwrapper.NewClient("https://api.pub1.infomaniak.cloud/")

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

// Scan returns fetched data from OpenStack API
func (provider *OpenStackPlugin) Scan() (*CombinedResources, error) {
	var resources CombinedResources

	resources.Source = "openstack"
	// Initialize the context for shared paramters
	ctx, err := provider.initContext()
	if err != nil {
		return nil, err
	}

	keystoneData, err := provider.FetchKeystoneData(ctx)
	if err != nil {
		return nil, err
	}
	//PrintAsJSON(keystoneData)

	resources.Data = append(resources.Data, *keystoneData)

	novaData, err := provider.FetchNovaData(ctx)
	if err != nil {
		return nil, err
	}
	// PrintAsJSON(novaData)
	resources.Data = append(resources.Data, *novaData)

	//PrintAsJSON(resources)
	/* jsonData, err := json.Marshal(resources)
		if err != nil {
	    	log.Fatalf("Failed to marshal data: %v", err)
		}
	jsonString := string(jsonData */
	return &resources, nil
}

func PrintAsJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	fmt.Println(string(data))
}
func (provider *OpenStackPlugin) initContext() (*APIContext, error) {
	ctx := &APIContext{}
	return ctx, nil
}

func (provider *OpenStackPlugin) FetchKeystoneData(ctx *APIContext) (*ServiceData, error) {
	projectID, err := provider.Keystone.Authenticate()
	if err != nil {
		fmt.Println(err)
	}

	projectModel, err := provider.Keystone.GetProjectDetailsByID(projectID)
	if err != nil {
		fmt.Println(err)
	}

	ctx.ProjectID = projectID

	return &ServiceData{
		ServiceSource: "identity",
		Data:          []interface{}{projectModel},
	}, nil
}

func (provider *OpenStackPlugin) FetchNovaData(ctx *APIContext) (*ServiceData, error) {
	// Use the project ID from the context
	var resources ServiceData

	resources.ServiceSource = "compute"

	projectID := ctx.ProjectID
	serverList, err := provider.Nova.GetServerListByProjectID(projectID)
	if err != nil {
		fmt.Println(err)
	}

	for _, server := range serverList.Servers {
		serverDetails, err := provider.Nova.GetServerByID(server.ID)
		// TODO append to NovaData
		resources.Data = append(resources.Data, serverDetails)
		if err != nil {
			return nil, err
		}
	}

	return &resources, nil
}
