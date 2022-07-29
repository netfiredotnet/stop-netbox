# \UsersApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigList**](UsersApi.md#UsersConfigList) | **Get** /users/config/ | 
[**UsersGroupsBulkDelete**](UsersApi.md#UsersGroupsBulkDelete) | **Delete** /users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersApi.md#UsersGroupsBulkPartialUpdate) | **Patch** /users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersApi.md#UsersGroupsBulkUpdate) | **Put** /users/groups/ | 
[**UsersGroupsCreate**](UsersApi.md#UsersGroupsCreate) | **Post** /users/groups/ | 
[**UsersGroupsDelete**](UsersApi.md#UsersGroupsDelete) | **Delete** /users/groups/{id}/ | 
[**UsersGroupsList**](UsersApi.md#UsersGroupsList) | **Get** /users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersApi.md#UsersGroupsPartialUpdate) | **Patch** /users/groups/{id}/ | 
[**UsersGroupsRead**](UsersApi.md#UsersGroupsRead) | **Get** /users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersApi.md#UsersGroupsUpdate) | **Put** /users/groups/{id}/ | 
[**UsersPermissionsBulkDelete**](UsersApi.md#UsersPermissionsBulkDelete) | **Delete** /users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersApi.md#UsersPermissionsBulkPartialUpdate) | **Patch** /users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersApi.md#UsersPermissionsBulkUpdate) | **Put** /users/permissions/ | 
[**UsersPermissionsCreate**](UsersApi.md#UsersPermissionsCreate) | **Post** /users/permissions/ | 
[**UsersPermissionsDelete**](UsersApi.md#UsersPermissionsDelete) | **Delete** /users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersApi.md#UsersPermissionsList) | **Get** /users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersApi.md#UsersPermissionsPartialUpdate) | **Patch** /users/permissions/{id}/ | 
[**UsersPermissionsRead**](UsersApi.md#UsersPermissionsRead) | **Get** /users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersApi.md#UsersPermissionsUpdate) | **Put** /users/permissions/{id}/ | 
[**UsersUsersBulkDelete**](UsersApi.md#UsersUsersBulkDelete) | **Delete** /users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersApi.md#UsersUsersBulkPartialUpdate) | **Patch** /users/users/ | 
[**UsersUsersBulkUpdate**](UsersApi.md#UsersUsersBulkUpdate) | **Put** /users/users/ | 
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /users/users/ | 
[**UsersUsersDelete**](UsersApi.md#UsersUsersDelete) | **Delete** /users/users/{id}/ | 
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /users/users/ | 
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | 
[**UsersUsersRead**](UsersApi.md#UsersUsersRead) | **Get** /users/users/{id}/ | 
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | 



## UsersConfigList

> UsersConfigList(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersConfigList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersConfigList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersConfigListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkDelete

> UsersGroupsBulkDelete(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkPartialUpdate

> Group UsersGroupsBulkPartialUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewGroup("Name_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkPartialUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkUpdate

> Group UsersGroupsBulkUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewGroup("Name_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsCreate

> Group UsersGroupsCreate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewGroup("Name_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsCreate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsDelete

> UsersGroupsDelete(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsList

> UsersGroupsList200Response UsersGroupsList(ctx).Id(id).Name(name).Q(q).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    idN := "idN_example" // string |  (optional)
    idLte := "idLte_example" // string |  (optional)
    idLt := "idLt_example" // string |  (optional)
    idGte := "idGte_example" // string |  (optional)
    idGt := "idGt_example" // string |  (optional)
    nameN := "nameN_example" // string |  (optional)
    nameIc := "nameIc_example" // string |  (optional)
    nameNic := "nameNic_example" // string |  (optional)
    nameIew := "nameIew_example" // string |  (optional)
    nameNiew := "nameNiew_example" // string |  (optional)
    nameIsw := "nameIsw_example" // string |  (optional)
    nameNisw := "nameNisw_example" // string |  (optional)
    nameIe := "nameIe_example" // string |  (optional)
    nameNie := "nameNie_example" // string |  (optional)
    nameEmpty := "nameEmpty_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsList(context.Background()).Id(id).Name(name).Q(q).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsList`: UsersGroupsList200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **q** | **string** |  | 
 **idN** | **string** |  | 
 **idLte** | **string** |  | 
 **idLt** | **string** |  | 
 **idGte** | **string** |  | 
 **idGt** | **string** |  | 
 **nameN** | **string** |  | 
 **nameIc** | **string** |  | 
 **nameNic** | **string** |  | 
 **nameIew** | **string** |  | 
 **nameNiew** | **string** |  | 
 **nameIsw** | **string** |  | 
 **nameNisw** | **string** |  | 
 **nameIe** | **string** |  | 
 **nameNie** | **string** |  | 
 **nameEmpty** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**UsersGroupsList200Response**](UsersGroupsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsPartialUpdate

> Group UsersGroupsPartialUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.
    data := *openapiclient.NewGroup("Name_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsPartialUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsRead

> Group UsersGroupsRead(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsRead`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsUpdate

> Group UsersGroupsUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.
    data := *openapiclient.NewGroup("Name_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkDelete

> UsersPermissionsBulkDelete(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkPartialUpdate

> ObjectPermission UsersPermissionsBulkPartialUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableObjectPermission("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkPartialUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkUpdate

> ObjectPermission UsersPermissionsBulkUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableObjectPermission("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsCreate

> ObjectPermission UsersPermissionsCreate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableObjectPermission("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsCreate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsDelete

> UsersPermissionsDelete(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsList

> UsersPermissionsList200Response UsersPermissionsList(ctx).Id(id).Name(name).Enabled(enabled).ObjectTypes(objectTypes).UserId(userId).User(user).GroupId(groupId).Group(group).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).ObjectTypesN(objectTypesN).UserIdN(userIdN).UserN(userN).GroupIdN(groupIdN).GroupN(groupN).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    enabled := "enabled_example" // string |  (optional)
    objectTypes := "objectTypes_example" // string |  (optional)
    userId := "userId_example" // string |  (optional)
    user := "user_example" // string |  (optional)
    groupId := "groupId_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    idN := "idN_example" // string |  (optional)
    idLte := "idLte_example" // string |  (optional)
    idLt := "idLt_example" // string |  (optional)
    idGte := "idGte_example" // string |  (optional)
    idGt := "idGt_example" // string |  (optional)
    nameN := "nameN_example" // string |  (optional)
    nameIc := "nameIc_example" // string |  (optional)
    nameNic := "nameNic_example" // string |  (optional)
    nameIew := "nameIew_example" // string |  (optional)
    nameNiew := "nameNiew_example" // string |  (optional)
    nameIsw := "nameIsw_example" // string |  (optional)
    nameNisw := "nameNisw_example" // string |  (optional)
    nameIe := "nameIe_example" // string |  (optional)
    nameNie := "nameNie_example" // string |  (optional)
    nameEmpty := "nameEmpty_example" // string |  (optional)
    objectTypesN := "objectTypesN_example" // string |  (optional)
    userIdN := "userIdN_example" // string |  (optional)
    userN := "userN_example" // string |  (optional)
    groupIdN := "groupIdN_example" // string |  (optional)
    groupN := "groupN_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsList(context.Background()).Id(id).Name(name).Enabled(enabled).ObjectTypes(objectTypes).UserId(userId).User(user).GroupId(groupId).Group(group).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).ObjectTypesN(objectTypesN).UserIdN(userIdN).UserN(userN).GroupIdN(groupIdN).GroupN(groupN).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsList`: UsersPermissionsList200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **enabled** | **string** |  | 
 **objectTypes** | **string** |  | 
 **userId** | **string** |  | 
 **user** | **string** |  | 
 **groupId** | **string** |  | 
 **group** | **string** |  | 
 **idN** | **string** |  | 
 **idLte** | **string** |  | 
 **idLt** | **string** |  | 
 **idGte** | **string** |  | 
 **idGt** | **string** |  | 
 **nameN** | **string** |  | 
 **nameIc** | **string** |  | 
 **nameNic** | **string** |  | 
 **nameIew** | **string** |  | 
 **nameNiew** | **string** |  | 
 **nameIsw** | **string** |  | 
 **nameNisw** | **string** |  | 
 **nameIe** | **string** |  | 
 **nameNie** | **string** |  | 
 **nameEmpty** | **string** |  | 
 **objectTypesN** | **string** |  | 
 **userIdN** | **string** |  | 
 **userN** | **string** |  | 
 **groupIdN** | **string** |  | 
 **groupN** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**UsersPermissionsList200Response**](UsersPermissionsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsPartialUpdate

> ObjectPermission UsersPermissionsPartialUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.
    data := *openapiclient.NewWritableObjectPermission("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsPartialUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsRead

> ObjectPermission UsersPermissionsRead(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsRead`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsUpdate

> ObjectPermission UsersPermissionsUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.
    data := *openapiclient.NewWritableObjectPermission("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkDelete

> UsersUsersBulkDelete(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkPartialUpdate

> User UsersUsersBulkPartialUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableUser("Username_example", "Password_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkPartialUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkUpdate

> User UsersUsersBulkUpdate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableUser("Username_example", "Password_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersCreate

> User UsersUsersCreate(ctx).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    data := *openapiclient.NewWritableUser("Username_example", "Password_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersCreate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersDelete

> UsersUsersDelete(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersList

> UsersUsersList200Response UsersUsersList(ctx).Id(id).Username(username).FirstName(firstName).LastName(lastName).Email(email).IsStaff(isStaff).IsActive(isActive).Q(q).GroupId(groupId).Group(group).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).UsernameN(usernameN).UsernameIc(usernameIc).UsernameNic(usernameNic).UsernameIew(usernameIew).UsernameNiew(usernameNiew).UsernameIsw(usernameIsw).UsernameNisw(usernameNisw).UsernameIe(usernameIe).UsernameNie(usernameNie).UsernameEmpty(usernameEmpty).FirstNameN(firstNameN).FirstNameIc(firstNameIc).FirstNameNic(firstNameNic).FirstNameIew(firstNameIew).FirstNameNiew(firstNameNiew).FirstNameIsw(firstNameIsw).FirstNameNisw(firstNameNisw).FirstNameIe(firstNameIe).FirstNameNie(firstNameNie).FirstNameEmpty(firstNameEmpty).LastNameN(lastNameN).LastNameIc(lastNameIc).LastNameNic(lastNameNic).LastNameIew(lastNameIew).LastNameNiew(lastNameNiew).LastNameIsw(lastNameIsw).LastNameNisw(lastNameNisw).LastNameIe(lastNameIe).LastNameNie(lastNameNie).LastNameEmpty(lastNameEmpty).EmailN(emailN).EmailIc(emailIc).EmailNic(emailNic).EmailIew(emailIew).EmailNiew(emailNiew).EmailIsw(emailIsw).EmailNisw(emailNisw).EmailIe(emailIe).EmailNie(emailNie).EmailEmpty(emailEmpty).GroupIdN(groupIdN).GroupN(groupN).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string |  (optional)
    username := "username_example" // string |  (optional)
    firstName := "firstName_example" // string |  (optional)
    lastName := "lastName_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    isStaff := "isStaff_example" // string |  (optional)
    isActive := "isActive_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    groupId := "groupId_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    idN := "idN_example" // string |  (optional)
    idLte := "idLte_example" // string |  (optional)
    idLt := "idLt_example" // string |  (optional)
    idGte := "idGte_example" // string |  (optional)
    idGt := "idGt_example" // string |  (optional)
    usernameN := "usernameN_example" // string |  (optional)
    usernameIc := "usernameIc_example" // string |  (optional)
    usernameNic := "usernameNic_example" // string |  (optional)
    usernameIew := "usernameIew_example" // string |  (optional)
    usernameNiew := "usernameNiew_example" // string |  (optional)
    usernameIsw := "usernameIsw_example" // string |  (optional)
    usernameNisw := "usernameNisw_example" // string |  (optional)
    usernameIe := "usernameIe_example" // string |  (optional)
    usernameNie := "usernameNie_example" // string |  (optional)
    usernameEmpty := "usernameEmpty_example" // string |  (optional)
    firstNameN := "firstNameN_example" // string |  (optional)
    firstNameIc := "firstNameIc_example" // string |  (optional)
    firstNameNic := "firstNameNic_example" // string |  (optional)
    firstNameIew := "firstNameIew_example" // string |  (optional)
    firstNameNiew := "firstNameNiew_example" // string |  (optional)
    firstNameIsw := "firstNameIsw_example" // string |  (optional)
    firstNameNisw := "firstNameNisw_example" // string |  (optional)
    firstNameIe := "firstNameIe_example" // string |  (optional)
    firstNameNie := "firstNameNie_example" // string |  (optional)
    firstNameEmpty := "firstNameEmpty_example" // string |  (optional)
    lastNameN := "lastNameN_example" // string |  (optional)
    lastNameIc := "lastNameIc_example" // string |  (optional)
    lastNameNic := "lastNameNic_example" // string |  (optional)
    lastNameIew := "lastNameIew_example" // string |  (optional)
    lastNameNiew := "lastNameNiew_example" // string |  (optional)
    lastNameIsw := "lastNameIsw_example" // string |  (optional)
    lastNameNisw := "lastNameNisw_example" // string |  (optional)
    lastNameIe := "lastNameIe_example" // string |  (optional)
    lastNameNie := "lastNameNie_example" // string |  (optional)
    lastNameEmpty := "lastNameEmpty_example" // string |  (optional)
    emailN := "emailN_example" // string |  (optional)
    emailIc := "emailIc_example" // string |  (optional)
    emailNic := "emailNic_example" // string |  (optional)
    emailIew := "emailIew_example" // string |  (optional)
    emailNiew := "emailNiew_example" // string |  (optional)
    emailIsw := "emailIsw_example" // string |  (optional)
    emailNisw := "emailNisw_example" // string |  (optional)
    emailIe := "emailIe_example" // string |  (optional)
    emailNie := "emailNie_example" // string |  (optional)
    emailEmpty := "emailEmpty_example" // string |  (optional)
    groupIdN := "groupIdN_example" // string |  (optional)
    groupN := "groupN_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersList(context.Background()).Id(id).Username(username).FirstName(firstName).LastName(lastName).Email(email).IsStaff(isStaff).IsActive(isActive).Q(q).GroupId(groupId).Group(group).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).UsernameN(usernameN).UsernameIc(usernameIc).UsernameNic(usernameNic).UsernameIew(usernameIew).UsernameNiew(usernameNiew).UsernameIsw(usernameIsw).UsernameNisw(usernameNisw).UsernameIe(usernameIe).UsernameNie(usernameNie).UsernameEmpty(usernameEmpty).FirstNameN(firstNameN).FirstNameIc(firstNameIc).FirstNameNic(firstNameNic).FirstNameIew(firstNameIew).FirstNameNiew(firstNameNiew).FirstNameIsw(firstNameIsw).FirstNameNisw(firstNameNisw).FirstNameIe(firstNameIe).FirstNameNie(firstNameNie).FirstNameEmpty(firstNameEmpty).LastNameN(lastNameN).LastNameIc(lastNameIc).LastNameNic(lastNameNic).LastNameIew(lastNameIew).LastNameNiew(lastNameNiew).LastNameIsw(lastNameIsw).LastNameNisw(lastNameNisw).LastNameIe(lastNameIe).LastNameNie(lastNameNie).LastNameEmpty(lastNameEmpty).EmailN(emailN).EmailIc(emailIc).EmailNic(emailNic).EmailIew(emailIew).EmailNiew(emailNiew).EmailIsw(emailIsw).EmailNisw(emailNisw).EmailIe(emailIe).EmailNie(emailNie).EmailEmpty(emailEmpty).GroupIdN(groupIdN).GroupN(groupN).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersList`: UsersUsersList200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **username** | **string** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **email** | **string** |  | 
 **isStaff** | **string** |  | 
 **isActive** | **string** |  | 
 **q** | **string** |  | 
 **groupId** | **string** |  | 
 **group** | **string** |  | 
 **idN** | **string** |  | 
 **idLte** | **string** |  | 
 **idLt** | **string** |  | 
 **idGte** | **string** |  | 
 **idGt** | **string** |  | 
 **usernameN** | **string** |  | 
 **usernameIc** | **string** |  | 
 **usernameNic** | **string** |  | 
 **usernameIew** | **string** |  | 
 **usernameNiew** | **string** |  | 
 **usernameIsw** | **string** |  | 
 **usernameNisw** | **string** |  | 
 **usernameIe** | **string** |  | 
 **usernameNie** | **string** |  | 
 **usernameEmpty** | **string** |  | 
 **firstNameN** | **string** |  | 
 **firstNameIc** | **string** |  | 
 **firstNameNic** | **string** |  | 
 **firstNameIew** | **string** |  | 
 **firstNameNiew** | **string** |  | 
 **firstNameIsw** | **string** |  | 
 **firstNameNisw** | **string** |  | 
 **firstNameIe** | **string** |  | 
 **firstNameNie** | **string** |  | 
 **firstNameEmpty** | **string** |  | 
 **lastNameN** | **string** |  | 
 **lastNameIc** | **string** |  | 
 **lastNameNic** | **string** |  | 
 **lastNameIew** | **string** |  | 
 **lastNameNiew** | **string** |  | 
 **lastNameIsw** | **string** |  | 
 **lastNameNisw** | **string** |  | 
 **lastNameIe** | **string** |  | 
 **lastNameNie** | **string** |  | 
 **lastNameEmpty** | **string** |  | 
 **emailN** | **string** |  | 
 **emailIc** | **string** |  | 
 **emailNic** | **string** |  | 
 **emailIew** | **string** |  | 
 **emailNiew** | **string** |  | 
 **emailIsw** | **string** |  | 
 **emailNisw** | **string** |  | 
 **emailIe** | **string** |  | 
 **emailNie** | **string** |  | 
 **emailEmpty** | **string** |  | 
 **groupIdN** | **string** |  | 
 **groupN** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**UsersUsersList200Response**](UsersUsersList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersPartialUpdate

> User UsersUsersPartialUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.
    data := *openapiclient.NewWritableUser("Username_example", "Password_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersPartialUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersRead

> User UsersUsersRead(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersRead`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersUpdate

> User UsersUsersUpdate(ctx, id).Data(data).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.
    data := *openapiclient.NewWritableUser("Username_example", "Password_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

