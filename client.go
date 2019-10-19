package togotime

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NewAPIClient Constructor for the APIClient object
func NewAPIClient(apiToken string) (*APIClient, error) {
	apiClient := APIClient{
		APIToken: apiToken,
		http: HTTPHelper{
			APIToken: apiToken,
		},
	}
	// Gets /me endpoint so we can fill the object with data.
	userData, err := apiClient.getUserData()
	if err != nil {
		return &apiClient, errors.New(err.Error())
	}
	apiClient.SelfID = userData.Data.ID
	apiClient.Workspaces = userData.Data.Workspaces
	for i := 0; i < len(apiClient.Workspaces); i++ {
		workspace := &apiClient.Workspaces[i] // Pointer so we can just return the Client object
		clients, clientErr := apiClient.getWorkspaceClients(workspace.ID)
		projects, projectErr := apiClient.getWorkspaceProjects(workspace.ID)
		if clientErr != nil {
			return &apiClient, errors.New(clientErr.Error())
		} else if projectErr != nil {
			return &apiClient, errors.New(projectErr.Error())
		}
		workspace.Clients = clients
		workspace.Projects = projects
	}
	return &apiClient, nil
}

// APIClient is a class to allow high level interation with the toggl api
type APIClient struct {
	APIToken   string
	http       HTTPHelper
	SelfID     int64
	Workspaces []Workspace
}

// getUserData gets the /me endpoint for the user
func (c APIClient) getUserData() (MeEndpoint, error) {
	endpoint := baseURL + "me"
	raw, err := c.http.GetRawJSON(endpoint)
	var userData MeEndpoint
	if err != nil {
		return userData, errors.New(err.Error())
	}
	json.Unmarshal(raw, &userData)
	return userData, nil
}

// getWorkspaceClients gets all clients for given workspace and returns them in a slice
func (c APIClient) getWorkspaceClients(workspaceID int64) ([]Client, error) {
	endpoint := baseURL + fmt.Sprintf("workspaces/%d/clients", workspaceID)
	raw, err := c.http.GetRawJSON(endpoint)
	var clients []Client
	if err != nil {
		return clients, errors.New(err.Error())
	}
	json.Unmarshal(raw, &clients)
	return clients, nil
}

// getWorkspaceProjects gets all clients for given workspace and returns them in a slice
func (c APIClient) getWorkspaceProjects(workspaceID int64) ([]Project, error) {
	endpoint := baseURL + fmt.Sprintf("workspaces/%d/projects", workspaceID)
	raw, err := c.http.GetRawJSON(endpoint)
	var projects []Project
	if err != nil {
		return projects, errors.New(err.Error())
	}
	json.Unmarshal(raw, &projects)
	return projects, nil
}
