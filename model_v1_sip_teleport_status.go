/*
 * Subspace Product API
 *
 * # Introduction  The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standard HTTP response codes.  The base URL for the API is:  `https://api.subspace.com/`  # Naming Convention  * Version name currently in use is: *v1*   * Example: `https://api.subspace.com/v1`  # Authentication  ## API Tokens  Subspace authenticates your API requests using JWT Bearer tokens. To use any Subspace API, you must pass a Bearer token with each request. If you do not include your Bearer token when making an API request, or use one that is incorrect or disabled, Subspace returns an error.  Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.     Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.   * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.   * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.  ## Getting a JWT Bearer Token  Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.  ## Protecting Your API Tokens    * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.   * The Subspace private token can be rotated from within the Subspace console.  Rotation may take up to 10 minutes for all systems to update state to invalidate the older token and enable the new one.   * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.   * **You may use the Subspace console to acquire an API token.**   * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.   * **Do not embed API keys directly in code.** Instead of directly embedding API keys in your application’s code, put them in environment variables, or within include files that are stored separately from the bulk of your code—outside the source repository of your application. Then, if you share your code, the API keys will not be included in the shared files.   * **Do not store API tokens inside your application’s source control.** If you store API tokens in files, keep the files outside your application’s source control system. This is particularly important if you use a public source code management system such as GitHub.   * **Limit access with restricted tokens.** The Subspace console will allow you to specify the IP addresses or referrer URLs associated with each token, reducing the impact of a compromised API token.   * **Use independent API tokens for different apps.** This limits the scope of each token. If an API token is compromised, you can rotate the impacted token without impacting other API tokens.  # Error Codes  Subspace uses HTTP response codes to indicate the success or failure of an API request.   General HTML status codes:   * 2xx Success.    * 4xx Errors based on information provided in the request.   * 5xx Errors on Subspace servers.    # Security  We provide a valid, signed certificate for our API methods. Be sure your connection library supports HTTPS with the SNI extension.  # REST How-To  Making your first REST API call is easy and can be done from your browser.  You will need:   * Your **secret** token and public client token, both found in the Console.   * The URL for the type of data you would like to request.  First, acquire a JWT Bearer Token.  Command line example:          curl --request POST \\          --url \"https://id.subspace.com/oauth/token\" \\          --header 'content-type: application/json' \\          --data '{ \"client_id\": YOURCLIENTID, \"client_secret\": YOURCLIENTSECRET, \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }'  REST calls are made up of:   * Base url: Example: `https://api.subspace.com`   * Version: Example: `v1`   * The API Endpoint and any parameters: `accelerators/acc_NDA3MUI5QzUtOTY4MC00Nz` where `acc_NDA3MUI5QzUtOTY4MC00Nz` is a valid accelerator ID   * Accelerator ids are always of the format `acc_NDA3MUI5QzUtOTY4MC00Nz`, with a \"acc_\" prefix followed by 22 characters.   * Project ids are always of the format `prj_00Iaqxjo71vNL1uLKKo5Kn`, with a \"prj_\" prefix followed by 22 characters.   * Token header: All REST requests require a valid JWT Bearer token which should be added as an “Authorization” header to the request:              `Authorization: Bearer YOUR_TOKEN_HERE`    ## Authorization header example  If your API token was “my_api_token”, you would add...      Authorization: Bearer my_api_token      ...to the header.  ## Command line examples  To list your current open packet_accelerators using the token “my_api_token”:      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerators      Alternately, to get the details of a specific accelerator whose id is 'abcd-ef01-2345':      curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerators/abcd-ef01-2345  # API Versioning  Subspace will release new versions when we make backwards-incompatible changes to the API. We will give advance notice before releasing a new version or retiring an old version.  Backwards compatible changes:   * Adding new response attributes   * Adding new endpoints   * Adding new methods to an existing endpoint   * Adding new query string parameters   * Adding new path parameters   * Adding new webhook events   * Adding new streaming endpoints   * Changing the order of existing response attributes    Versions are added to the base url, for example:   * `https://api.subspace.com/v1`  Current Version is **v1:** `https://api.subspace.com/v1` 
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

// V1SipTeleportStatus the model 'V1SipTeleportStatus'
type V1SipTeleportStatus string

// List of v1SipTeleportStatus
const (
	DISABLED V1SipTeleportStatus = "DISABLED"
	ENABLED V1SipTeleportStatus = "ENABLED"
)

var allowedV1SipTeleportStatusEnumValues = []V1SipTeleportStatus{
	"DISABLED",
	"ENABLED",
}

func (v *V1SipTeleportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1SipTeleportStatus(value)
	for _, existing := range allowedV1SipTeleportStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1SipTeleportStatus", value)
}

// NewV1SipTeleportStatusFromValue returns a pointer to a valid V1SipTeleportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1SipTeleportStatusFromValue(v string) (*V1SipTeleportStatus, error) {
	ev := V1SipTeleportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1SipTeleportStatus: valid values are %v", v, allowedV1SipTeleportStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1SipTeleportStatus) IsValid() bool {
	for _, existing := range allowedV1SipTeleportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1SipTeleportStatus value
func (v V1SipTeleportStatus) Ptr() *V1SipTeleportStatus {
	return &v
}

type NullableV1SipTeleportStatus struct {
	value *V1SipTeleportStatus
	isSet bool
}

func (v NullableV1SipTeleportStatus) Get() *V1SipTeleportStatus {
	return v.value
}

func (v *NullableV1SipTeleportStatus) Set(val *V1SipTeleportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SipTeleportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SipTeleportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SipTeleportStatus(val *V1SipTeleportStatus) *NullableV1SipTeleportStatus {
	return &NullableV1SipTeleportStatus{value: val, isSet: true}
}

func (v NullableV1SipTeleportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SipTeleportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

