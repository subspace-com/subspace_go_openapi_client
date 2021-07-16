/*
 * Subspace Product API
 *
 * # Introduction  The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standardHTTP response codes.  The base URL for the API is `https://api.subspace.com/`  # Naming Convention  **EARLY ACCESS NOTE:** There is no “stable” version yet.  Once there is, the version name **stable** will be used to access the latest stable API version.   * Example: `https://api.subspace.com/stable` * Version name currently in use is: *v1*   * Example: `https://api.subspace.com/v1`  # Authentication  ## API Tokens  Subspace authenticates your API requests using JWT Bearer tokens.  The provided client library requires this JWT to be set before it can be used, by setting the local access token in the local configuration.  This is done by updating the configuration line marked \"YOUR ACCESS TOKEN\" by replacing the text \"YOUR ACCESS TOKEN\" with your JWT Bearer token.  Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.     Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.   * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.   * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.  ## Getting a JWT Bearer Token  Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.  ## Protecting Your API Tokens    * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.   * The Subspace private token can be rotated from within the Subspace console.  Rotation may take up to 10 minutes for all systems to update state to invalidate the older token and enable the new one. * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.   * **You may use the Subspace console to acquire an API token.**  * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.  **Acquiring a valid JWT**  Command line example: ``` curl --request POST           --url 'https://id.subspace.com/oauth/token'           --header 'content-type: application/json'           --data '{ \"client_id\": YOURCLIENTID, \"client_secret\": YOURCLIENTSECRET, \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }' ``` 
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
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	DateUpdated *time.Time `json:"dateUpdated,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *V1SipTeleportStatus `json:"status,omitempty"`
	TeleportEntryPoints *[]V1TeleportAddresses `json:"teleportEntryPoints,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewV1SipTeleportResponse instantiates a new V1SipTeleportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SipTeleportResponse() *V1SipTeleportResponse {
	this := V1SipTeleportResponse{}
	var status V1SipTeleportStatus = DISABLED
	this.Status = &status
	return &this
}

// NewV1SipTeleportResponseWithDefaults instantiates a new V1SipTeleportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SipTeleportResponseWithDefaults() *V1SipTeleportResponse {
	this := V1SipTeleportResponse{}
	var status V1SipTeleportStatus = DISABLED
	this.Status = &status
	return &this
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

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1SipTeleportResponse) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SipTeleportResponse) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1SipTeleportResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *V1SipTeleportResponse) SetVersion(v int64) {
	o.Version = &v
}

func (o V1SipTeleportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.DateUpdated != nil {
		toSerialize["dateUpdated"] = o.DateUpdated
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TeleportEntryPoints != nil {
		toSerialize["teleportEntryPoints"] = o.TeleportEntryPoints
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
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


