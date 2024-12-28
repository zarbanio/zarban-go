# Documentation for Zarban API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://api.zarban.io*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AccountsApi* | [**getAccountByAddress**](Apis/AccountsApi.md#getaccountbyaddress) | **GET** /v2/accounts/{address} | Get account by address |
| *AddressesApi* | [**getAllAddresses**](Apis/AddressesApi.md#getalladdresses) | **GET** /v2/addresses | Get all addresses |
| *BorrowsApi* | [**getUserBorrows**](Apis/BorrowsApi.md#getuserborrows) | **GET** /v2/lendingpool/borrows | Get user borrows of lendingpool |
| *CollectorApi* | [**getCollectorData**](Apis/CollectorApi.md#getcollectordata) | **GET** /v2/stats | Get collector data |
| *DepositsApi* | [**getUserDeposits**](Apis/DepositsApi.md#getuserdeposits) | **GET** /v2/lendingpool/deposits | Get user deposits of Lendingpool |
| *IlksApi* | [**getAllIlks**](Apis/IlksApi.md#getallilks) | **GET** /v2/ilks | Get all Ilks |
*IlksApi* | [**getIlkByName**](Apis/IlksApi.md#getilkbyname) | **GET** /v2/ilks/{name} | Get Ilk by name |
| *LendingPoolApi* | [**createLendingPoolBorrow**](Apis/LendingPoolApi.md#createlendingpoolborrow) | **POST** /v2/lendingpool/tx/borrow | Borrow from lending pool |
*LendingPoolApi* | [**createLendingPoolDeposit**](Apis/LendingPoolApi.md#createlendingpooldeposit) | **POST** /v2/lendingpool/tx/deposit | Deposit to lending pool |
*LendingPoolApi* | [**createLendingPoolRepay**](Apis/LendingPoolApi.md#createlendingpoolrepay) | **POST** /v2/lendingpool/tx/repay | Repay to lending pool |
*LendingPoolApi* | [**createLendingPoolWithdraw**](Apis/LendingPoolApi.md#createlendingpoolwithdraw) | **POST** /v2/lendingpool/tx/withdraw | Withdraw from lending pool |
*LendingPoolApi* | [**fetchReserveDataByAsset**](Apis/LendingPoolApi.md#fetchreservedatabyasset) | **GET** /v2/lendingpool/reserves | Fetch Reserve Data By Asset |
*LendingPoolApi* | [**getUserBorrows**](Apis/LendingPoolApi.md#getuserborrows) | **GET** /v2/lendingpool/borrows | Get user borrows of lendingpool |
*LendingPoolApi* | [**getUserDeposits**](Apis/LendingPoolApi.md#getuserdeposits) | **GET** /v2/lendingpool/deposits | Get user deposits of Lendingpool |
*LendingPoolApi* | [**setLendingPoolAssetCollateral**](Apis/LendingPoolApi.md#setlendingpoolassetcollateral) | **POST** /v2/lendingpool/tx/useassetascollateral | Enable/Disable asset as collateral |
| *LogsApi* | [**getLogsByTransactionHash**](Apis/LogsApi.md#getlogsbytransactionhash) | **GET** /v2/logs/{txHash} | Get raw and decoded logs by transaction hash |
| *OrdersApi* | [**getUnfilledOrders**](Apis/OrdersApi.md#getunfilledorders) | **GET** /v2/orders | Fetch Unfilled Orders |
*OrdersApi* | [**syncOrder**](Apis/OrdersApi.md#syncorder) | **POST** /v2/orders/sync | Updates Order Entity |
| *PermitApi* | [**getSingleTokenPermit**](Apis/PermitApi.md#getsingletokenpermit) | **GET** /v2/permit/single | Get permit for single token |
| *PointsApi* | [**getScoreboard**](Apis/PointsApi.md#getscoreboard) | **GET** /v2/points/scoreboard | Get scoreboard |
| *PricesApi* | [**listPrices**](Apis/PricesApi.md#listprices) | **GET** /v2/prices | List prices based on query parameters |
| *StableCoinSystemApi* | [**approveAndJoinZarTransaction**](Apis/StableCoinSystemApi.md#approveandjoinzartransaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarjoin | approve and join ZAR token into Vat contract |
*StableCoinSystemApi* | [**createStableCoinVault**](Apis/StableCoinSystemApi.md#createstablecoinvault) | **POST** /v2/stablecoinsystem/tx/createvault | Create vault |
*StableCoinSystemApi* | [**depositStableCoinCollateral**](Apis/StableCoinSystemApi.md#depositstablecoincollateral) | **POST** /v2/stablecoinsystem/tx/depositcollateral | Deposit collateral |
*StableCoinSystemApi* | [**exitGemTransaction**](Apis/StableCoinSystemApi.md#exitgemtransaction) | **POST** /v2/stablecoinsystem/auctions/tx/gemexit | exit Gem token (which can be used as collateral) from Vat contract |
*StableCoinSystemApi* | [**exitZarTransaction**](Apis/StableCoinSystemApi.md#exitzartransaction) | **POST** /v2/stablecoinsystem/auctions/tx/zarexit | exit ZAR token from Vat contract |
*StableCoinSystemApi* | [**getAllIlks**](Apis/StableCoinSystemApi.md#getallilks) | **GET** /v2/ilks | Get all Ilks |
*StableCoinSystemApi* | [**getCollectorData**](Apis/StableCoinSystemApi.md#getcollectordata) | **GET** /v2/stats | Get collector data |
*StableCoinSystemApi* | [**getIlkByName**](Apis/StableCoinSystemApi.md#getilkbyname) | **GET** /v2/ilks/{name} | Get Ilk by name |
*StableCoinSystemApi* | [**getVaultById**](Apis/StableCoinSystemApi.md#getvaultbyid) | **GET** /v2/vaults/{id} | Get a vault by ID |
*StableCoinSystemApi* | [**getVaultEventsById**](Apis/StableCoinSystemApi.md#getvaulteventsbyid) | **GET** /v2/vaults/{id}/events | Get vault events by ID |
*StableCoinSystemApi* | [**getVaultsByOwner**](Apis/StableCoinSystemApi.md#getvaultsbyowner) | **GET** /v2/vaults | Get vaults by owner query |
*StableCoinSystemApi* | [**liquidateVaultTransaction**](Apis/StableCoinSystemApi.md#liquidatevaulttransaction) | **POST** /v2/stablecoinsystem/tx/bark | liquidate a vault |
*StableCoinSystemApi* | [**mintZarTransaction**](Apis/StableCoinSystemApi.md#mintzartransaction) | **POST** /v2/stablecoinsystem/tx/mintzar | Mint ZAR |
*StableCoinSystemApi* | [**repayZarTransaction**](Apis/StableCoinSystemApi.md#repayzartransaction) | **POST** /v2/stablecoinsystem/tx/repayzar | Repay ZAR |
*StableCoinSystemApi* | [**resetAuctionTransaction**](Apis/StableCoinSystemApi.md#resetauctiontransaction) | **POST** /v2/stablecoinsystem/auctions/tx/redo | reset a auction |
*StableCoinSystemApi* | [**takeAuctionTransaction**](Apis/StableCoinSystemApi.md#takeauctiontransaction) | **POST** /v2/stablecoinsystem/auctions/tx/take | take a auction |
*StableCoinSystemApi* | [**withdrawCollateralTransaction**](Apis/StableCoinSystemApi.md#withdrawcollateraltransaction) | **POST** /v2/stablecoinsystem/tx/withdrawcollateral | Withdraw collateral |
| *StakingApi* | [**collectStakingReward**](Apis/StakingApi.md#collectstakingreward) | **POST** /v2/staking/tx/collectreward | Collect staking reward |
*StakingApi* | [**getStakingPlans**](Apis/StakingApi.md#getstakingplans) | **GET** /v2/staking/plans | Get staking plans |
*StakingApi* | [**getUserStakingStats**](Apis/StakingApi.md#getuserstakingstats) | **GET** /v2/staking/stats | Get user staking stats |
*StakingApi* | [**stakeToStakingContract**](Apis/StakingApi.md#staketostakingcontract) | **POST** /v2/staking/tx/stake | Stake to staking contract |
*StakingApi* | [**withdrawStakedAsset**](Apis/StakingApi.md#withdrawstakedasset) | **POST** /v2/staking/tx/withdraw | Withdraw staked asset |
| *SwapApi* | [**getSwapQuote**](Apis/SwapApi.md#getswapquote) | **POST** /v2/swap/quote | Get a quote for a swap |
| *VaultsApi* | [**getVaultById**](Apis/VaultsApi.md#getvaultbyid) | **GET** /v2/vaults/{id} | Get a vault by ID |
*VaultsApi* | [**getVaultEventsById**](Apis/VaultsApi.md#getvaulteventsbyid) | **GET** /v2/vaults/{id}/events | Get vault events by ID |
*VaultsApi* | [**getVaultsByOwner**](Apis/VaultsApi.md#getvaultsbyowner) | **GET** /v2/vaults | Get vaults by owner query |
| *WebsocketApi* | [**getUnfilledOrdersWebsocket**](Apis/WebsocketApi.md#getunfilledorderswebsocket) | **GET** /v2/ws | Websocket Upgrade |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Account](./Models/Account.md)
 - [AccountLendingpoolSummary](./Models/AccountLendingpoolSummary.md)
 - [AccountStablecoinSystemSummary](./Models/AccountStablecoinSystemSummary.md)
 - [AccountStakingSummary](./Models/AccountStakingSummary.md)
 - [Address](./Models/Address.md)
 - [AddressResponse](./Models/AddressResponse.md)
 - [Balance](./Models/Balance.md)
 - [BasicEvent](./Models/BasicEvent.md)
 - [ChainActivity](./Models/ChainActivity.md)
 - [ChainActivityStep](./Models/ChainActivityStep.md)
 - [ChainActivityStepData](./Models/ChainActivityStepData.md)
 - [DutchAmount](./Models/DutchAmount.md)
 - [EIP712SignRequest](./Models/EIP712SignRequest.md)
 - [Error](./Models/Error.md)
 - [ErrorMessage](./Models/ErrorMessage.md)
 - [EventDetailsResponse](./Models/EventDetailsResponse.md)
 - [EventDomain](./Models/EventDomain.md)
 - [EventName](./Models/EventName.md)
 - [EventType](./Models/EventType.md)
 - [ExtendedEvent](./Models/ExtendedEvent.md)
 - [FormattedReserveData](./Models/FormattedReserveData.md)
 - [Ilk](./Models/Ilk.md)
 - [IlksResponse](./Models/IlksResponse.md)
 - [LendingpoolBorrow](./Models/LendingpoolBorrow.md)
 - [LendingpoolBorrowTxRequest](./Models/LendingpoolBorrowTxRequest.md)
 - [LendingpoolBorrowTxResponse](./Models/LendingpoolBorrowTxResponse.md)
 - [LendingpoolDeposit](./Models/LendingpoolDeposit.md)
 - [LendingpoolDepositTxRequest](./Models/LendingpoolDepositTxRequest.md)
 - [LendingpoolDepositTxResponse](./Models/LendingpoolDepositTxResponse.md)
 - [LendingpoolRepayTxRequest](./Models/LendingpoolRepayTxRequest.md)
 - [LendingpoolRepayTxResponse](./Models/LendingpoolRepayTxResponse.md)
 - [LendingpoolStats](./Models/LendingpoolStats.md)
 - [LendingpoolTxResponse](./Models/LendingpoolTxResponse.md)
 - [LendingpoolUseAssetAsCollateralTxRequest](./Models/LendingpoolUseAssetAsCollateralTxRequest.md)
 - [LendingpoolUseAssetAsCollateralTxResponse](./Models/LendingpoolUseAssetAsCollateralTxResponse.md)
 - [LendingpoolWithdrawTxRequest](./Models/LendingpoolWithdrawTxRequest.md)
 - [LendingpoolWithdrawTxResponse](./Models/LendingpoolWithdrawTxResponse.md)
 - [Log](./Models/Log.md)
 - [MethodParameters](./Models/MethodParameters.md)
 - [Order](./Models/Order.md)
 - [OrderInfo](./Models/OrderInfo.md)
 - [OrderResponse](./Models/OrderResponse.md)
 - [OrderType](./Models/OrderType.md)
 - [PermitSingle](./Models/PermitSingle.md)
 - [PersonalSignRequest](./Models/PersonalSignRequest.md)
 - [PreparedTx](./Models/PreparedTx.md)
 - [Price](./Models/Price.md)
 - [PriceListResponse](./Models/PriceListResponse.md)
 - [QuoteRequest](./Models/QuoteRequest.md)
 - [QuoteRequestOptions](./Models/QuoteRequestOptions.md)
 - [QuoteResponse](./Models/QuoteResponse.md)
 - [RawDutchAmount](./Models/RawDutchAmount.md)
 - [RouteItem](./Models/RouteItem.md)
 - [Scoreboard](./Models/Scoreboard.md)
 - [ScoreboardItem](./Models/ScoreboardItem.md)
 - [StablecoinSystemBarkTxRequest](./Models/StablecoinSystemBarkTxRequest.md)
 - [StablecoinSystemCreateVaultTxRequest](./Models/StablecoinSystemCreateVaultTxRequest.md)
 - [StablecoinSystemDepositCollateralTxRequest](./Models/StablecoinSystemDepositCollateralTxRequest.md)
 - [StablecoinSystemGemexitTxRequest](./Models/StablecoinSystemGemexitTxRequest.md)
 - [StablecoinSystemMintZarTxRequest](./Models/StablecoinSystemMintZarTxRequest.md)
 - [StablecoinSystemRedoTxRequest](./Models/StablecoinSystemRedoTxRequest.md)
 - [StablecoinSystemRepayZarTxRequest](./Models/StablecoinSystemRepayZarTxRequest.md)
 - [StablecoinSystemStats](./Models/StablecoinSystemStats.md)
 - [StablecoinSystemTakeTxRequest](./Models/StablecoinSystemTakeTxRequest.md)
 - [StablecoinSystemWithdrawCollateralTxRequest](./Models/StablecoinSystemWithdrawCollateralTxRequest.md)
 - [StablecoinSystemZarexitTxRequest](./Models/StablecoinSystemZarexitTxRequest.md)
 - [StablecoinSystemZarjoinTxRequest](./Models/StablecoinSystemZarjoinTxRequest.md)
 - [StakeBalance](./Models/StakeBalance.md)
 - [StakePlan](./Models/StakePlan.md)
 - [StakePlansResponse](./Models/StakePlansResponse.md)
 - [StakingCollectRewardTxRequest](./Models/StakingCollectRewardTxRequest.md)
 - [StakingCollectRewardTxResponse](./Models/StakingCollectRewardTxResponse.md)
 - [StakingStakeTxRequest](./Models/StakingStakeTxRequest.md)
 - [StakingStakeTxResponse](./Models/StakingStakeTxResponse.md)
 - [StakingWithdrawTxRequest](./Models/StakingWithdrawTxRequest.md)
 - [StakingWithdrawTxResponse](./Models/StakingWithdrawTxResponse.md)
 - [Stats](./Models/Stats.md)
 - [Symbol](./Models/Symbol.md)
 - [SystemBadDebt](./Models/SystemBadDebt.md)
 - [SystemDebt](./Models/SystemDebt.md)
 - [SystemDebtCeiling](./Models/SystemDebtCeiling.md)
 - [SystemSurplus](./Models/SystemSurplus.md)
 - [SystemSurplusBuffer](./Models/SystemSurplusBuffer.md)
 - [SystemSurplusLotSize](./Models/SystemSurplusLotSize.md)
 - [TimeRange](./Models/TimeRange.md)
 - [Timestamp](./Models/Timestamp.md)
 - [Token](./Models/Token.md)
 - [TradeType](./Models/TradeType.md)
 - [Type](./Models/Type.md)
 - [TypedData](./Models/TypedData.md)
 - [TypedDataDomain](./Models/TypedDataDomain.md)
 - [UnclaimedReward](./Models/UnclaimedReward.md)
 - [UpdateOrderRequest](./Models/UpdateOrderRequest.md)
 - [UserBorrowsResponse](./Models/UserBorrowsResponse.md)
 - [UserDepositsResponse](./Models/UserDepositsResponse.md)
 - [UserError](./Models/UserError.md)
 - [UserStake](./Models/UserStake.md)
 - [UserStakesResponse](./Models/UserStakesResponse.md)
 - [V3PoolInRoute](./Models/V3PoolInRoute.md)
 - [Vault](./Models/Vault.md)
 - [VaultEvent](./Models/VaultEvent.md)
 - [VaultEventExtension](./Models/VaultEventExtension.md)
 - [VaultEventItem](./Models/VaultEventItem.md)
 - [VaultEventsResponse](./Models/VaultEventsResponse.md)
 - [VaultsResponse](./Models/VaultsResponse.md)
 - [WalletBalance](./Models/WalletBalance.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
