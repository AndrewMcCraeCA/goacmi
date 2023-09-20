# Go API client for goacmi

A public API for ACMI's collection data including Films, TV Shows, Videogames and Art.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.acmi.net.au](https://www.acmi.net.au)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import goacmi "github.com/AndrewMcCraeCA/goacmi"
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
ctx := context.WithValue(context.Background(), goacmi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), goacmi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), goacmi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), goacmi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.acmi.net.au*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**RootGet**](docs/DefaultAPI.md#rootget) | **Get** / | 
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Get** /search/ | 
*WorkAPI* | [**GetWork**](docs/WorkAPI.md#getwork) | **Get** /works/{id}/ | 
*WorksAPI* | [**GetWorks**](docs/WorksAPI.md#getworks) | **Get** /works/ | 


## Documentation For Models

 - [ConstellationSummary](docs/ConstellationSummary.md)
 - [CreatorSummary](docs/CreatorSummary.md)
 - [Get200Response](docs/Get200Response.md)
 - [VideoMetadata](docs/VideoMetadata.md)
 - [Work](docs/Work.md)
 - [WorkDetailsInner](docs/WorkDetailsInner.md)
 - [WorkExternalReferencesInner](docs/WorkExternalReferencesInner.md)
 - [WorkHoldingsInner](docs/WorkHoldingsInner.md)
 - [WorkImagesInner](docs/WorkImagesInner.md)
 - [WorkLinksInner](docs/WorkLinksInner.md)
 - [WorkProductionDatesInner](docs/WorkProductionDatesInner.md)
 - [WorkProductionPlacesInner](docs/WorkProductionPlacesInner.md)
 - [WorkRecommendationsInner](docs/WorkRecommendationsInner.md)
 - [WorkSource](docs/WorkSource.md)
 - [WorkStats](docs/WorkStats.md)
 - [WorkSummary](docs/WorkSummary.md)
 - [WorkThumbnail](docs/WorkThumbnail.md)
 - [WorkVideoLinksInner](docs/WorkVideoLinksInner.md)
 - [WorkVideosInner](docs/WorkVideosInner.md)


## Documentation For Authorization

Endpoints do not require authorization.


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


