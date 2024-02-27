#!/bin/bash

# Convert Swagger to OpenAPI v3
swagger2openapi --yaml --outfile backend-api.yaml ./pkg/generated/social_service/v2/apidocs.swagger.json

# Generate TypeScript client using Axios
openapi-generator generate \
    -i ./backend-api.yaml \
    -g typescript-axios -o ./sdk/typescript-sdk/src/typescript-axios

# Generate documentation in YAML format
openapi-generator generate \
    -i ./backend-api.yaml \
    -g openapi-yaml -o ./documentation/openapi-yaml

# Generate documentation in Markdown format
openapi-generator generate \
    -i ./backend-api.yaml \
    -g markdown -o ./documentation/markdown

# Generate Postman collection
openapi-generator generate \
    -i ./backend-api.yaml \
    -g postman-collection -o ./documentation/postman-collection

swagger2openapi --outfile backend-api.json ./pkg/generated/social_service/v2/apidocs.swagger.json

# from the generate openapiv3 json file, we generate the krakend.json file
# TODO: actual generation
go run ./cmd/openapi2krakend/main.go -directory=path/to/swagger -output=my_output.json

# we copy the generated file into the backend api gateway folder which will be auto deployed upon changes being merged into prod
