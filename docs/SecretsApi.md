# \SecretsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretsGenerateRsaKeyPairList**](SecretsApi.md#SecretsGenerateRsaKeyPairList) | **Get** /secrets/generate-rsa-key-pair/ | This endpoint can be used to generate a new RSA key pair. The keys are returned in PEM format.
[**SecretsGetSessionKeyCreate**](SecretsApi.md#SecretsGetSessionKeyCreate) | **Post** /secrets/get-session-key/ | 
[**SecretsSecretRolesBulkDelete**](SecretsApi.md#SecretsSecretRolesBulkDelete) | **Delete** /secrets/secret-roles/ | 
[**SecretsSecretRolesBulkPartialUpdate**](SecretsApi.md#SecretsSecretRolesBulkPartialUpdate) | **Patch** /secrets/secret-roles/ | 
[**SecretsSecretRolesBulkUpdate**](SecretsApi.md#SecretsSecretRolesBulkUpdate) | **Put** /secrets/secret-roles/ | 
[**SecretsSecretRolesCreate**](SecretsApi.md#SecretsSecretRolesCreate) | **Post** /secrets/secret-roles/ | 
[**SecretsSecretRolesDelete**](SecretsApi.md#SecretsSecretRolesDelete) | **Delete** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesList**](SecretsApi.md#SecretsSecretRolesList) | **Get** /secrets/secret-roles/ | 
[**SecretsSecretRolesPartialUpdate**](SecretsApi.md#SecretsSecretRolesPartialUpdate) | **Patch** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesRead**](SecretsApi.md#SecretsSecretRolesRead) | **Get** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesUpdate**](SecretsApi.md#SecretsSecretRolesUpdate) | **Put** /secrets/secret-roles/{id}/ | 
[**SecretsSecretsBulkDelete**](SecretsApi.md#SecretsSecretsBulkDelete) | **Delete** /secrets/secrets/ | 
[**SecretsSecretsBulkPartialUpdate**](SecretsApi.md#SecretsSecretsBulkPartialUpdate) | **Patch** /secrets/secrets/ | 
[**SecretsSecretsBulkUpdate**](SecretsApi.md#SecretsSecretsBulkUpdate) | **Put** /secrets/secrets/ | 
[**SecretsSecretsCreate**](SecretsApi.md#SecretsSecretsCreate) | **Post** /secrets/secrets/ | 
[**SecretsSecretsDelete**](SecretsApi.md#SecretsSecretsDelete) | **Delete** /secrets/secrets/{id}/ | 
[**SecretsSecretsList**](SecretsApi.md#SecretsSecretsList) | **Get** /secrets/secrets/ | 
[**SecretsSecretsPartialUpdate**](SecretsApi.md#SecretsSecretsPartialUpdate) | **Patch** /secrets/secrets/{id}/ | 
[**SecretsSecretsRead**](SecretsApi.md#SecretsSecretsRead) | **Get** /secrets/secrets/{id}/ | 
[**SecretsSecretsUpdate**](SecretsApi.md#SecretsSecretsUpdate) | **Put** /secrets/secrets/{id}/ | 



## SecretsGenerateRsaKeyPairList

> SecretsGenerateRsaKeyPairList(ctx).Execute()

This endpoint can be used to generate a new RSA key pair. The keys are returned in PEM format.



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
    resp, r, err := apiClient.SecretsApi.SecretsGenerateRsaKeyPairList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsGenerateRsaKeyPairList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsGenerateRsaKeyPairListRequest struct via the builder pattern


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


## SecretsGetSessionKeyCreate

> SecretsGetSessionKeyCreate(ctx).Execute()





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
    resp, r, err := apiClient.SecretsApi.SecretsGetSessionKeyCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsGetSessionKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsGetSessionKeyCreateRequest struct via the builder pattern


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


## SecretsSecretRolesBulkDelete

> SecretsSecretRolesBulkDelete(ctx).Execute()



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
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesBulkDeleteRequest struct via the builder pattern


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


## SecretsSecretRolesBulkPartialUpdate

> SecretRole SecretsSecretRolesBulkPartialUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewSecretRole("Name_example", "Slug_example") // SecretRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesBulkPartialUpdate`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**SecretRole**](SecretRole.md) |  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesBulkUpdate

> SecretRole SecretsSecretRolesBulkUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewSecretRole("Name_example", "Slug_example") // SecretRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesBulkUpdate`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**SecretRole**](SecretRole.md) |  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesCreate

> SecretRole SecretsSecretRolesCreate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewSecretRole("Name_example", "Slug_example") // SecretRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesCreate`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**SecretRole**](SecretRole.md) |  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesDelete

> SecretsSecretRolesDelete(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesDeleteRequest struct via the builder pattern


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


## SecretsSecretRolesList

> SecretsSecretRolesList200Response SecretsSecretRolesList(ctx).Id(id).Name(name).Slug(slug).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesList(context.Background()).Id(id).Name(name).Slug(slug).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).SlugN(slugN).SlugIc(slugIc).SlugNic(slugNic).SlugIew(slugIew).SlugNiew(slugNiew).SlugIsw(slugIsw).SlugNisw(slugNisw).SlugIe(slugIe).SlugNie(slugNie).SlugEmpty(slugEmpty).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesList`: SecretsSecretRolesList200Response
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesListRequest struct via the builder pattern


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
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**SecretsSecretRolesList200Response**](SecretsSecretRolesList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesPartialUpdate

> SecretRole SecretsSecretRolesPartialUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret role.
    data := *openapiclient.NewSecretRole("Name_example", "Slug_example") // SecretRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesPartialUpdate`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**SecretRole**](SecretRole.md) |  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesRead

> SecretRole SecretsSecretRolesRead(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesRead`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretRolesUpdate

> SecretRole SecretsSecretRolesUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret role.
    data := *openapiclient.NewSecretRole("Name_example", "Slug_example") // SecretRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretRolesUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretRolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretRolesUpdate`: SecretRole
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretRolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**SecretRole**](SecretRole.md) |  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsBulkDelete

> SecretsSecretsBulkDelete(ctx).Execute()



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
    resp, r, err := apiClient.SecretsApi.SecretsSecretsBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsBulkDeleteRequest struct via the builder pattern


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


## SecretsSecretsBulkPartialUpdate

> Secret SecretsSecretsBulkPartialUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableSecret("AssignedObjectType_example", int32(123), int32(123), "Plaintext_example") // WritableSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsBulkPartialUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsBulkPartialUpdate`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableSecret**](WritableSecret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsBulkUpdate

> Secret SecretsSecretsBulkUpdate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableSecret("AssignedObjectType_example", int32(123), int32(123), "Plaintext_example") // WritableSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsBulkUpdate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsBulkUpdate`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableSecret**](WritableSecret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsCreate

> Secret SecretsSecretsCreate(ctx).Data(data).Execute()



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
    data := *openapiclient.NewWritableSecret("AssignedObjectType_example", int32(123), int32(123), "Plaintext_example") // WritableSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsCreate(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsCreate`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**WritableSecret**](WritableSecret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsDelete

> SecretsSecretsDelete(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsDeleteRequest struct via the builder pattern


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


## SecretsSecretsList

> SecretsSecretsList200Response SecretsSecretsList(ctx).Id(id).Name(name).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).RoleId(roleId).Role(role).Device(device).DeviceId(deviceId).VirtualMachine(virtualMachine).VirtualMachineId(virtualMachineId).Tag(tag).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).RoleIdN(roleIdN).RoleN(roleN).DeviceN(deviceN).DeviceIdN(deviceIdN).VirtualMachineN(virtualMachineN).VirtualMachineIdN(virtualMachineIdN).TagN(tagN).Limit(limit).Offset(offset).Execute()



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
    created := "created_example" // string |  (optional)
    createdGte := "createdGte_example" // string |  (optional)
    createdLte := "createdLte_example" // string |  (optional)
    lastUpdated := "lastUpdated_example" // string |  (optional)
    lastUpdatedGte := "lastUpdatedGte_example" // string |  (optional)
    lastUpdatedLte := "lastUpdatedLte_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    roleId := "roleId_example" // string |  (optional)
    role := "role_example" // string |  (optional)
    device := "device_example" // string |  (optional)
    deviceId := "deviceId_example" // string |  (optional)
    virtualMachine := "virtualMachine_example" // string |  (optional)
    virtualMachineId := "virtualMachineId_example" // string |  (optional)
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
    roleIdN := "roleIdN_example" // string |  (optional)
    roleN := "roleN_example" // string |  (optional)
    deviceN := "deviceN_example" // string |  (optional)
    deviceIdN := "deviceIdN_example" // string |  (optional)
    virtualMachineN := "virtualMachineN_example" // string |  (optional)
    virtualMachineIdN := "virtualMachineIdN_example" // string |  (optional)
    tagN := "tagN_example" // string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsList(context.Background()).Id(id).Name(name).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Q(q).RoleId(roleId).Role(role).Device(device).DeviceId(deviceId).VirtualMachine(virtualMachine).VirtualMachineId(virtualMachineId).Tag(tag).IdN(idN).IdLte(idLte).IdLt(idLt).IdGte(idGte).IdGt(idGt).NameN(nameN).NameIc(nameIc).NameNic(nameNic).NameIew(nameIew).NameNiew(nameNiew).NameIsw(nameIsw).NameNisw(nameNisw).NameIe(nameIe).NameNie(nameNie).NameEmpty(nameEmpty).RoleIdN(roleIdN).RoleN(roleN).DeviceN(deviceN).DeviceIdN(deviceIdN).VirtualMachineN(virtualMachineN).VirtualMachineIdN(virtualMachineIdN).TagN(tagN).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsList`: SecretsSecretsList200Response
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **created** | **string** |  | 
 **createdGte** | **string** |  | 
 **createdLte** | **string** |  | 
 **lastUpdated** | **string** |  | 
 **lastUpdatedGte** | **string** |  | 
 **lastUpdatedLte** | **string** |  | 
 **q** | **string** |  | 
 **roleId** | **string** |  | 
 **role** | **string** |  | 
 **device** | **string** |  | 
 **deviceId** | **string** |  | 
 **virtualMachine** | **string** |  | 
 **virtualMachineId** | **string** |  | 
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
 **roleIdN** | **string** |  | 
 **roleN** | **string** |  | 
 **deviceN** | **string** |  | 
 **deviceIdN** | **string** |  | 
 **virtualMachineN** | **string** |  | 
 **virtualMachineIdN** | **string** |  | 
 **tagN** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**SecretsSecretsList200Response**](SecretsSecretsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsPartialUpdate

> Secret SecretsSecretsPartialUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret.
    data := *openapiclient.NewWritableSecret("AssignedObjectType_example", int32(123), int32(123), "Plaintext_example") // WritableSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsPartialUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsPartialUpdate`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableSecret**](WritableSecret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsRead

> Secret SecretsSecretsRead(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsRead`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsSecretsUpdate

> Secret SecretsSecretsUpdate(ctx, id).Data(data).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this secret.
    data := *openapiclient.NewWritableSecret("AssignedObjectType_example", int32(123), int32(123), "Plaintext_example") // WritableSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.SecretsSecretsUpdate(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.SecretsSecretsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretsSecretsUpdate`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.SecretsSecretsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsSecretsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WritableSecret**](WritableSecret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

