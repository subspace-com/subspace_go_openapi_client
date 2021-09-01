# Go API client for subspace_openapi_client

# Introduction

The Subspace API is based on REST, has resource-oriented URLs, returns JSON-encoded responses, and returns standard HTTP response codes.

The base URL for the API is:  `https://api.subspace.com/`

# Naming Convention

* Version name currently in use is: *v1*
  * Example: `https://api.subspace.com/v1`

# Authentication

## API Tokens

Subspace authenticates your API requests using JWT Bearer tokens. The provided client library requires this JWT to be set before it can be used, by setting the local access token in the local configuration.  This is done by updating the configuration line marked "ACCESSTOKENSTRING" by replacing the text "ACCESSTOKENSTRING" with your JWT Bearer token.

Bearer tokens are granted by requesting one (as noted below) and presenting your publishable (client_id) and secret (client_secret) tokens.   

Subspace provides two types of API tokens: publishable (client_id) and secret (client_secret).  These are available in the Subspace console.
  * **Publishable** API tokens (client_id) are meant solely to identify your account with Subspace, they aren’t secret. They can be published in places like your website JavaScript code, or in an iPhone or Android app.
  * **Secret** API tokens (client_secret) should be kept confidential and only stored on your own servers. Your account’s secret API token will allow you to acquire a valid JWT token authorized to perform any API request to Subspace.

## Getting a JWT Bearer Token

Subspace uses auth0 for JWT token management.  You can acquire a JWT token by utilizing `https://id.subspace.com` and following the instructions in the curl example below.

## Protecting Your API Tokens

  * **JWT tokens have a expiration time of 24 hours.**  Once expired, you will have to use your Subspace private API and public token to request a new one.
  * The Subspace private token can be rotated from within the Subspace console.  Rotation may take up to 10 minutes for all systems to update state to invalidate the older token and enable the new one.
  * **Keep your secret token safe.** Your secret token can make any API call on behalf of your account, including changes that may impact billing such as enabling pay-as-you-go charges. Do not store your secret token in your version control system. Do not use your secret token outside your web server, such as a browser, mobile app, or distributed file.
  * **You may use the Subspace console to acquire an API token.**
  * **You may use the Subspace console to disable pay-as-you-go.** This may prevent unexpected charges due to unauthorized or abnormal usage.
  * **Do not embed API keys directly in code.** Instead of directly embedding API keys in your application’s code, put them in environment variables, or within include files that are stored separately from the bulk of your code—outside the source repository of your application. Then, if you share your code, the API keys will not be included in the shared files.
  * **Do not store API tokens inside your application’s source control.** If you store API tokens in files, keep the files outside your application’s source control system. This is particularly important if you use a public source code management system such as GitHub.
  * **Limit access with restricted tokens.** The Subspace console will allow you to specify the IP addresses or referrer URLs associated with each token, reducing the impact of a compromised API token.
  * **Use independent API tokens for different apps.** This limits the scope of each token. If an API token is compromised, you can rotate the impacted token without impacting other API tokens.

# Error Codes

Subspace uses HTTP response codes to indicate the success or failure of an API request. 

General HTML status codes:
  * 2xx Success. 
  * 4xx Errors based on information provided in the request.
  * 5xx Errors on Subspace servers.
  
# Security

We provide a valid, signed certificate for our API methods. Be sure your connection library supports HTTPS with the SNI extension.

# REST How-To

Making your first REST API call is easy and can be done from your browser.  You will need:
  * Your **secret** token and public client token, both found in the Console.
  * The URL for the type of data you would like to request.

First, acquire a JWT Bearer Token.  Command line example:

```    
    curl --request POST \\
         --url \"https://id.subspace.com/oauth/token\" \\
         --header 'content-type: application/json' \\
         --data '{ \"client_id\": YOURCLIENTID, \"client_secret\": YOURCLIENTSECRET, \"audience\": \"https://api.subspace.com/\", \"grant_type\": \"client_credentials\" }'
```

REST calls are made up of:
  * Base url: Example: `https://api.subspace.com`
  * Version: Example: `v1`
  * The API Endpoint and any parameters: `accelerators/acc_NDA3MUI5QzUtOTY4MC00Nz` where `acc_NDA3MUI5QzUtOTY4MC00Nz` is a valid accelerator ID
  * Accelerator ids are always of the format `acc_NDA3MUI5QzUtOTY4MC00Nz`, with a \"acc_\" prefix followed by 22 characters.
  * Project ids are always of the format `prj_00Iaqxjo71vNL1uLKKo5Kn`, with a \"prj_\" prefix followed by 22 characters.
  * Token header: All REST requests require a valid JWT Bearer token which should be added as an “Authorization” header to the request:
      
      `Authorization: Bearer YOUR_TOKEN_HERE`
  
## Authorization header example

If your API token was “my_api_token”, you would add...

    Authorization: Bearer my_api_token
    
...to the header.

## Command line examples

To list your current open packet_accelerators using the token “my_api_token”:

    curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerators
    
Alternately, to get the details of a specific accelerator whose id is 'abcd-ef01-2345':

    curl -H “Authorization: Bearer my_api_token” https://api.subspace.com/v1/accelerators/abcd-ef01-2345

