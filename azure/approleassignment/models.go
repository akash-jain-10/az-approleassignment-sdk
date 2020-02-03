package approleassignment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//AzureClient a Client that manages communication with the Azure Graph API
type AzureClient struct {
	// Organization Id to be used for subsequent API Calls
	tenantID string
	// Base URL for API requests. Defaults to the Graph API Base URL
	baseURL *url.URL
	// HTTP client used to communicate with the API
	httpClient *http.Client
}

// Error Response Object from Graph API
type odataResponse struct {
	Resp *http.Response
	// Odata Error Body
	OdataErr *odataError `json:"odata.error"`
}

// Odata Error Message
type odataErrorMessage struct {
	// Language of Error
	Lang *string `json:"lang"`
	// Error Message
	Value *string `json:"value"`
}

// OdataError error information.
type odataError struct {
	// Code - Error code.
	Code *string `json:"code"`
	// ErrorMessage - OData Error Message.
	Message *odataErrorMessage `json:"message"`
	// RequestId - Request Id generated for Error.
	RequestID *string `json:"requestId"`
}

// Format String for Error Response
func (r *odataResponse) Error() string {
	return fmt.Sprintf("%v: %d %v %+v", r.Resp.Request.Method, r.Resp.StatusCode, *r.OdataErr.Code, *r.OdataErr.Message.Value)
}

// ObjectType enumerates the values for object type.
type ObjectType string

const (
	// ObjectTypeDirectoryObject ...
	ObjectTypeDirectoryObject ObjectType = "DirectoryObject"
	// ObjectTypeGroup ...
	ObjectTypeGroup ObjectType = "Group"
	// ObjectTypeServicePrincipal ...
	ObjectTypeServicePrincipal ObjectType = "ServicePrincipal"
	// ObjectTypeAppRoleAssignment ...
	ObjectTypeAppRoleAssignment ObjectType = "AppRoleAssignment"
)

// PossibleObjectTypeValues returns an array of possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{ObjectTypeDirectoryObject, ObjectTypeGroup, ObjectTypeServicePrincipal, ObjectTypeAppRoleAssignment}
}

// AppRolesAssignment - Used to record when a user or group is assigned to an application.
type AppRolesAssignment struct {
	// ObjectType - Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeAppRoleAssignment'
	ObjectType *ObjectType `json:"objectType,omitempty"`
	// OdataType - Microsoft.DirectoryServices.AppRoleAssignment
	OdataType *string `json:"odata.type"`
	// ObjectId - AppRolesAssignment Id for the AppRolesAssignment
	ObjectID *string `json:"objectId,omitempty"`
	// The role id that was assigned to the principal.
	ID *string `json:"id,omitempty"`
	// The display name of the principal that was granted the access.
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	// The unique identifier (id) for the principal being granted the access. Required on create.
	PrincipalID *string `json:"principalId,omitempty"`
	// The type of principal. This can either be "User", "Group" or "ServicePrincipal".
	PrincipalType *string `json:"principalType,omitempty"`
	// The display name of the resource to which the assignment was made.
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// The unique identifier (id) for the target resource (service principal) for which the assignment was made.
	ResourceID *string `json:"resourceId,omitempty"`
	// Other Properties
	AdditionalProperties map[string]interface{} `json:""`
}

// Group active Directory group approlesassignment information.
type Group struct {
	// OdataMetadata - The URL representing edm equivalent.
	OdataMetadata *string `json:"odata.metadata,omitempty"`
	// Value - A collection of AppRolesAssignment.
	Value *[]AppRolesAssignment `json:"value"`
}

// ServicePrincipal approleassignment information.
type ServicePrincipal struct {
	// OdataMetadata - The URL representing edm equivalent.
	OdataMetadata *string `json:"odata.metadata,omitempty"`
	// Value - A collection of AppRolesAssignment.
	Value *[]AppRolesAssignment `json:"value"`
}

// UnmarshalJSON is the custom unmarshaler for AppRoleAssignment struct.
func (a *AppRolesAssignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				var err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "objectId":
			if v != nil {
				var objectID string
				var err = json.Unmarshal(*v, &objectID)
				if err != nil {
					return err
				}
				a.ObjectID = &objectID
			}
		case "odata.type":
			if v != nil {
				var OdataType string
				var err = json.Unmarshal(*v, &OdataType)
				if err != nil {
					return err
				}
				a.OdataType = &OdataType
			}
		case "objectType":
			if v != nil {
				var ObjectType ObjectType
				var err = json.Unmarshal(*v, &ObjectType)
				if err != nil {
					return err
				}
				a.ObjectType = &ObjectType
			}
		case "principalDisplayName":
			if v != nil {
				var PrincipalDisplayName string
				var err = json.Unmarshal(*v, &PrincipalDisplayName)
				if err != nil {
					return err
				}
				a.PrincipalDisplayName = &PrincipalDisplayName
			}
		case "principalId":
			if v != nil {
				var principalID string
				var err = json.Unmarshal(*v, &principalID)
				if err != nil {
					return err
				}
				a.PrincipalID = &principalID
			}
		case "principalType":
			if v != nil {
				var PrincipalType string
				var err = json.Unmarshal(*v, &PrincipalType)
				if err != nil {
					return err
				}
				a.PrincipalType = &PrincipalType
			}
		case "resourceDisplayName":
			if v != nil {
				var ResourceDisplayName string
				var err = json.Unmarshal(*v, &ResourceDisplayName)
				if err != nil {
					return err
				}
				a.ResourceDisplayName = &ResourceDisplayName
			}
		case "resourceId":
			if v != nil {
				var resourceID string
				var err = json.Unmarshal(*v, &resourceID)
				if err != nil {
					return err
				}
				a.ResourceID = &resourceID
			}
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if a.AdditionalProperties == nil {
					a.AdditionalProperties = make(map[string]interface{})
				}
				a.AdditionalProperties[k] = additionalProperties
			}
		}
	}
	return nil
}

// AssignAppRoleRequest Request struct
type AssignAppRoleRequest struct {
	// GroupId - Id of the Group
	GroupID *string
	// AppRoleId - Id of App Role to be assigned to; A zero guid if default AppRole to Assign
	AppRoleID *string
	// ResourceId - Service Principal Id of the Azure App
	ResourceID *string
}
