package approleassignment

const (
	//API_VERSION API Version of Azure AD Graph Endpoints
	API_VERSION = "1.6"
	// GRAPH_API_BASE_URL Azure AD Graph Endpoint Base URL
	GRAPH_API_BASE_URL = "https://graph.windows.net/"
	// GRAPH_TOKEN_URL Azure Token Endpoint for 2 legged OAuth
	GRAPH_TOKEN_URL = "https://login.microsoftonline.com/%s/oauth2/token"
	// GRAPH_DEFAULT_SCOPE Default Graph Endpoint Token Scope
	GRAPH_DEFAULT_SCOPE = "https://graph.windows.net/.default"
)
