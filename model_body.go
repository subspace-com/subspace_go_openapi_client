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

// Body struct for Body
type Body struct {
	// Name of PacketAccelerator
	Name string `json:"name"`
	// Destination IP, e.g. 1.2.3.4
	DestinationIp string `json:"destination_ip"`
	// Destination port, e.g. 443
	DestinationPort int32 `json:"destination_port"`
	// Optional per plan, Specify the Subspace-returned ingress port
	SubspacePort *int32 `json:"subspace_port,omitempty"`
}

// NewBody instantiates a new Body object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBody(name string, destinationIp string, destinationPort int32) *Body {
	this := Body{}
	this.Name = name
	this.DestinationIp = destinationIp
	this.DestinationPort = destinationPort
	return &this
}

// NewBodyWithDefaults instantiates a new Body object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyWithDefaults() *Body {
	this := Body{}
	return &this
}

// GetName returns the Name field value
func (o *Body) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Body) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Body) SetName(v string) {
	o.Name = v
}

// GetDestinationIp returns the DestinationIp field value
func (o *Body) GetDestinationIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIp
}

// GetDestinationIpOk returns a tuple with the DestinationIp field value
// and a boolean to check if the value has been set.
func (o *Body) GetDestinationIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationIp, true
}

// SetDestinationIp sets field value
func (o *Body) SetDestinationIp(v string) {
	o.DestinationIp = v
}

// GetDestinationPort returns the DestinationPort field value
func (o *Body) GetDestinationPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DestinationPort
}

// GetDestinationPortOk returns a tuple with the DestinationPort field value
// and a boolean to check if the value has been set.
func (o *Body) GetDestinationPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationPort, true
}

// SetDestinationPort sets field value
func (o *Body) SetDestinationPort(v int32) {
	o.DestinationPort = v
}

// GetSubspacePort returns the SubspacePort field value if set, zero value otherwise.
func (o *Body) GetSubspacePort() int32 {
	if o == nil || o.SubspacePort == nil {
		var ret int32
		return ret
	}
	return *o.SubspacePort
}

// GetSubspacePortOk returns a tuple with the SubspacePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Body) GetSubspacePortOk() (*int32, bool) {
	if o == nil || o.SubspacePort == nil {
		return nil, false
	}
	return o.SubspacePort, true
}

// HasSubspacePort returns a boolean if a field has been set.
func (o *Body) HasSubspacePort() bool {
	if o != nil && o.SubspacePort != nil {
		return true
	}

	return false
}

// SetSubspacePort gets a reference to the given int32 and assigns it to the SubspacePort field.
func (o *Body) SetSubspacePort(v int32) {
	o.SubspacePort = &v
}

func (o Body) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["destination_ip"] = o.DestinationIp
	}
	if true {
		toSerialize["destination_port"] = o.DestinationPort
	}
	if o.SubspacePort != nil {
		toSerialize["subspace_port"] = o.SubspacePort
	}
	return json.Marshal(toSerialize)
}

type NullableBody struct {
	value *Body
	isSet bool
}

func (v NullableBody) Get() *Body {
	return v.value
}

func (v *NullableBody) Set(val *Body) {
	v.value = val
	v.isSet = true
}

func (v NullableBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBody(val *Body) *NullableBody {
	return &NullableBody{value: val, isSet: true}
}

func (v NullableBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


