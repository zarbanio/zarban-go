# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier of the transaction | 
**time** | [**Timestamp**](Timestamp.md) |  | 
**type** | [**TransactionType**](TransactionType.md) |  | 
**_from** | **str** | The ID of the user from whom the transaction originated | 
**to** | **str** | The ID of the user to whom the transaction is directed | 
**symbol** | [**Symbol**](Symbol.md) |  | 
**amount** | [**Currency**](Currency.md) |  | 
**direction** | **str** | Equal to \&quot;Inbound\&quot; if the transaction is directed to the user, and \&quot;Outbound\&quot; if the transaction is directed from the user. | 
**external_transaction** | [**ExternalTransaction**](ExternalTransaction.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


