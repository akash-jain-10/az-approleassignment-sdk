package approleassignment

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

// GetAppRoleAssignedToForServicePrincipal gets AppRoleAssignments assigned to a service principal Id.
// Returns ServicePrincipal Struct on Successful Response or Error incase something went wrong
func (c *AzureClient) GetAppRoleAssignedToForServicePrincipal(servicePrincipalID string) (*ServicePrincipal, error) {
	url := GRAPH_API_BASE_URL + c.tenantID + "/servicePrincipals/" + servicePrincipalID + "/appRoleAssignedTo"
	req, err := c.NewRequest("GET", url, nil)
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
	servicePrincipal := new(ServicePrincipal)

	if json.Unmarshal(body, &servicePrincipal) != nil {
		return nil, err
	}
	return servicePrincipal, err
}
