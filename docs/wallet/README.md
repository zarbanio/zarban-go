# Documentation for Zarban Wallet API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://wapi.zarban.io*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AuthApi* | [**authenticateWithTelegram**](Apis/AuthApi.md#authenticatewithtelegram) | **POST** /auth/telegram | Authenticate with Telegram |
*AuthApi* | [**generateJwtToken**](Apis/AuthApi.md#generatejwttoken) | **GET** /auth/token | Generate a JWT token |
*AuthApi* | [**getOtp**](Apis/AuthApi.md#getotp) | **GET** /auth/otp | Get OTP |
*AuthApi* | [**loginWithEmailAndPassword**](Apis/AuthApi.md#loginwithemailandpassword) | **POST** /auth/login | Login with email and password |
*AuthApi* | [**signupWithEmailAndPassword**](Apis/AuthApi.md#signupwithemailandpassword) | **POST** /auth/signup | signup with email and password |
*AuthApi* | [**verifyUserEmail**](Apis/AuthApi.md#verifyuseremail) | **GET** /verify-email | Verify email |
| *BalanceApi* | [**getBalanceBySymbol**](Apis/BalanceApi.md#getbalancebysymbol) | **GET** /balance/{symbol} | Get balance |
*BalanceApi* | [**getWalletBalance**](Apis/BalanceApi.md#getwalletbalance) | **GET** /balance | Get wallet balance |
| *CoinsApi* | [**getCoinDetails**](Apis/CoinsApi.md#getcoindetails) | **GET** /coins/{symbol} | Get coin |
*CoinsApi* | [**getSupportedCoins**](Apis/CoinsApi.md#getsupportedcoins) | **GET** /coins | Get coins |
| *DepositApi* | [**depositMoney**](Apis/DepositApi.md#depositmoney) | **GET** /deposit | Deposit money |
| *HealthApi* | [**checkApiHealth**](Apis/HealthApi.md#checkapihealth) | **GET** /healthz | Health check |
| *LoansApi* | [**createLoanVault**](Apis/LoansApi.md#createloanvault) | **POST** /loans/create | Create vault |
*LoansApi* | [**estimateLoanCollateral**](Apis/LoansApi.md#estimateloancollateral) | **GET** /loans/estimate | Get collateral and loan amount estimation |
*LoansApi* | [**getAllLoanPlans**](Apis/LoansApi.md#getallloanplans) | **GET** /loans/plans | Get all plan loans |
*LoansApi* | [**getLoanDetails**](Apis/LoansApi.md#getloandetails) | **GET** /loans/{id} | Get loan |
*LoansApi* | [**getUserLoans**](Apis/LoansApi.md#getuserloans) | **GET** /loans | Get user loans |
*LoansApi* | [**repayLoan**](Apis/LoansApi.md#repayloan) | **POST** /loans/repay | Repay Loan |
| *PaymentApi* | [**createPayment**](Apis/PaymentApi.md#createpayment) | **POST** /payments | Create a payment |
| *PointsApi* | [**getFriendsPoints**](Apis/PointsApi.md#getfriendspoints) | **GET** /points/frineds | Get friends points. |
| *RedemptionsApi* | [**getAllRedemptions**](Apis/RedemptionsApi.md#getallredemptions) | **GET** /admin/redemptions | Get all redemptions |
*RedemptionsApi* | [**getRedemptionDetails**](Apis/RedemptionsApi.md#getredemptiondetails) | **GET** /admin/redemptions/{id} | Get redemption |
*RedemptionsApi* | [**redeemZar**](Apis/RedemptionsApi.md#redeemzar) | **POST** /redemptions | Redeem zar |
*RedemptionsApi* | [**updateRedemptionStatus**](Apis/RedemptionsApi.md#updateredemptionstatus) | **POST** /admin/redemptions/{id} | Update redemption |
| *ReferralsApi* | [**getReferralById**](Apis/ReferralsApi.md#getreferralbyid) | **GET** /referrals/{referralId} | Get referral by ID |
*ReferralsApi* | [**getReferrals**](Apis/ReferralsApi.md#getreferrals) | **GET** /referrals | Get referrals |
*ReferralsApi* | [**redeemReferral**](Apis/ReferralsApi.md#redeemreferral) | **POST** /referrals/{referralId}/redeem | Redeem a referral |
*ReferralsApi* | [**validateReferral**](Apis/ReferralsApi.md#validatereferral) | **POST** /referrals/{referralId}/validate | Validate a referral |
| *SwapApi* | [**swapCoins**](Apis/SwapApi.md#swapcoins) | **POST** /swap | Swap coins |
| *TasksApi* | [**getTasks**](Apis/TasksApi.md#gettasks) | **GET** /tasks | Get tasks |
| *TransactionsApi* | [**getUserTransactions**](Apis/TransactionsApi.md#getusertransactions) | **GET** /transactions | Get user transactions |
| *UserApi* | [**confirmKyc**](Apis/UserApi.md#confirmkyc) | **POST** /users/kyc/confirm | Confirm KYC |
*UserApi* | [**confirmPhoneNumber**](Apis/UserApi.md#confirmphonenumber) | **POST** /users/phone/confirm | Confirm phone number |
*UserApi* | [**createChildUser**](Apis/UserApi.md#createchilduser) | **POST** /users/children | create a child user |
*UserApi* | [**getUserProfile**](Apis/UserApi.md#getuserprofile) | **GET** /profile | Get profile |
*UserApi* | [**submitEmailConfirmationOtp**](Apis/UserApi.md#submitemailconfirmationotp) | **POST** /users/email/confirm | Submit email confirmation OTP |
*UserApi* | [**submitKyc**](Apis/UserApi.md#submitkyc) | **POST** /users/kyc | Submit KYC |
*UserApi* | [**verifyPhoneNumber**](Apis/UserApi.md#verifyphonenumber) | **POST** /users/phone | Verify phone number |
*UserApi* | [**verifyUserEmailAddress**](Apis/UserApi.md#verifyuseremailaddress) | **POST** /users/email | Verify email |
| *WithdrawApi* | [**getUserWithdrawRequests**](Apis/WithdrawApi.md#getuserwithdrawrequests) | **GET** /withdraws | Get user withdraw requests |
*WithdrawApi* | [**getWithdrawalStatus**](Apis/WithdrawApi.md#getwithdrawalstatus) | **GET** /withdraws/{id} | Get withdrawal status |
*WithdrawApi* | [**previewWithdrawal**](Apis/WithdrawApi.md#previewwithdrawal) | **POST** /withdraws/preview | Withdraw request |
*WithdrawApi* | [**requestWithdrawal**](Apis/WithdrawApi.md#requestwithdrawal) | **POST** /withdraws/request | Withdraw request |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [AdminRedemptionUpdateRequest](./Models/AdminRedemptionUpdateRequest.md)
 - [AuthTelegramRequest](./Models/AuthTelegramRequest.md)
 - [Balance](./Models/Balance.md)
 - [BankInfo](./Models/BankInfo.md)
 - [BulletContent](./Models/BulletContent.md)
 - [Coin](./Models/Coin.md)
 - [CoinConfig](./Models/CoinConfig.md)
 - [CoinResponse](./Models/CoinResponse.md)
 - [CreateChildUserRequest](./Models/CreateChildUserRequest.md)
 - [Currency](./Models/Currency.md)
 - [DepositResponse](./Models/DepositResponse.md)
 - [DetailedLoanToValueOptions](./Models/DetailedLoanToValueOptions.md)
 - [EmailOtpSubmitRequest](./Models/EmailOtpSubmitRequest.md)
 - [Error](./Models/Error.md)
 - [ErrorDetail](./Models/ErrorDetail.md)
 - [ExternalTransaction](./Models/ExternalTransaction.md)
 - [FriendPoints](./Models/FriendPoints.md)
 - [FriendPointsResponse](./Models/FriendPointsResponse.md)
 - [HealthStatus](./Models/HealthStatus.md)
 - [JwtResponse](./Models/JwtResponse.md)
 - [KycConfirmRequest](./Models/KycConfirmRequest.md)
 - [KycRequest](./Models/KycRequest.md)
 - [KycResponse](./Models/KycResponse.md)
 - [LoanCreateRequest](./Models/LoanCreateRequest.md)
 - [LoanPlan](./Models/LoanPlan.md)
 - [LoanPlanResponse](./Models/LoanPlanResponse.md)
 - [LoanToValueOptions](./Models/LoanToValueOptions.md)
 - [LoansResponse](./Models/LoansResponse.md)
 - [LoansResponseList](./Models/LoansResponseList.md)
 - [LoginRequest](./Models/LoginRequest.md)
 - [Network](./Models/Network.md)
 - [Payment](./Models/Payment.md)
 - [PaymentRequest](./Models/PaymentRequest.md)
 - [PhoneOtpSubmitRequest](./Models/PhoneOtpSubmitRequest.md)
 - [ProfileResponse](./Models/ProfileResponse.md)
 - [Redemption](./Models/Redemption.md)
 - [RedemptionRequest](./Models/RedemptionRequest.md)
 - [RedemptionResponse](./Models/RedemptionResponse.md)
 - [Referral](./Models/Referral.md)
 - [ReferralResponse](./Models/ReferralResponse.md)
 - [RepayLoanRequest](./Models/RepayLoanRequest.md)
 - [SignUpRequest](./Models/SignUpRequest.md)
 - [SimpleResponse](./Models/SimpleResponse.md)
 - [SwapRequest](./Models/SwapRequest.md)
 - [SwapResponse](./Models/SwapResponse.md)
 - [Symbol](./Models/Symbol.md)
 - [Task](./Models/Task.md)
 - [TaskResponse](./Models/TaskResponse.md)
 - [TelegramProfile](./Models/TelegramProfile.md)
 - [Timestamp](./Models/Timestamp.md)
 - [Transaction](./Models/Transaction.md)
 - [TransactionResponse](./Models/TransactionResponse.md)
 - [TransactionStatus](./Models/TransactionStatus.md)
 - [TransactionType](./Models/TransactionType.md)
 - [UpdateEmailRequest](./Models/UpdateEmailRequest.md)
 - [UpdatePhoneRequest](./Models/UpdatePhoneRequest.md)
 - [User](./Models/User.md)
 - [UserError](./Models/UserError.md)
 - [WalletBalance](./Models/WalletBalance.md)
 - [WithdrawRequest](./Models/WithdrawRequest.md)
 - [WithdrawRequestBody](./Models/WithdrawRequestBody.md)
 - [WithdrawRequestPreview](./Models/WithdrawRequestPreview.md)
 - [WithdrawRequestResponse](./Models/WithdrawRequestResponse.md)
 - [WithdrawResponseBody](./Models/WithdrawResponseBody.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="bearerAuth"></a>
### bearerAuth

- **Type**: HTTP Bearer Token authentication (JWT)

