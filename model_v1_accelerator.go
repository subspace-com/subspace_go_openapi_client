/*
Subspace Product API

# Introduction  The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standard HTTP response codes.  The base URL for the API is:  `https://api.subspace.com/`  # Naming Convention  * Version name currently in use is: *v1*   * Example: `https://api.subspace.com/v1`  # Authentication  ## API Tokens  Subspace authenticates your API requests using JWT Bearer tokens. To use any Subspace API, you must pass a Bearer token with each request. If you do not include your Bearer token when making an API request, or use one that is incorrect or disabled, Subspace returns an error.  Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.     Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.   * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.   * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.  ## Getting a JWT Bearer Token  Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.  ## Protecting Your API Tokens    * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.   * The Subspace private token can be rotated from within the Subspace console.   * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.   * **You may use the Subspace console to acquire an API token.**   * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.   * **Do not embed API keys directly in code.** Instead of directly embedding API keys in your application’s code, put them in environment variables, or within include files that are stored separately from the bulk of your code—outside the source repository of your application. Then, if you share your code, the API keys will not be included in the shared files.   * **Do not store API tokens inside your application’s source control.** If you store API tokens in files, keep the files outside your application’s source control system. This is particularly important if you use a public source code management system such as GitHub.   * **Limit access with restricted tokens.** The Subspace console will allow you to specify the IP addresses or referrer URLs associated with each token, reducing the impact of a compromised API token.   * **Use independent API tokens for different apps.** This limits the scope of each token. If an API token is compromised, you can rotate the impacted token without impacting other API tokens.  # Error Codes  Subspace uses HTTP response codes to indicate the success or failure of an API request.   General HTML status codes:   * 2xx Success.    * 4xx Errors based on information provided in the request.   * 5xx Errors on Subspace servers.    # Security  We provide a valid, signed certificate for our API methods. Be sure your connection library supports HTTPS with the SNI extension.  # REST How-To  Making your first REST API call is easy and can be done from your browser.  You will need:   * Your **secret** token and public client token, both found in the Console.   * The URL for the type of data you would like to request.  First, acquire a JWT Bearer Token.  Command line example:          curl --request POST \\          --url \"https://id.subspace.com/oauth/token\" \\          --header 'content-type: application/json' \\          --data '{ \"client_id\": \"YOURCLIENTID\", \"client_secret\": \"YOURCLIENTSECRET\", \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }'  REST calls are made up of:   * Base url: Example: `https://api.subspace.com`   * Version: Example: `v1`   * The API Endpoint and any parameters: `accelerator/acc_NDA3MUI5QzUtOTY4MC00Nz` where `acc_NDA3MUI5QzUtOTY4MC00Nz` is a valid accelerator ID   * Accelerator ids are always of the format `acc_NDA3MUI5QzUtOTY4MC00Nz`, with a \"acc_\" prefix followed by 22 characters.   * Token header: All REST requests require a valid JWT Bearer token which should be added as an “Authorization” header to the request:              `Authorization: Bearer YOUR_TOKEN_HERE`    ## Authorization header example  If your API token was “my_api_token”, you would add...      Authorization: Bearer my_api_token      ...to the header.  ## Command line examples  To list your current open packet_accelerators using the token “my_api_token”:      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator      Alternately, to get the details of a specific accelerator whose id is 'abcd-ef01-2345':      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator/abcd-ef01-2345  # API Versioning  Subspace will release new versions when we make backwards-incompatible changes to the API. We will give advance notice before releasing a new version or retiring an old version.  Backwards compatible changes:   * Adding new response attributes   * Adding new endpoints   * Adding new methods to an existing endpoint   * Adding new query string parameters   * Adding new path parameters   * Adding new webhook events   * Adding new streaming endpoints   * Changing the order of existing response attributes    Versions are added to the base url, for example:   * `https://api.subspace.com/v1`  Current Version is **v1:** `https://api.subspace.com/v1` 

API version: 1.0
Contact: sales@subspace.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subspace_openapi_client

import (
	"encoding/json"
)

// V1Accelerator struct for V1Accelerator
type V1Accelerator struct {
	DestinationIp *string `json:"destination_ip,omitempty"`
	DestinationPort *int64 `json:"destination_port,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SubspaceIpv4 *string `json:"subspace_ipv4,omitempty"`
	SubspacePort *int64 `json:"subspace_port,omitempty"`
}

// NewV1Accelerator instantiates a new V1Accelerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Accelerator() *V1Accelerator {
	this := V1Accelerator{}
	return &this
}

// NewV1AcceleratorWithDefaults instantiates a new V1Accelerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AcceleratorWithDefaults() *V1Accelerator {
	this := V1Accelerator{}
	return &this
}

// GetDestinationIp returns the DestinationIp field value if set, zero value otherwise.
func (o *V1Accelerator) GetDestinationIp() string {
	if o == nil || o.DestinationIp == nil {
		var ret string
		return ret
	}
	return *o.DestinationIp
}

// GetDestinationIpOk returns a tuple with the DestinationIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetDestinationIpOk() (*string, bool) {
	if o == nil || o.DestinationIp == nil {
		return nil, false
	}
	return o.DestinationIp, true
}

// HasDestinationIp returns a boolean if a field has been set.
func (o *V1Accelerator) HasDestinationIp() bool {
	if o != nil && o.DestinationIp != nil {
		return true
	}

	return false
}

// SetDestinationIp gets a reference to the given string and assigns it to the DestinationIp field.
func (o *V1Accelerator) SetDestinationIp(v string) {
	o.DestinationIp = &v
}

// GetDestinationPort returns the DestinationPort field value if set, zero value otherwise.
func (o *V1Accelerator) GetDestinationPort() int64 {
	if o == nil || o.DestinationPort == nil {
		var ret int64
		return ret
	}
	return *o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetDestinationPortOk() (*int64, bool) {
	if o == nil || o.DestinationPort == nil {
		return nil, false
	}
	return o.DestinationPort, true
}

// HasDestinationPort returns a boolean if a field has been set.
func (o *V1Accelerator) HasDestinationPort() bool {
	if o != nil && o.DestinationPort != nil {
		return true
	}

	return false
}

// SetDestinationPort gets a reference to the given int64 and assigns it to the DestinationPort field.
func (o *V1Accelerator) SetDestinationPort(v int64) {
	o.DestinationPort = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1Accelerator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1Accelerator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1Accelerator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1Accelerator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1Accelerator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1Accelerator) SetName(v string) {
	o.Name = &v
}

// GetSubspaceIpv4 returns the SubspaceIpv4 field value if set, zero value otherwise.
func (o *V1Accelerator) GetSubspaceIpv4() string {
	if o == nil || o.SubspaceIpv4 == nil {
		var ret string
		return ret
	}
	return *o.SubspaceIpv4
}

// GetSubspaceIpv4Ok returns a tuple with the SubspaceIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetSubspaceIpv4Ok() (*string, bool) {
	if o == nil || o.SubspaceIpv4 == nil {
		return nil, false
	}
	return o.SubspaceIpv4, true
}

// HasSubspaceIpv4 returns a boolean if a field has been set.
func (o *V1Accelerator) HasSubspaceIpv4() bool {
	if o != nil && o.SubspaceIpv4 != nil {
		return true
	}

	return false
}

// SetSubspaceIpv4 gets a reference to the given string and assigns it to the SubspaceIpv4 field.
func (o *V1Accelerator) SetSubspaceIpv4(v string) {
	o.SubspaceIpv4 = &v
}

// GetSubspacePort returns the SubspacePort field value if set, zero value otherwise.
func (o *V1Accelerator) GetSubspacePort() int64 {
	if o == nil || o.SubspacePort == nil {
		var ret int64
		return ret
	}
	return *o.SubspacePort
}

// GetSubspacePortOk returns a tuple with the SubspacePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Accelerator) GetSubspacePortOk() (*int64, bool) {
	if o == nil || o.SubspacePort == nil {
		return nil, false
	}
	return o.SubspacePort, true
}

// HasSubspacePort returns a boolean if a field has been set.
func (o *V1Accelerator) HasSubspacePort() bool {
	if o != nil && o.SubspacePort != nil {
		return true
	}

	return false
}

// SetSubspacePort gets a reference to the given int64 and assigns it to the SubspacePort field.
func (o *V1Accelerator) SetSubspacePort(v int64) {
	o.SubspacePort = &v
}

func (o V1Accelerator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationIp != nil {
		toSerialize["destination_ip"] = o.DestinationIp
	}
	if o.DestinationPort != nil {
		toSerialize["destination_port"] = o.DestinationPort
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SubspaceIpv4 != nil {
		toSerialize["subspace_ipv4"] = o.SubspaceIpv4
	}
	if o.SubspacePort != nil {
		toSerialize["subspace_port"] = o.SubspacePort
	}
	return json.Marshal(toSerialize)
}

type NullableV1Accelerator struct {
	value *V1Accelerator
	isSet bool
}

func (v NullableV1Accelerator) Get() *V1Accelerator {
	return v.value
}

func (v *NullableV1Accelerator) Set(val *V1Accelerator) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Accelerator) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Accelerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Accelerator(val *V1Accelerator) *NullableV1Accelerator {
	return &NullableV1Accelerator{value: val, isSet: true}
}

func (v NullableV1Accelerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Accelerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


