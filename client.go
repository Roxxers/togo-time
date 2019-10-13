package togotime

import (
	"encoding/json"
	"fmt"
)

// NewAPIClient Constructor for the APIClient object
func NewAPIClient(apiToken string) *APIClient {
	apiClient := APIClient{
		APIToken: apiToken,
		http: HTTPHelper{
			APIToken: apiToken,
		},
	}
	// Gets /me endpoint so we can fill the object with data.
	userData := apiClient.getUserData()
	apiClient.SelfID = userData.Data.ID
	apiClient.Workspaces = userData.Data.Workspaces
	for i := 0; i < len(apiClient.Workspaces); i++ {
		workspace := &apiClient.Workspaces[i] // Pointer so we can just return the Client object
		workspace.Clients = apiClient.getWorkspaceClients(workspace.ID)
		workspace.Projects = apiClient.getWorkspaceProjects(workspace.ID)
	}
	return &apiClient
}

// APIClient is a class to allow high level interation with the toggl api
type APIClient struct {
	APIToken   string
	http       HTTPHelper
	SelfID     int64
	Workspaces []Workspace
}

// getUserData gets the /me endpoint for the user
func (c APIClient) getUserData() MeEndpoint {
	endpoint := baseURL + "me"
	raw := c.http.GetRawJSON(endpoint)
	var userData MeEndpoint
	json.Unmarshal(raw, &userData)
	return userData
}

// getWorkspaceClients gets all clients for given workspace and returns them in a slice
func (c APIClient) getWorkspaceClients(workspaceID int64) []Client {
	endpoint := baseURL + fmt.Sprintf("workspaces/%d/clients", workspaceID)
	raw := c.http.GetRawJSON(endpoint)
	var clients []Client
	json.Unmarshal(raw, &clients)
	return clients
}

// getWorkspaceProjects gets all clients for given workspace and returns them in a slice
func (c APIClient) getWorkspaceProjects(workspaceID int64) []Project {
	endpoint := baseURL + fmt.Sprintf("workspaces/%d/projects", workspaceID)
	raw := c.http.GetRawJSON(endpoint)
	var projects []Project
	json.Unmarshal(raw, &projects)
	return projects
}
