# Transaction
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **Long** | Unique identifier of the transaction | [default to null] |
| **time** | [**Timestamp**](Timestamp.md) |  | [default to null] |
| **type** | [**TransactionType**](TransactionType.md) |  | [default to null] |
| **from** | **String** | The ID of the user from whom the transaction originated | [default to null] |
| **to** | **String** | The ID of the user to whom the transaction is directed | [default to null] |
| **symbol** | [**Symbol**](Symbol.md) |  | [default to null] |
| **amount** | [**Currency**](Currency.md) |  | [default to null] |
| **direction** | **String** | Equal to \&quot;Inbound\&quot; if the transaction is directed to the user, and \&quot;Outbound\&quot; if the transaction is directed from the user. | [default to null] |
| **externalTransaction** | [**ExternalTransaction**](ExternalTransaction.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

