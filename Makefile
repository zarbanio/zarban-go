

codegen:
	oapi-codegen -package wallet -generate "types,client" api_specs/wallet.openapi.yaml > wallet/client.go
	oapi-codegen -package service -generate "types,client" api_specs/service.openapi.yaml > service/client.go
