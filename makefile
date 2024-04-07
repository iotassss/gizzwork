.PHONY:
	generate,
	generate-petstore,
	generate2,
	generate-employee,

# openapi.ymlからGo言語のコードを生成
generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/api/employee.yml \
		-g go-gin-server \
		-o /local \
		--additional-properties packageName=api

generate-petstore:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
		-g go-server \
		-o /local  \
		--additional-properties packageName=petstore \
		--type-mappings number=float64

generate2:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/api/employee.yml \
		-g go-gin-server \
		-o /local \
		--additional-properties packageName=api
		--additional-properties withGoMod=false \
		--openapi-generator-ignore-list api/*,test/*,.gitignore,.travis.yml,git_push.sh

generate-employee-server:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/api/employee.yml \
		-g go-server \
		-o /local/employee \
		-t /local/employee/templates \
		--additional-properties gitUserId=iotassss \
		--additional-properties packageName=api

# generate-employee-client-package:
generate-employee-server-templates:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli author template \
		-g go-server \
		-o /local/employee/templates
