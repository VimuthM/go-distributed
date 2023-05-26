proto:
	@echo "--> Generating gRPC clients"
	@docker run \
		-v `pwd`/backend/api/runner:/api \
		-v `pwd`/backend/cmd/client:/goclient \
		-v `pwd`/frontend/src/jsclient:/jsclient \
		jfbrandhorst/grpc-web-generators \
		protoc -I /api \
			--go_out=plugins=grpc,paths=source_relative:/api \
			/api/runner.proto
	@echo "--> Generating gRPC-web clients"
	@cd frontend && \
		npm install && \
		npx protoc \
			--ts_out src/jsclient \
			--ts_opt long_type_string \
			--proto_path ../backend/api/runner \
			../backend/api/runner/runner.proto

# --plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
#-grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \

run-servers:
	@echo "--> Starting servers"
	@docker-compose up

run-frontend:
	@echo "--> Starting frontend"
	cd ./frontend
	npm run serve
