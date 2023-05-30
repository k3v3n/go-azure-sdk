
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/usages` Documentation

The `usages` SDK allows for interaction with the Azure Resource Manager Service `cognitive` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/usages"
```


### Client Initialization

```go
client := usages.NewUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsagesClient.List`

```go
ctx := context.TODO()
id := usages.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.List(ctx, id, usages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, usages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```