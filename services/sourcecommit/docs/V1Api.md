# \V1Api

All URIs are relative to *https://sourcecommit.apigw.ntruss.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Repository**](V1Api.md#Repository) | **Post** /repository/{repositoryName}/branch/default | 
[**RepositoryRepositoryNameDelete**](V1Api.md#RepositoryRepositoryNameDelete) | **Delete** /repository/{repositoryName} | 
[**RepositoryRepositoryNamePatch**](V1Api.md#RepositoryRepositoryNamePatch) | **Patch** /repository/{repositoryName} | 
[**Repository_0**](V1Api.md#Repository_0) | **Post** /repository | 


# **Repository**
> OkResponse Repository(body, repositoryName, xNcpRegionCode)


set defalut branch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*DefaultBranch](DefaultBranch.md)** |  | **repositoryName** | **string** | repositoryName | **xNcpRegionCode** | **string** | NCP 리전 코드  | 

### Return type

*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryRepositoryNameDelete**
> OkResponse RepositoryRepositoryNameDelete(repositoryName, xNcpRegionCode)


delete repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**repositoryName** | **string** | repositoryName | **xNcpRegionCode** | **string** | NCP 리전 코드 | 

### Return type

*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryRepositoryNamePatch**
> OkResponse RepositoryRepositoryNamePatch(body, repositoryName, xNcpRegionCode)


change repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*ChangeRepository](ChangeRepository.md)** |  | **repositoryName** | **string** | repositoryName | **xNcpRegionCode** | **string** | NCP 리전 코드 | 

### Return type

*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Repository_0**
> OkResponse Repository_0(body, xNcpRegionCode)


create repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | **[\*CreateRepository](CreateRepository.md)** |  | **xNcpRegionCode** | **string** | NCP 리전 코드 | 

### Return type

*[**OkResponse**](OkResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

