protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative api/runner/runner.proto

docker run -v `pwd`/backend/api:/api -v `pwd`/backend/cmd/goclient:/goclient -v `pwd`/frontend/src/jsclient:/jsclient jfbrandhorst/grpc-web-generators \
    protoc \
	  --go_out=. --go_opt=paths=source_relative:/goclient \
	  --js_out=import_style=commonjs:/jsclient \
	  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
	api/runner/runner.proto


docker run -v `pwd`/api:/api -v `pwd`/backend/goclient:/goclient -v `pwd`/frontend/src/jsclient:/jsclient jfbrandhorst/grpc-web-generators protoc -I /api \
	  --go_out=plugins=grpc,paths=source_relative:/goclient \
	  --js_out=import_style=commonjs:/jsclient \
	 /api/runner/runner.proto

docker run -v "/c/Users/vimut/Desktop/go-distributed/api:/api" -v "/c/Users/vimut/Desktop/go-distributed/backend/goclient:/goclient" -v "/c/Users/vimut/Desktop/go-distributed/frontend/src/jsclient:/jsclient" jfbrandhorst/grpc-web-generators protoc -I /api \
  --go_out=plugins=grpc,paths=source_relative:/goclient \
  --js_out=import_style=commonjs:/jsclient \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
  "/api/runner/runner.proto"


protoc   --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. api/runner/runner.proto