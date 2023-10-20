package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/regulatory-transparency-monitor/openstack-provider-plugin/pkg/services"
)

func main() {

	// Create a new OpenStackPlugin instance
	plugin := &services.OpenStackPlugin{}

	// Initialize the plugin (authenticates and sets up the provider)
	err := plugin.Initialize()
	if err != nil {
		log.Fatalf("Error during initialization: %s", err)
	}

	// Fetch data
	fetchResources, err := plugin.FetchData()
	PrintAsJSON(fetchResources)
	if err != nil {
		log.Fatalf("Error while fetching data: %s", err)
	}
}

func PrintAsJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	fmt.Println(string(data))
}
