/*
 * Subspace Product API
 *
 * # Introduction  The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standard HTTP response codes.  The base URL for the API is:  `https://api.subspace.com/`  # Naming Convention  * Version name currently in use is: *v1*   * Example: `https://api.subspace.com/v1`  # Authentication  ## API Tokens  Subspace authenticates your API requests using JWT Bearer tokens. To use any Subspace API, you must pass a Bearer token with each request. If you do not include your Bearer token when making an API request, or use one that is incorrect or disabled, Subspace returns an error.  Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.     Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.   * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.   * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.  ## Getting a JWT Bearer Token  Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.  ## Protecting Your API Tokens    * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.   * The Subspace private token can be rotated from within the Subspace console.  Rotation may take up to 10 minutes for all systems to update state to invalidate the older token and enable the new one.   * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.   * **You may use the Subspace console to acquire an API token.**   * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.   * **Do not embed API keys directly in code.** Instead of directly embedding API keys in your application’s code, put them in environment variables, or within include files that are stored separately from the bulk of your code—outside the source repository of your application. Then, if you share your code, the API keys will not be included in the shared files.   * **Do not store API tokens inside your application’s source control.** If you store API tokens in files, keep the files outside your application’s source control system. This is particularly important if you use a public source code management system such as GitHub.   * **Limit access with restricted tokens.** The Subspace console will allow you to specify the IP addresses or referrer URLs associated with each token, reducing the impact of a compromised API token.   * **Use independent API tokens for different apps.** This limits the scope of each token. If an API token is compromised, you can rotate the impacted token without impacting other API tokens.  # Error Codes  Subspace uses HTTP response codes to indicate the success or failure of an API request.   General HTML status codes:   * 2xx Success.    * 4xx Errors based on information provided in the request.   * 5xx Errors on Subspace servers.    # Security  We provide a valid, signed certificate for our API methods. Be sure your connection library supports HTTPS with the SNI extension.  # REST How-To  Making your first REST API call is easy and can be done from your browser.  You will need:   * Your **secret** token and public client token, both found in the Console.   * The URL for the type of data you would like to request.  First, acquire a JWT Bearer Token.  Command line example:          curl --request POST \\          --url \"https://id.subspace.com/oauth/token\" \\          --header 'content-type: application/json' \\          --data '{ \"client_id\": \"YOURCLIENTID\", \"client_secret\": \"YOURCLIENTSECRET\", \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }'  REST calls are made up of:   * Base url: Example: `https://api.subspace.com`   * Version: Example: `v1`   * The API Endpoint and any parameters: `accelerator/acc_NDA3MUI5QzUtOTY4MC00Nz` where `acc_NDA3MUI5QzUtOTY4MC00Nz` is a valid accelerator ID   * Accelerator ids are always of the format `acc_NDA3MUI5QzUtOTY4MC00Nz`, with a \"acc_\" prefix followed by 22 characters.   * Token header: All REST requests require a valid JWT Bearer token which should be added as an “Authorization” header to the request:              `Authorization: Bearer YOUR_TOKEN_HERE`    ## Authorization header example  If your API token was “my_api_token”, you would add...      Authorization: Bearer my_api_token      ...to the header.  ## Command line examples  To list your current open packet_accelerators using the token “my_api_token”:      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator      Alternately, to get the details of a specific accelerator whose id is 'abcd-ef01-2345':      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator/abcd-ef01-2345  # API Versioning  Subspace will release new versions when we make backwards-incompatible changes to the API. We will give advance notice before releasing a new version or retiring an old version.  Backwards compatible changes:   * Adding new response attributes   * Adding new endpoints   * Adding new methods to an existing endpoint   * Adding new query string parameters   * Adding new path parameters   * Adding new webhook events   * Adding new streaming endpoints   * Changing the order of existing response attributes    Versions are added to the base url, for example:   * `https://api.subspace.com/v1`  Current Version is **v1:** `https://api.subspace.com/v1` 
 *
 * API version: 1.0
 * Contact: sales@subspace.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subspace_openapi_client

import (
	"encoding/json"
)

// V1GlobalTurnResponse struct for V1GlobalTurnResponse
type V1GlobalTurnResponse struct {
	IceServers *[]V1GlobalTurnServer `json:"ice_servers,omitempty"`
	Ttl *int64 `json:"ttl,omitempty"`
}

// NewV1GlobalTurnResponse instantiates a new V1GlobalTurnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GlobalTurnResponse() *V1GlobalTurnResponse {
	this := V1GlobalTurnResponse{}
	return &this
}

// NewV1GlobalTurnResponseWithDefaults instantiates a new V1GlobalTurnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GlobalTurnResponseWithDefaults() *V1GlobalTurnResponse {
	this := V1GlobalTurnResponse{}
	return &this
}

// GetIceServers returns the IceServers field value if set, zero value otherwise.
func (o *V1GlobalTurnResponse) GetIceServers() []V1GlobalTurnServer {
	if o == nil || o.IceServers == nil {
		var ret []V1GlobalTurnServer
		return ret
	}
	return *o.IceServers
}

// GetIceServersOk returns a tuple with the IceServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalTurnResponse) GetIceServersOk() (*[]V1GlobalTurnServer, bool) {
	if o == nil || o.IceServers == nil {
		return nil, false
	}
	return o.IceServers, true
}

// HasIceServers returns a boolean if a field has been set.
func (o *V1GlobalTurnResponse) HasIceServers() bool {
	if o != nil && o.IceServers != nil {
		return true
	}

	return false
}

// SetIceServers gets a reference to the given []V1GlobalTurnServer and assigns it to the IceServers field.
func (o *V1GlobalTurnResponse) SetIceServers(v []V1GlobalTurnServer) {
	o.IceServers = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *V1GlobalTurnResponse) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GlobalTurnResponse) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *V1GlobalTurnResponse) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *V1GlobalTurnResponse) SetTtl(v int64) {
	o.Ttl = &v
}

func (o V1GlobalTurnResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IceServers != nil {
		toSerialize["ice_servers"] = o.IceServers
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableV1GlobalTurnResponse struct {
	value *V1GlobalTurnResponse
	isSet bool
}

func (v NullableV1GlobalTurnResponse) Get() *V1GlobalTurnResponse {
	return v.value
}

func (v *NullableV1GlobalTurnResponse) Set(val *V1GlobalTurnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GlobalTurnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GlobalTurnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GlobalTurnResponse(val *V1GlobalTurnResponse) *NullableV1GlobalTurnResponse {
	return &NullableV1GlobalTurnResponse{value: val, isSet: true}
}

func (v NullableV1GlobalTurnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GlobalTurnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


