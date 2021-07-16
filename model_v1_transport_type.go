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
	"fmt"
)

// V1TransportType the model 'V1TransportType'
type V1TransportType string

// List of v1TransportType
const (
	UDP_TCP V1TransportType = "UDP_TCP"
	TLS V1TransportType = "TLS"
)

var allowedV1TransportTypeEnumValues = []V1TransportType{
	"UDP_TCP",
	"TLS",
}

func (v *V1TransportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1TransportType(value)
	for _, existing := range allowedV1TransportTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1TransportType", value)
}

// NewV1TransportTypeFromValue returns a pointer to a valid V1TransportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1TransportTypeFromValue(v string) (*V1TransportType, error) {
	ev := V1TransportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1TransportType: valid values are %v", v, allowedV1TransportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1TransportType) IsValid() bool {
	for _, existing := range allowedV1TransportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1TransportType value
func (v V1TransportType) Ptr() *V1TransportType {
	return &v
}

type NullableV1TransportType struct {
	value *V1TransportType
	isSet bool
}

func (v NullableV1TransportType) Get() *V1TransportType {
	return v.value
}

func (v *NullableV1TransportType) Set(val *V1TransportType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TransportType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TransportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TransportType(val *V1TransportType) *NullableV1TransportType {
	return &NullableV1TransportType{value: val, isSet: true}
}

func (v NullableV1TransportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TransportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

