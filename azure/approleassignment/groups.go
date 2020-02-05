package approleassignment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetAppRoleAssignmentOnGroup gets AppRoleAssignments for a group Id and App Roles Assignment Object Id.
// Returns App Role Assignment Struct on Successful Response or Error incase something went wrong
func (c *AzureClient) GetAppRoleAssignmentOnGroup(groupID, appRoleAssignmentID string) (*AppRolesAssignment, error) {
	url := GRAPH_API_BASE_URL + c.tenantID + "/groups/" + groupID + "/appRoleAssignments/" + appRoleAssignmentID
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("api-version", API_VERSION)
	req.URL.RawQuery = q.Encode()

	appRoleAssignment := new(AppRolesAssignment)

	resp, err := c.Do(context.Background(), req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if json.Unmarshal(body, &appRoleAssignment) != nil {
		return nil, err
	}

	return appRoleAssignment, err
}

// GetAppRoleAssignmentsForGroup gets AppRoleAssignments for a group Id.
// Returns Group Struct on Successful Response or Error incase something went wrong
func (c *AzureClient) GetAppRoleAssignmentsForGroup(groupID string) (*Group, error) {
	url := GRAPH_API_BASE_URL + c.tenantID + "/groups/" + groupID + "/appRoleAssignments"
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("api-version", API_VERSION)
	req.URL.RawQuery = q.Encode()

	group := new(Group)

	resp, err := c.Do(context.Background(), req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if json.Unmarshal(body, &group) != nil {
		return nil, err
	}

	return group, err
}

// AddAppRoleAssignmentToGroup adds AppRoleAssignments.
// Needs AssignAppRoleRequest Struct which needs GroupId, AppRoleId and ResourceId
// Returns AppRoleAssignment Struct on Successful Response or Error incase something went wrong
func (c *AzureClient) AddAppRoleAssignmentToGroup(assignAppRoleRequest *AssignAppRoleRequest) (*AppRolesAssignment, error) {
	url := fmt.Sprint(GRAPH_API_BASE_URL, c.tenantID, "/groups/", *assignAppRoleRequest.GroupID, "/appRoleAssignments/")

	reqbody := map[string]interface{}{
		"id":          *assignAppRoleRequest.AppRoleID,
		"resourceId":  *assignAppRoleRequest.ResourceID, //Service Principal Id
		"principalId": *assignAppRoleRequest.GroupID,
	}

	req, err := c.NewRequest("POST", url, reqbody)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("api-version", API_VERSION)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(context.Background(), req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	appRoleAssignment := new(AppRolesAssignment)

	if json.Unmarshal(body, &appRoleAssignment) != nil {
		return nil, err
	}
	return appRoleAssignment, err
}

// RemoveAppRoleAssignmentFromGroup removes AppRoleAssignments.
// Needs AppRolesAssignment Struct of Group which needs to be deleted
// Returns string on Successful Response or Error incase something went wrong
func (c *AzureClient) RemoveAppRoleAssignmentFromGroup(appRoleAssignment *AppRolesAssignment) (string, error) {
	url := fmt.Sprint(GRAPH_API_BASE_URL, c.tenantID, "/groups/", *appRoleAssignment.PrincipalID, "/appRoleAssignments/", *appRoleAssignment.ObjectID)

	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("api-version", API_VERSION)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(context.Background(), req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
