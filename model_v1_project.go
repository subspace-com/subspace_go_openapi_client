/*
 * Subspace Product API
 *
 * # Introduction  The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standard HTTP response codes.  The base URL for the API is:  `https://api.subspace.com/`  # Naming Convention  * Version name currently in use is: *v1*   * Example: `https://api.subspace.com/v1`  # Authentication  ## API Tokens  Subspace authenticates your API requests using JWT Bearer tokens. To use any Subspace API, you must pass a Bearer token with each request. If you do not include your Bearer token when making an API request, or use one that is incorrect or disabled, Subspace returns an error.  Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.     Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.   * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.   * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.  ## Getting a JWT Bearer Token  Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.  ## Protecting Your API Tokens    * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.   * The Subspace private token can be rotated from within the Subspace console.  Rotation may take up to 10 minutes for all systems to update state to invalidate the older token and enable the new one.   * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.   * **You may use the Subspace console to acquire an API token.**   * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.   * **Do not embed API keys directly in code.** Instead of directly embedding API keys in your application’s code, put them in environment variables, or within include files that are stored separately from the bulk of your code—outside the source repository of your application. Then, if you share your code, the API keys will not be included in the shared files.   * **Do not store API tokens inside your application’s source control.** If you store API tokens in files, keep the files outside your application’s source control system. This is particularly important if you use a public source code management system such as GitHub.   * **Limit access with restricted tokens.** The Subspace console will allow you to specify the IP addresses or referrer URLs associated with each token, reducing the impact of a compromised API token.   * **Use independent API tokens for different apps.** This limits the scope of each token. If an API token is compromised, you can rotate the impacted token without impacting other API tokens.  # Error Codes  Subspace uses HTTP response codes to indicate the success or failure of an API request.   General HTML status codes:   * 2xx Success.    * 4xx Errors based on information provided in the request.   * 5xx Errors on Subspace servers.    # Security  We provide a valid, signed certificate for our API methods. Be sure your connection library supports HTTPS with the SNI extension.  # REST How-To  Making your first REST API call is easy and can be done from your browser.  You will need:   * Your **secret** token and public client token, both found in the Console.   * The URL for the type of data you would like to request.  First, acquire a JWT Bearer Token.  Command line example:          curl --request POST \\          --url \"https://id.subspace.com/oauth/token\" \\          --header 'content-type: application/json' \\          --data '{ \"client_id\": \"YOURCLIENTID\", \"client_secret\": \"YOURCLIENTSECRET\", \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }'  REST calls are made up of:   * Base url: Example: `https://api.subspace.com`   * Version: Example: `v1`   * The API Endpoint and any parameters: `accelerator/acc_NDA3MUI5QzUtOTY4MC00Nz` where `acc_NDA3MUI5QzUtOTY4MC00Nz` is a valid accelerator ID   * Accelerator ids are always of the format `acc_NDA3MUI5QzUtOTY4MC00Nz`, with a \"acc_\" prefix followed by 22 characters.   * Project ids are always of the format `prj_00Iaqxjo71vNL1uLKKo5Kn`, with a \"prj_\" prefix followed by 22 characters.   * Token header: All REST requests require a valid JWT Bearer token which should be added as an “Authorization” header to the request:              `Authorization: Bearer YOUR_TOKEN_HERE`    ## Authorization header example  If your API token was “my_api_token”, you would add...      Authorization: Bearer my_api_token      ...to the header.  ## Command line examples  To list your current open packet_accelerators using the token “my_api_token”:      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator      Alternately, to get the details of a specific accelerator whose id is 'abcd-ef01-2345':      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerator/abcd-ef01-2345  # API Versioning  Subspace will release new versions when we make backwards-incompatible changes to the API. We will give advance notice before releasing a new version or retiring an old version.  Backwards compatible changes:   * Adding new response attributes   * Adding new endpoints   * Adding new methods to an existing endpoint   * Adding new query string parameters   * Adding new path parameters   * Adding new webhook events   * Adding new streaming endpoints   * Changing the order of existing response attributes    Versions are added to the base url, for example:   * `https://api.subspace.com/v1`  Current Version is **v1:** `https://api.subspace.com/v1` 
 *
 * API version: 1.0
 * Contact: sales@subspace.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subspace_openapi_client

import (
	"encoding/json"
)

// V1Project struct for V1Project
type V1Project struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AcceleratorQuota *int64 `json:"accelerator_quota,omitempty"`
	AcceleratorRequestPort *bool `json:"accelerator_request_port,omitempty"`
	GlobalturnLimitGb *int64 `json:"globalturn_limit_gb,omitempty"`
	SipteleportQuota *int64 `json:"sipteleport_quota,omitempty"`
	SipteleportCallQuota *int64 `json:"sipteleport_call_quota,omitempty"`
	RtpspeedLimitGb *int64 `json:"rtpspeed_limit_gb,omitempty"`
}

// NewV1Project instantiates a new V1Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Project() *V1Project {
	this := V1Project{}
	return &this
}

// NewV1ProjectWithDefaults instantiates a new V1Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProjectWithDefaults() *V1Project {
	this := V1Project{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1Project) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1Project) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1Project) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1Project) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1Project) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1Project) SetName(v string) {
	o.Name = &v
}

// GetAcceleratorQuota returns the AcceleratorQuota field value if set, zero value otherwise.
func (o *V1Project) GetAcceleratorQuota() int64 {
	if o == nil || o.AcceleratorQuota == nil {
		var ret int64
		return ret
	}
	return *o.AcceleratorQuota
}

// GetAcceleratorQuotaOk returns a tuple with the AcceleratorQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetAcceleratorQuotaOk() (*int64, bool) {
	if o == nil || o.AcceleratorQuota == nil {
		return nil, false
	}
	return o.AcceleratorQuota, true
}

// HasAcceleratorQuota returns a boolean if a field has been set.
func (o *V1Project) HasAcceleratorQuota() bool {
	if o != nil && o.AcceleratorQuota != nil {
		return true
	}

	return false
}

// SetAcceleratorQuota gets a reference to the given int64 and assigns it to the AcceleratorQuota field.
func (o *V1Project) SetAcceleratorQuota(v int64) {
	o.AcceleratorQuota = &v
}

// GetAcceleratorRequestPort returns the AcceleratorRequestPort field value if set, zero value otherwise.
func (o *V1Project) GetAcceleratorRequestPort() bool {
	if o == nil || o.AcceleratorRequestPort == nil {
		var ret bool
		return ret
	}
	return *o.AcceleratorRequestPort
}

// GetAcceleratorRequestPortOk returns a tuple with the AcceleratorRequestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetAcceleratorRequestPortOk() (*bool, bool) {
	if o == nil || o.AcceleratorRequestPort == nil {
		return nil, false
	}
	return o.AcceleratorRequestPort, true
}

// HasAcceleratorRequestPort returns a boolean if a field has been set.
func (o *V1Project) HasAcceleratorRequestPort() bool {
	if o != nil && o.AcceleratorRequestPort != nil {
		return true
	}

	return false
}

// SetAcceleratorRequestPort gets a reference to the given bool and assigns it to the AcceleratorRequestPort field.
func (o *V1Project) SetAcceleratorRequestPort(v bool) {
	o.AcceleratorRequestPort = &v
}

// GetGlobalturnLimitGb returns the GlobalturnLimitGb field value if set, zero value otherwise.
func (o *V1Project) GetGlobalturnLimitGb() int64 {
	if o == nil || o.GlobalturnLimitGb == nil {
		var ret int64
		return ret
	}
	return *o.GlobalturnLimitGb
}

// GetGlobalturnLimitGbOk returns a tuple with the GlobalturnLimitGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetGlobalturnLimitGbOk() (*int64, bool) {
	if o == nil || o.GlobalturnLimitGb == nil {
		return nil, false
	}
	return o.GlobalturnLimitGb, true
}

// HasGlobalturnLimitGb returns a boolean if a field has been set.
func (o *V1Project) HasGlobalturnLimitGb() bool {
	if o != nil && o.GlobalturnLimitGb != nil {
		return true
	}

	return false
}

// SetGlobalturnLimitGb gets a reference to the given int64 and assigns it to the GlobalturnLimitGb field.
func (o *V1Project) SetGlobalturnLimitGb(v int64) {
	o.GlobalturnLimitGb = &v
}

// GetSipteleportQuota returns the SipteleportQuota field value if set, zero value otherwise.
func (o *V1Project) GetSipteleportQuota() int64 {
	if o == nil || o.SipteleportQuota == nil {
		var ret int64
		return ret
	}
	return *o.SipteleportQuota
}

// GetSipteleportQuotaOk returns a tuple with the SipteleportQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetSipteleportQuotaOk() (*int64, bool) {
	if o == nil || o.SipteleportQuota == nil {
		return nil, false
	}
	return o.SipteleportQuota, true
}

// HasSipteleportQuota returns a boolean if a field has been set.
func (o *V1Project) HasSipteleportQuota() bool {
	if o != nil && o.SipteleportQuota != nil {
		return true
	}

	return false
}

// SetSipteleportQuota gets a reference to the given int64 and assigns it to the SipteleportQuota field.
func (o *V1Project) SetSipteleportQuota(v int64) {
	o.SipteleportQuota = &v
}

// GetSipteleportCallQuota returns the SipteleportCallQuota field value if set, zero value otherwise.
func (o *V1Project) GetSipteleportCallQuota() int64 {
	if o == nil || o.SipteleportCallQuota == nil {
		var ret int64
		return ret
	}
	return *o.SipteleportCallQuota
}

// GetSipteleportCallQuotaOk returns a tuple with the SipteleportCallQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetSipteleportCallQuotaOk() (*int64, bool) {
	if o == nil || o.SipteleportCallQuota == nil {
		return nil, false
	}
	return o.SipteleportCallQuota, true
}

// HasSipteleportCallQuota returns a boolean if a field has been set.
func (o *V1Project) HasSipteleportCallQuota() bool {
	if o != nil && o.SipteleportCallQuota != nil {
		return true
	}

	return false
}

// SetSipteleportCallQuota gets a reference to the given int64 and assigns it to the SipteleportCallQuota field.
func (o *V1Project) SetSipteleportCallQuota(v int64) {
	o.SipteleportCallQuota = &v
}

// GetRtpspeedLimitGb returns the RtpspeedLimitGb field value if set, zero value otherwise.
func (o *V1Project) GetRtpspeedLimitGb() int64 {
	if o == nil || o.RtpspeedLimitGb == nil {
		var ret int64
		return ret
	}
	return *o.RtpspeedLimitGb
}

// GetRtpspeedLimitGbOk returns a tuple with the RtpspeedLimitGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Project) GetRtpspeedLimitGbOk() (*int64, bool) {
	if o == nil || o.RtpspeedLimitGb == nil {
		return nil, false
	}
	return o.RtpspeedLimitGb, true
}

// HasRtpspeedLimitGb returns a boolean if a field has been set.
func (o *V1Project) HasRtpspeedLimitGb() bool {
	if o != nil && o.RtpspeedLimitGb != nil {
		return true
	}

	return false
}

// SetRtpspeedLimitGb gets a reference to the given int64 and assigns it to the RtpspeedLimitGb field.
func (o *V1Project) SetRtpspeedLimitGb(v int64) {
	o.RtpspeedLimitGb = &v
}

func (o V1Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AcceleratorQuota != nil {
		toSerialize["accelerator_quota"] = o.AcceleratorQuota
	}
	if o.AcceleratorRequestPort != nil {
		toSerialize["accelerator_request_port"] = o.AcceleratorRequestPort
	}
	if o.GlobalturnLimitGb != nil {
		toSerialize["globalturn_limit_gb"] = o.GlobalturnLimitGb
	}
	if o.SipteleportQuota != nil {
		toSerialize["sipteleport_quota"] = o.SipteleportQuota
	}
	if o.SipteleportCallQuota != nil {
		toSerialize["sipteleport_call_quota"] = o.SipteleportCallQuota
	}
	if o.RtpspeedLimitGb != nil {
		toSerialize["rtpspeed_limit_gb"] = o.RtpspeedLimitGb
	}
	return json.Marshal(toSerialize)
}

type NullableV1Project struct {
	value *V1Project
	isSet bool
}

func (v NullableV1Project) Get() *V1Project {
	return v.value
}

func (v *NullableV1Project) Set(val *V1Project) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Project) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Project) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Project(val *V1Project) *NullableV1Project {
	return &NullableV1Project{value: val, isSet: true}
}

func (v NullableV1Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Project) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


