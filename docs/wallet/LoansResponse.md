# LoansResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Identifier for the vault. | [optional] 
**user_id** | **int** | Identifier for the user. | 
**liquidation_price** | [**Currency**](Currency.md) |  | 
**collateral** | [**Currency**](Currency.md) |  | 
**collateralization_ratio** | **str** | The collateralization ratio of the vault. | 
**loan_to_value** | **str** | The loan to value of the vault. | 
**debt** | [**Currency**](Currency.md) |  | 
**principal** | [**Currency**](Currency.md) |  | 
**scaled_debt** | [**Currency**](Currency.md) |  | 
**plan** | [**LoanPlan**](LoanPlan.md) |  | 
**state** | **dict(str, str)** |  | 
**updated_collateral_token_balance** | [**Currency**](Currency.md) |  | [optional] 
**updated_debt_token_balance** | [**Currency**](Currency.md) |  | [optional] 
**terms_and_conditions** | [**BulletContent**](BulletContent.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