# API Versioning

Subspace will release new versions when we make backwards-incompatible changes to the API. We will give advance notice before releasing a new version or retiring an old version.

Backwards compatible changes:
  * Adding new response attributes
  * Adding new endpoints
  * Adding new methods to an existing endpoint
  * Adding new query string parameters
  * Adding new path parameters
  * Adding new webhook events
  * Adding new streaming endpoints
  * Changing the order of existing response attributes
  
Versions are added to the base url, for example:
  * `https://api.subspace.com/v1`

Current Version is **v1:** `https://api.subspace.com/v1`


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://subspace.com](https://subspace.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./subspace_openapi_client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.subspace.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AcceleratorServiceApi* | [**AcceleratorServiceCreate**](docs/AcceleratorServiceApi.md#acceleratorservicecreate) | **Post** /v1/accelerators | 
*AcceleratorServiceApi* | [**AcceleratorServiceDelete**](docs/AcceleratorServiceApi.md#acceleratorservicedelete) | **Delete** /v1/accelerators/{id} | 
*AcceleratorServiceApi* | [**AcceleratorServiceGet**](docs/AcceleratorServiceApi.md#acceleratorserviceget) | **Get** /v1/accelerators/{id} | 
*AcceleratorServiceApi* | [**AcceleratorServiceList**](docs/AcceleratorServiceApi.md#acceleratorservicelist) | **Get** /v1/accelerators | 
*AcceleratorServiceApi* | [**AcceleratorServiceUpdate**](docs/AcceleratorServiceApi.md#acceleratorserviceupdate) | **Put** /v1/accelerators/{id} | 
*ProjectServiceApi* | [**ProjectServiceCreate**](docs/ProjectServiceApi.md#projectservicecreate) | **Post** /v1/projects | 
*ProjectServiceApi* | [**ProjectServiceGet**](docs/ProjectServiceApi.md#projectserviceget) | **Get** /v1/projects/{id} | 
*ProjectServiceApi* | [**ProjectServiceList**](docs/ProjectServiceApi.md#projectservicelist) | **Get** /v1/projects | 
*ProjectServiceApi* | [**ProjectServiceUpdate**](docs/ProjectServiceApi.md#projectserviceupdate) | **Put** /v1/projects/{id} | 
*SessionServiceApi* | [**SessionServiceList**](docs/SessionServiceApi.md#sessionservicelist) | **Get** /v1/accelerators/{accelerator_id}/sessions | 
*SipTeleportServiceApi* | [**SipTeleportServiceCreate**](docs/SipTeleportServiceApi.md#sipteleportservicecreate) | **Post** /v1/sip-teleports | 
*SipTeleportServiceApi* | [**SipTeleportServiceDelete**](docs/SipTeleportServiceApi.md#sipteleportservicedelete) | **Delete** /v1/sip-teleports/{id} | 
*SipTeleportServiceApi* | [**SipTeleportServiceGet**](docs/SipTeleportServiceApi.md#sipteleportserviceget) | **Get** /v1/sip-teleports/{id} | 
*SipTeleportServiceApi* | [**SipTeleportServiceList**](docs/SipTeleportServiceApi.md#sipteleportservicelist) | **Get** /v1/sip-teleports | 
*SipTeleportServiceApi* | [**SipTeleportServiceUpdate**](docs/SipTeleportServiceApi.md#sipteleportserviceupdate) | **Put** /v1/sip-teleports/{id} | 


## Documentation For Models

 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [V1Accelerator](docs/V1Accelerator.md)
 - [V1CreateSipTeleport](docs/V1CreateSipTeleport.md)
 - [V1ListAcceleratorsResponse](docs/V1ListAcceleratorsResponse.md)
 - [V1ListProjectsResponse](docs/V1ListProjectsResponse.md)
 - [V1ListSessionsResponse](docs/V1ListSessionsResponse.md)
 - [V1ListSipTeleportResponse](docs/V1ListSipTeleportResponse.md)
 - [V1NextPage](docs/V1NextPage.md)
 - [V1Project](docs/V1Project.md)
 - [V1Protocol](docs/V1Protocol.md)
 - [V1Session](docs/V1Session.md)
 - [V1SipTeleportResponse](docs/V1SipTeleportResponse.md)
 - [V1SipTeleportStatus](docs/V1SipTeleportStatus.md)
 - [V1TeleportAddresses](docs/V1TeleportAddresses.md)
 - [V1TransportType](docs/V1TransportType.md)
 - [V1UpdateSipTeleport](docs/V1UpdateSipTeleport.md)


## Documentation For Authorization



### accessCode


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **accelerators:read**: allows reading details about provisioned PacketAccelerators
 - **accelerators:write**: allows administration of PacketAccelerators
 - **sipteleport:read**: allows reading details about provisioned SIPTeleport
 - **sipteleport:write**: allows administration of SIPTeleport
 - **sessions:read**: allows reading details about PacketAccelerator sessions
 - **sessions:write**: allows administration of PacketAccelerator sessions
 - **projects:read**: allows reading details about projects
 - **projects:write**: allows administration of projects

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

sales@subspace.com

