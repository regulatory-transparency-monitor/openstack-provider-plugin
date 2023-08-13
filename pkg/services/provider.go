package services

import (
	"collector-service/internal/repository"
	"collector-service/pkg/logger"
	"context"
)

type Provider struct {
	Keystone repository.KeystoneRepository
	Nova     repository.NovaRepository
}

func (p Provider) InitializeScan(appLogger *logger.APIlogger) error {
	ctx := context.Background()

	// returns a auth token and project ID
	auth, err := p.Keystone.Authenticate(ctx)
	if err != nil {
		appLogger.Infof("Authentication Error: %+v\n", err)
		return err
	}
	
	// returns a list of projects
	projectModel, err := p.Keystone.GetProjects()
	if err != nil {
		appLogger.Fatalf("Error while coverting: %s\n", err)
	}
	
	// returns a list of servers
	serverModel, err := p.Nova.GetServers(auth)
	if err != nil {
		appLogger.Infof("G %+v\n", err)

		return err
	}
	
	appLogger.Infof("Project Details: %+v\n", projectModel)

	appLogger.Infof("Server List: %+v\n", serverModel)

	// send to collector service ?!
	return err
}

/*
func (p Provider) GetProjects() error {
	// Authenticate using Keystone
	err := p.Keystone.Authenticate()
	if err != nil {
		return err
	}
	ProjectDetails, err := p.Keystone.GetProjects()

	fmt.Printf("Project Details: %+v", ProjectDetails)
	// Get projects
	// projects, _ := p.keystone.GetProjects()

	// ... process data ...

	return err
} */

func (p Provider) GetServerInfo() {
	// Authenticate using Keystone

	// Get projects, users, and servers
	/*projects, _ := p.keystone.GetProjects()
	users, _ := p.keystone.GetUsers()
	servers, _ := p.nova.GetServers()*/

	// ... process data ...

	// return servers, nil
}
