SWAGGER_SOCIAL_SERVICE_CODEGEN_PATH:=./core-library/pkg/generated/social_service/v2/apidocs.swagger.json
SWAGGER_USER_SERVICE_CODEGEN_PATH:=./core-library/pkg/generated/user_service/v1/apidocs.swagger.json
SWAGGER_FINANCIAL_SERVICE_CODEGEN_PATH:=./core-library/pkg/generated/financial_service/v1/apidocs.swagger.json
SWAGGER_ACCOUNTING_SERVICE_CODEGEN_PATH:=./core-library/pkg/generated/accounting_service/v1/apidocs.swagger.json
SWAGGER_WORKSPACE_SERVICE_CODEGEN_PATH:=./core-library/pkg/generated/workspace_service/v1/apidocs.swagger.json

VERSION = 3
FILE_PROD = ./krakend-config/final-krakend.prod.json
FILE_STAGING = ./krakend-config/final-krakend.staging.json

gen: 
	cd core-library && make gen && cd ../

copy-swagger:
	cp -rf $(SWAGGER_SOCIAL_SERVICE_CODEGEN_PATH) ./swagger/social-service.json
	cp -rf $(SWAGGER_USER_SERVICE_CODEGEN_PATH) ./swagger/user-service.json
	cp -rf $(SWAGGER_FINANCIAL_SERVICE_CODEGEN_PATH) ./swagger/financial-service.json
	cp -rf $(SWAGGER_ACCOUNTING_SERVICE_CODEGEN_PATH) ./swagger/accounting-service.json
	cp -rf $(SWAGGER_WORKSPACE_SERVICE_CODEGEN_PATH) ./swagger/workspace-service.json 


convert-swagger-to-openapiv3:
	swagger2openapi --yaml --outfile ./swagger/social-service-backend-api.yaml ./swagger/social-service.json
	swagger2openapi --yaml --outfile ./swagger/user-service-backend-api.yaml ./swagger/user-service.json
	swagger2openapi --yaml --outfile ./swagger/financial-service-backend-api.yaml ./swagger/financial-service.json
	swagger2openapi --yaml --outfile ./swagger/accounting-service-backend-api.yaml ./swagger/accounting-service.json
	swagger2openapi --yaml --outfile ./swagger/workspace-service-backend-api.yaml ./swagger/workspace-service.json

	swagger2openapi --outfile ./swagger/social-service-backend-api.json ./swagger/social-service.json
	swagger2openapi --outfile ./swagger/user-service-backend-api.json ./swagger/user-service.json
	swagger2openapi --outfile ./swagger/financial-service-backend-api.json ./swagger/financial-service.json
	swagger2openapi --outfile ./swagger/accounting-service-backend-api.json ./swagger/accounting-service.json
	swagger2openapi --outfile ./swagger/workspace-service-backend-api.json ./swagger/workspace-service.json
	
	swagger2openapi --yaml --outfile ./swagger/workspace-service-rest-api.yaml ./api-packages/workspace-service-http/docs/swagger.json
	swagger2openapi --outfile ./swagger/workspace-service-rest-api.json ./api-packages/workspace-service-http/docs/swagger.json

	npx openapi-merge-cli --config ./swagger/merge-config.json
	swagger2openapi --outfile ./swagger/backend-api.json ./swagger/backend-api.yaml

# generate the various typescript sdks for this core-lib
update-typescript-sdk:
	openapi-generator generate \
    	-i ./swagger/backend-api.yaml \
    	-g typescript-fetch -o ./solomon-ai-typescript-sdk/src/typescript-fetch

update-docs:
	openapi-generator generate \
    -i ./swagger/backend-api.yaml \
    -g openapi-yaml -o ./documentation/autogen/openapi-yaml

	openapi-generator generate \
		-i ./swagger/backend-api.yaml \
		-g markdown -o ./documentation/autogen/markdown

	openapi-generator generate \
		-i ./swagger/backend-api.yaml \
		-g postman-collection -o ./documentation/autogen/postman-collection

generate-krakend-config:
	mkdir temp-swagger && cp ./swagger/user-service-backend-api.json ./temp-swagger/user-service-backend-api.json && cp ./swagger/social-service-backend-api.json ./temp-swagger/social-service-backend-api.json && cp ./swagger/financial-service-backend-api.json ./temp-swagger/financial-service-backend-api.json && cp ./swagger/accounting-service-backend-api.json ./temp-swagger/accounting-service-backend-api.json && cp ./swagger/workspace-service-backend-api.json ./temp-swagger/workspace-service-backend-api.json && cp ./swagger/workspace-service-rest-api.json ./temp-swagger/workspace-service-rest-api.json
	go run ./openapi2krakend/pkg/main.go -directory ./temp-swagger -output ./krakend-config/krakend.prod.json -environment production -webhook-config-path ./webhooks.yaml
	go run ./openapi2krakend/pkg/main.go -directory ./temp-swagger -output ./krakend-config/krakend.staging.json -environment staging -webhook-config-path ./webhooks.yaml
	rm -rf temp-swagger

prettify-krakend:
	cat ./krakend-config/krakend.json | jq >> ./krakend-config/krakend.compiled.json
	mv ./krakend-config/krakend.compiled.json ./krakend-config/krakend.json

merge-configs: 
	jq -s '.[0] * .[1]' ./krakend-config/gateway-configurations.json ./krakend-config/krakend.prod.json > ./krakend-config/final-krakend.prod.json
	@sed -i '' 's/"version": "3"/"version": $(VERSION)/' $(FILE_PROD)
	jq -s '.[0] * .[1]' ./krakend-config/gateway-configurations.json ./krakend-config/krakend.staging.json > ./krakend-config/final-krakend.staging.json
	@sed -i '' 's/"version": "3"/"version": $(VERSION)/' $(FILE_STAGING)

validate:
	krakend check --config ./krakend-config/final-krakend.prod.json -ddd
	krakend check -tlc ./krakend-config/final-krakend.prod.json
	krakend check --config ./krakend-config/final-krakend.staging.json -ddd
	krakend check -tlc ./krakend-config/final-krakend.staging.json

copy-configs-to-gateway:
	cp ./krakend-config/final-krakend.prod.json ./backend-api-gateway/krakend.prod.json
	cp ./krakend-config/final-krakend.staging.json ./backend-api-gateway/krakend.staging.json

gen-workspace-service-api:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag/gen@latest
	go get github.com/swaggo/swag/cmd/swag@latest
	cd api-packages/workspace-service-http && $$(go env GOPATH)/bin/swag init -g api.go

autogen: gen-workspace-service-api gen copy-swagger convert-swagger-to-openapiv3 update-typescript-sdk update-docs generate-krakend-config prettify-krakend merge-configs validate copy-configs-to-gateway

generate: gen-workspace-service-api gen copy-swagger convert-swagger-to-openapiv3 update-typescript-sdk update-docs generate-krakend-config prettify-krakend merge-configs validate copy-configs-to-gateway
