# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScribeServiceCheck**](ScribeServiceApi.md#ScribeServiceCheck) | **Post** /grpc/v1/health_check | see: https://github.com/grpc/grpc/blob/master/doc/health-checking.md
[**ScribeServiceFilterLogs**](ScribeServiceApi.md#ScribeServiceFilterLogs) | **Post** /grpc/v1/filter_logs | 
[**ScribeServiceStreamLogs**](ScribeServiceApi.md#ScribeServiceStreamLogs) | **Post** /grpc/v1/stream_logs | 
[**ScribeServiceWatch**](ScribeServiceApi.md#ScribeServiceWatch) | **Post** /grpc/v1/health_watch | 

# **ScribeServiceCheck**
> V1HealthCheckResponse ScribeServiceCheck(ctx, body)
see: https://github.com/grpc/grpc/blob/master/doc/health-checking.md

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1HealthCheckRequest**](V1HealthCheckRequest.md)|  | 

### Return type

[**V1HealthCheckResponse**](v1HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScribeServiceFilterLogs**
> V1FilterLogsResponse ScribeServiceFilterLogs(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1FilterLogsRequest**](V1FilterLogsRequest.md)|  | 

### Return type

[**V1FilterLogsResponse**](v1FilterLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScribeServiceStreamLogs**
> StreamResultOfV1StreamLogsResponse ScribeServiceStreamLogs(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1StreamLogsRequest**](V1StreamLogsRequest.md)|  | 

### Return type

[**StreamResultOfV1StreamLogsResponse**](Stream result of v1StreamLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScribeServiceWatch**
> StreamResultOfV1HealthCheckResponse ScribeServiceWatch(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1HealthCheckRequest**](V1HealthCheckRequest.md)|  | 

### Return type

[**StreamResultOfV1HealthCheckResponse**](Stream result of v1HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

