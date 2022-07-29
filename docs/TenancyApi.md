# \TenancyApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenancyTenantGroupsBulkDelete**](TenancyApi.md#TenancyTenantGroupsBulkDelete) | **Delete** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkPartialUpdate**](TenancyApi.md#TenancyTenantGroupsBulkPartialUpdate) | **Patch** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkUpdate**](TenancyApi.md#TenancyTenantGroupsBulkUpdate) | **Put** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsCreate**](TenancyApi.md#TenancyTenantGroupsCreate) | **Post** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsDelete**](TenancyApi.md#TenancyTenantGroupsDelete) | **Delete** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsList**](TenancyApi.md#TenancyTenantGroupsList) | **Get** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsPartialUpdate**](TenancyApi.md#TenancyTenantGroupsPartialUpdate) | **Patch** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsRead**](TenancyApi.md#TenancyTenantGroupsRead) | **Get** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsUpdate**](TenancyApi.md#TenancyTenantGroupsUpdate) | **Put** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantsBulkDelete**](TenancyApi.md#TenancyTenantsBulkDelete) | **Delete** /tenancy/tenants/ | 
[**TenancyTenantsBulkPartialUpdate**](TenancyApi.md#TenancyTenantsBulkPartialUpdate) | **Patch** /tenancy/tenants/ | 
[**TenancyTenantsBulkUpdate**](TenancyApi.md#TenancyTenantsBulkUpdate) | **Put** /tenancy/tenants/ | 
[**TenancyTenantsCreate**](TenancyApi.md#TenancyTenantsCreate) | **Post** /tenancy/tenants/ | 
[**TenancyTenantsDelete**](TenancyApi.md#TenancyTenantsDelete) | **Delete** /tenancy/tenants/{id}/ | 
[**TenancyTenantsList**](TenancyApi.md#TenancyTenantsList) | **Get** /tenancy/tenants/ | 
[**TenancyTenantsPartialUpdate**](TenancyApi.md#TenancyTenantsPartialUpdate) | **Patch** /tenancy/tenants/{id}/ | 
[**TenancyTenantsRead**](TenancyApi.md#TenancyTenantsRead) | **Get** /tenancy/tenants/{id}/ | 
[**TenancyTenantsUpdate**](TenancyApi.md#TenancyTenantsUpdate) | **Put** /tenancy/tenants/{id}/ | 



## TenancyTenantGroupsBulkDelete

> TenancyTenantGroupsBulkDelete(ctx).Execute()



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
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkDeleteRequest struct via the builder pattern


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


## TenancyTenantGroupsBulkPartialUpdate

> TenantGroup TenancyTenantGroupsBulkPartialUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenantGroup("Name_example", "Slug_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsBulkPartialUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsBulkUpdate

> TenantGroup TenancyTenantGroupsBulkUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenantGroup("Name_example", "Slug_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsBulkUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsCreate

> TenantGroup TenancyTenantGroupsCreate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenantGroup("Name_example", "Slug_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsCreate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsDelete

> TenancyTenantGroupsDelete(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsDeleteRequest struct via the builder pattern


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


## TenancyTenantGroupsList

> TenancyTenantGroupsList200Response TenancyTenantGroupsList(ctx).Id(id).Name(name).Slug(slug).Description(description).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).ParentId(parentId).Parent(parent).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).DescriptionN(descriptionN).DescriptionIc(descriptionIc).DescriptionNic(descriptionNic).DescriptionIew(descriptionIew).DescriptionNiew(descriptionNiew).DescriptionIsw(descriptionIsw).DescriptionNisw(descriptionNisw).DescriptionIe(descriptionIe).DescriptionNie(descriptionNie).DescriptionEmpty(descriptionEmpty).ParentIdN(parentIdN).ParentN(parentN).Limit(limit).Offset(offset).Execute()



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
    slug := "slug_example" // string |  (optional)
    description := "description_example" // string |  (optional)
    created := "created_example" // string |  (optional)
    createdGte := "createdGte_example" // string |  (optional)
    createdLte := "createdLte_example" // string |  (optional)
    lastUpdated := "lastUpdated_example" // string |  (optional)
    lastUpdatedGte := "lastUpdatedGte_example" // string |  (optional)
    lastUpdatedLte := "lastUpdatedLte_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    parent := "parent_example" // string |  (optional)
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
    slugN := "slugN_example" // string |  (optional)
    slugIc := "slugIc_example" // string |  (optional)
    slugNic := "slugNic_example" // string |  (optional)
    slugIew := "slugIew_example" // string |  (optional)
    slugNiew := "slugNiew_example" // string |  (optional)
    slugIsw := "slugIsw_example" // string |  (optional)
    slugNisw := "slugNisw_example" // string |  (optional)
    slugIe := "slugIe_example" // string |  (optional)
    slugNie := "slugNie_example" // string |  (optional)
    slugEmpty := "slugEmpty_example" // string |  (optional)
    descriptionN := "descriptionN_example" // string |  (optional)
    descriptionIc := "descriptionIc_example" // string |  (optional)
    descriptionNic := "descriptionNic_example" // string |  (optional)
    descriptionIew := "descriptionIew_example" // string |  (optional)
    descriptionNiew := "descriptionNiew_example" // string |  (optional)
    descriptionIsw := "descriptionIsw_example" // string |  (optional)
    descriptionNisw := "descriptionNisw_example" // string |  (optional)
    descriptionIe := "descriptionIe_example" // string |  (optional)
    descriptionNie := "descriptionNie_example" // string |  (optional)
    descriptionEmpty := "descriptionEmpty_example" // string |  (optional)
    parentIdN := "parentIdN_example" // string |  (optional)
    parentN := "parentN_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsList(context.Background()).Id(id).Name(name).Slug(slug).Description(description).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).ParentId(parentId).Parent(parent).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).DescriptionN(descriptionN).DescriptionIc(descriptionIc).DescriptionNic(descriptionNic).DescriptionIew(descriptionIew).DescriptionNiew(descriptionNiew).DescriptionIsw(descriptionIsw).DescriptionNisw(descriptionNisw).DescriptionIe(descriptionIe).DescriptionNie(descriptionNie).DescriptionEmpty(descriptionEmpty).ParentIdN(parentIdN).ParentN(parentN).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsList`: TenancyTenantGroupsList200Response
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **slug** | **string** |  | 
 **description** | **string** |  | 
 **created** | **string** |  | 
 **createdGte** | **string** |  | 
 **createdLte** | **string** |  | 
 **lastUpdated** | **string** |  | 
 **lastUpdatedGte** | **string** |  | 
 **lastUpdatedLte** | **string** |  | 
 **q** | **string** |  | 
 **parentId** | **string** |  | 
 **parent** | **string** |  | 
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
 **slugN** | **string** |  | 
 **slugIc** | **string** |  | 
 **slugNic** | **string** |  | 
 **slugIew** | **string** |  | 
 **slugNiew** | **string** |  | 
 **slugIsw** | **string** |  | 
 **slugNisw** | **string** |  | 
 **slugIe** | **string** |  | 
 **slugNie** | **string** |  | 
 **slugEmpty** | **string** |  | 
 **descriptionN** | **string** |  | 
 **descriptionIc** | **string** |  | 
 **descriptionNic** | **string** |  | 
 **descriptionIew** | **string** |  | 
 **descriptionNiew** | **string** |  | 
 **descriptionIsw** | **string** |  | 
 **descriptionNisw** | **string** |  | 
 **descriptionIe** | **string** |  | 
 **descriptionNie** | **string** |  | 
 **descriptionEmpty** | **string** |  | 
 **parentIdN** | **string** |  | 
 **parentN** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**TenancyTenantGroupsList200Response**](TenancyTenantGroupsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsPartialUpdate

> TenantGroup TenancyTenantGroupsPartialUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant group.
    data := *openapiclient.NewWritableTenantGroup("Name_example", "Slug_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsPartialUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsRead

> TenantGroup TenancyTenantGroupsRead(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsRead`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsUpdate

> TenantGroup TenancyTenantGroupsUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant group.
    data := *openapiclient.NewWritableTenantGroup("Name_example", "Slug_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkDelete

> TenancyTenantsBulkDelete(ctx).Execute()



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
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkDeleteRequest struct via the builder pattern


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


## TenancyTenantsBulkPartialUpdate

> Tenant TenancyTenantsBulkPartialUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenant("Name_example", "Slug_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsBulkPartialUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkUpdate

> Tenant TenancyTenantsBulkUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenant("Name_example", "Slug_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsBulkUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsCreate

> Tenant TenancyTenantsCreate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableTenant("Name_example", "Slug_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsCreate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsDelete

> TenancyTenantsDelete(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsDeleteRequest struct via the builder pattern


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


## TenancyTenantsList

> TenancyTenantsList200Response TenancyTenantsList(ctx).Id(id).Name(name).Slug(slug).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).GroupId(groupId).Group(group).Tag(tag).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).GroupIdN(groupIdN).GroupN(groupN).TagN(tagN).Limit(limit).Offset(offset).Execute()



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
    slug := "slug_example" // string |  (optional)
    created := "created_example" // string |  (optional)
    createdGte := "createdGte_example" // string |  (optional)
    createdLte := "createdLte_example" // string |  (optional)
    lastUpdated := "lastUpdated_example" // string |  (optional)
    lastUpdatedGte := "lastUpdatedGte_example" // string |  (optional)
    lastUpdatedLte := "lastUpdatedLte_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    groupId := "groupId_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    tag := "tag_example" // string |  (optional)
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
    slugN := "slugN_example" // string |  (optional)
    slugIc := "slugIc_example" // string |  (optional)
    slugNic := "slugNic_example" // string |  (optional)
    slugIew := "slugIew_example" // string |  (optional)
    slugNiew := "slugNiew_example" // string |  (optional)
    slugIsw := "slugIsw_example" // string |  (optional)
    slugNisw := "slugNisw_example" // string |  (optional)
    slugIe := "slugIe_example" // string |  (optional)
    slugNie := "slugNie_example" // string |  (optional)
    slugEmpty := "slugEmpty_example" // string |  (optional)
    groupIdN := "groupIdN_example" // string |  (optional)
    groupN := "groupN_example" // string |  (optional)
    tagN := "tagN_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsList(context.Background()).Id(id).Name(name).Slug(slug).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).GroupId(groupId).Group(group).Tag(tag).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).GroupIdN(groupIdN).GroupN(groupN).TagN(tagN).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsList`: TenancyTenantsList200Response
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **slug** | **string** |  | 
 **created** | **string** |  | 
 **createdGte** | **string** |  | 
 **createdLte** | **string** |  | 
 **lastUpdated** | **string** |  | 
 **lastUpdatedGte** | **string** |  | 
 **lastUpdatedLte** | **string** |  | 
 **q** | **string** |  | 
 **groupId** | **string** |  | 
 **group** | **string** |  | 
 **tag** | **string** |  | 
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
 **slugN** | **string** |  | 
 **slugIc** | **string** |  | 
 **slugNic** | **string** |  | 
 **slugIew** | **string** |  | 
 **slugNiew** | **string** |  | 
 **slugIsw** | **string** |  | 
 **slugNisw** | **string** |  | 
 **slugIe** | **string** |  | 
 **slugNie** | **string** |  | 
 **slugEmpty** | **string** |  | 
 **groupIdN** | **string** |  | 
 **groupN** | **string** |  | 
 **tagN** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**TenancyTenantsList200Response**](TenancyTenantsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsPartialUpdate

> Tenant TenancyTenantsPartialUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant.
    data := *openapiclient.NewWritableTenant("Name_example", "Slug_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsPartialUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsRead

> Tenant TenancyTenantsRead(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsRead`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsUpdate

> Tenant TenancyTenantsUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this tenant.
    data := *openapiclient.NewWritableTenant("Name_example", "Slug_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

