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
	"time"
)

// V1SipTeleportResponse struct for V1SipTeleportResponse
type V1SipTeleportResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Destination *string `json:"destination,omitempty"`
	TeleportEntryPoints *[]V1TeleportAddresses `json:"teleport_entry_points,omitempty"`
	Status *V1SipTeleportStatus `json:"status,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}

// NewV1SipTeleportResponse instantiates a new V1SipTeleportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SipTeleportResponse() *V1SipTeleportResponse {
	this := V1SipTeleportResponse{}
	var status V1SipTeleportStatus = UNKNOWN
	this.Status = &status
	return &this
}

// NewV1SipTeleportResponseWithDefaults instantiates a new V1SipTeleportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SipTeleportResponseWithDefaults() *V1SipTeleportResponse {
	this := V1SipTeleportResponse{}
	var status V1SipTeleportStatus = UNKNOWN
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1SipTeleportResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1SipTeleportResponse) SetName(v string) {
	o.Name = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *V1SipTeleportResponse) SetDestination(v string) {
	o.Destination = &v
}

// GetTeleportEntryPoints returns the TeleportEntryPoints field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetTeleportEntryPoints() []V1TeleportAddresses {
	if o == nil || o.TeleportEntryPoints == nil {
		var ret []V1TeleportAddresses
		return ret
	}
	return *o.TeleportEntryPoints
}

// GetTeleportEntryPointsOk returns a tuple with the TeleportEntryPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetTeleportEntryPointsOk() (*[]V1TeleportAddresses, bool) {
	if o == nil || o.TeleportEntryPoints == nil {
		return nil, false
	}
	return o.TeleportEntryPoints, true
}

// HasTeleportEntryPoints returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasTeleportEntryPoints() bool {
	if o != nil && o.TeleportEntryPoints != nil {
		return true
	}

	return false
}

// SetTeleportEntryPoints gets a reference to the given []V1TeleportAddresses and assigns it to the TeleportEntryPoints field.
func (o *V1SipTeleportResponse) SetTeleportEntryPoints(v []V1TeleportAddresses) {
	o.TeleportEntryPoints = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetStatus() V1SipTeleportStatus {
	if o == nil || o.Status == nil {
		var ret V1SipTeleportStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetStatusOk() (*V1SipTeleportStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1SipTeleportStatus and assigns it to the Status field.
func (o *V1SipTeleportResponse) SetStatus(v V1SipTeleportStatus) {
	o.Status = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *V1SipTeleportResponse) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *V1SipTeleportResponse) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

func (o V1SipTeleportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.TeleportEntryPoints != nil {
		toSerialize["teleport_entry_points"] = o.TeleportEntryPoints
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.DateUpdated != nil {
		toSerialize["date_updated"] = o.DateUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableV1SipTeleportResponse struct {
	value *V1SipTeleportResponse
	isSet bool
}

func (v NullableV1SipTeleportResponse) Get() *V1SipTeleportResponse {
	return v.value
}

func (v *NullableV1SipTeleportResponse) Set(val *V1SipTeleportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SipTeleportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SipTeleportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SipTeleportResponse(val *V1SipTeleportResponse) *NullableV1SipTeleportResponse {
	return &NullableV1SipTeleportResponse{value: val, isSet: true}
}

func (v NullableV1SipTeleportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SipTeleportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


