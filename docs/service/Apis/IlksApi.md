# IlksApi

All URIs are relative to *https://api.zarban.io*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAllIlks**](IlksApi.md#getAllIlks) | **GET** /v2/ilks | Get all Ilks |
| [**getIlkByName**](IlksApi.md#getIlkByName) | **GET** /v2/ilks/{name} | Get Ilk by name |


<a name="getAllIlks"></a>
# **getAllIlks**
> IlksResponse getAllIlks()

Get all Ilks

    Retrieve a list of all available ilks.

### Parameters
This endpoint does not need any parameter.

### Return type

[**IlksResponse**](../Models/IlksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getIlkByName"></a>
# **getIlkByName**
> Ilk getIlkByName(name)

Get Ilk by name

    Retrieve an Ilk by providing its name.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **name** | **String**| Name of the ILK | [default to null] |

### Return type

[**Ilk**](../Models/Ilk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

