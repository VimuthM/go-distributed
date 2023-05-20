proto:
	@echo "--> Generating gRPC clients"
	@docker run -v `pwd`/backend/api:/backend/api -v `pwd`/backend/cmd/goclient:/goclient -v `pwd`/frontend/src/jsclient:/jsclient jfbrandhorst/grpc-web-generators protoc -I /backend/api \
	  --go_out=plugins=grpc,paths=source_relative:/goclient \
	  --js_out=import_style=commonjs:/jsclient \
	  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
	 /backend/api/runner/runner.proto

run-servers:
	@echo "--> Starting servers"
	@docker-compose up

run-frontend:
	@echo "--> Starting frontend"
	cd ./frontend
	npm run serve
