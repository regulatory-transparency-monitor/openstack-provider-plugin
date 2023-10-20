package services

import (
	"encoding/json"
	"fmt"
	"log"

	shared "github.com/regulatory-transparency-monitor/commons/models"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/interfaces"
	m "github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/models"
)

// OpenStackPlugin is a struct holding Keystone and Nova service
type OpenStackPlugin struct {
	Keystone interfaces.KeystoneAPI
	Nova     interfaces.NovaAPI
	Cinder   interfaces.CinderAPI
}
type APIContext struct {
	ProjectID string
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
	provider.Cinder = &api.CinderService{
		BaseURL: "volume/v3/",
		Client:  httpClient,
	}

	return nil

}

func (provider *OpenStackPlugin) FetchData() (shared.RawData, error) {
	data := make(shared.RawData)
	// Initialize the context for shared paramters
	ctx, err := provider.initContext()
	if err != nil {
		return data, err
	}

	keystoneData, err := provider.FetchKeystoneData(ctx)
	if err != nil {
		return data, err
	}

	novaData, err := provider.FetchNovaData(ctx)
	if err != nil {
		return data, err
	}

	cinderData, err := provider.FetchCinderData(ctx)
	if err != nil {
		return data, err
	}

	cinderSnapshots, err := provider.FetchCinderSnapshots(ctx)
	if err != nil {
		return data, err
	}

	//PrintAsJSON(cinderSnapshots)
	data["os_project"] = []interface{}{keystoneData}
	data["os_instance"] = novaData
	data["os_volume"] = cinderData
	data["os_snapshot"] = cinderSnapshots

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

func (provider *OpenStackPlugin) FetchCinderData(ctx *APIContext) ([]interface{}, error) {
	var volumeDetailsList []interface{}
	// Use the project ID from the context
	projectID := ctx.ProjectID
	// List of hard-coded volume IDs
	volumeIDs := []string{
		"b5d527e9-0441-4db8-b2f1-b0de5f0ca4ea",
		"8227897a-61d2-408b-a22d-c98705fcb39b",
	}
	for _, volumeID := range volumeIDs {
		volumeDetails, err := provider.Cinder.GetVolumeByID(volumeID, projectID)

		if err != nil {
			fmt.Println(err)
		}
		volumeDetailsList = append(volumeDetailsList, volumeDetails)
	}

	return volumeDetailsList, nil
}
func (provider *OpenStackPlugin) FetchCinderSnapshots(ctx *APIContext) ([]interface{}, error) {
	var snapshotDetailsList []interface{}
	// Use the project ID from the context
	projectID := ctx.ProjectID

	snapshotList, err := provider.Cinder.GetSnapshots(projectID)
	if err != nil {
		return snapshotDetailsList, err
	}

	for _, snapshot := range snapshotList {
		snapshotDetailsList = append(snapshotDetailsList, snapshot)

	}

	return snapshotDetailsList, nil
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
