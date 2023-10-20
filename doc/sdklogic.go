package services

/*

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/pagination"
	shared "github.com/regulatory-transparency-monitor/commons/models"
)

// OpenStackPlugin is a struct holding Keystone and Nova service
type OpenStackPlugin struct {
	Provider *gophercloud.ProviderClient
	//Nova     interfaces.NovaAPI
	//Cinder   interfaces.CinderAPI
}

// TODO pass base url and credentials as parameters from the graph builder service
func (plugin *OpenStackPlugin) Initialize() error {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint:            "https://api.pub1.infomaniak.cloud/identity/v3/",
		ApplicationCredentialID:     "e83a436ec77348e78c14d39c4d06fead",
		ApplicationCredentialSecret: "NF5hNRTfAUew9DAMb-X8bMwLCqViFsubTWz45XnqzECrTDx3npz1D2wQrbue48k1uObKbkPFdQEV4IbNX31btA",
	}

	// Authenticate
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Store the provider client in the OpenStackPlugin struct
	plugin.Provider = provider

	return nil

}

func (plugin *OpenStackPlugin) FetchData() (shared.RawData, error) {
	data := make(shared.RawData)
	// Get the identity client for Keystone
	identityClient, err := openstack.NewIdentityV3(plugin.Provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Failed to get identity client: %v", err)
	}

	// Extract project details
	projectID := "c6b10513ac5542fd844a1f6c7c4bca80"
	log.Printf("Project ID: %s", projectID)
	project, err := projects.Get(identityClient, projectID).Extract()
	if err != nil {
		log.Fatalf("Failed to get project details: %v", err)
	}
	data["project"] = []interface{}{project}

	// Get the compute client for Nova
	computeClient, err := openstack.NewComputeV2(plugin.Provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Failed to get compute client: %v", err)
	}

	// Extract server details for the project
	var serverList []interface{}
	serverPager := servers.List(computeClient, servers.ListOpts{})
	err = serverPager.EachPage(func(page pagination.Page) (bool, error) {
		sList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		for _, s := range sList {
			serverList = append(serverList, s)
		}
		return true, nil
	})
	if err != nil {
		log.Fatalf("Failed to list servers: %v", err)
	}
	data["servers"] = serverList

	// Initialize the Block Storage v3 client
	blockStorageClient, err := openstack.NewBlockStorageV3(plugin.Provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Failed to get block storage client: %v", err)
	}

	// Fetch the list of volumes
	var allVolumes []volumes.Volume
	pager := volumes.List(blockStorageClient, volumes.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		volumeList, err := volumes.ExtractVolumes(page)
		if err != nil {
			return false, err
		}
		allVolumes = append(allVolumes, volumeList...)
		return true, nil
	})
	if err != nil {
		log.Fatalf("Failed to list volumes: %v", err)
	}
	data["volumes"] = allVolumes

	// You can add logic to extract volumes attached to each server similarly using the Cinder client
	PrintAsJSON(data)
	return data, nil
}

func (provider *OpenStackPlugin) FetchKeystoneData(ctx *APIContext) (*m.ProjectDetails, error) {
	projectID, err := provider.Keystone.Authenticate()
	if err != nil {
		fmt.Println(err)
	}

	projectModel, err := provider.Keystone.GetProjectDetailsByID(projectID)
	if err != nil {
		fmt.Println(err)
	}

	ctx.ProjectID = projectID

	return &projectModel, nil
}

func (provider *OpenStackPlugin) FetchNovaData(ctx *APIContext) ([]interface{}, error) {
	var serverDetailsList []interface{}

	projectID := ctx.ProjectID

	serverList, err := provider.Nova.GetServerListByProjectID(projectID)
	if err != nil {
		return serverDetailsList, err
	}

	for _, server := range serverList.Servers {
		serverDetails, err := provider.Nova.GetServerByID(server.ID)
		if err != nil {
			return serverDetailsList, err
		}
		serverDetailsList = append(serverDetailsList, serverDetails)
	}

	return serverDetailsList, nil
}

func (provider *OpenStackPlugin) FetchCinderData(ctx *APIContext) (*m.Volume, error) {

	// Use the project ID from the context
	projectID := ctx.ProjectID
	volumeModel, err := provider.Cinder.GetVolumeByID(projectID)
	if err != nil {
		fmt.Println(err)
	}
	//PrintAsJSON(volumeModel)

	return &volumeModel, nil
}

// Scan returns fetched data from OpenStack API
func (provider *OpenStackPlugin) OLDScan() (models.CombinedResources, error) {
	var resources models.CombinedResources

	resources.Source = "openstack"
	// Initialize the context for shared paramters
	ctx, err := provider.initContext()
	if err != nil {
		return resources, err
	}

	keystoneData, err := provider.FetchKeystoneData(ctx)
	if err != nil {
		return resources, err
	}
	//PrintAsJSON(keystoneData)

	resources.Data = append(resources.Data, *keystoneData)

	novaData, err := provider.FetchNovaData(ctx)
	if err != nil {
		return resources, err
	}
	// PrintAsJSON(novaData)
	resources.Data = append(resources.Data, *novaData)

	cinderData, err := provider.FetchCinderData(ctx)
	if err != nil {
		return resources, err
	}
	resources.Data = append(resources.Data, *cinderData)
	//PrintAsJSON(cinderData)
	return resources, nil
}

func (provider *OpenStackPlugin) OLDFetchKeystoneData(ctx *APIContext) (*models.ServiceData, error) {
	projectID, err := provider.Keystone.Authenticate()
	if err != nil {
		fmt.Println(err)
	}

	projectModel, err := provider.Keystone.GetProjectDetailsByID(projectID)
	if err != nil {
		fmt.Println(err)
	}

	ctx.ProjectID = projectID

	return &models.ServiceData{
		ServiceSource: "identity",
		Data:          []interface{}{projectModel},
	}, nil
}

func (provider *OpenStackPlugin) OLDFetchNovaData(ctx *APIContext) (*models.ServiceData, error) {
	var resources models.ServiceData
	resour

	n &resources, nil
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
*/
