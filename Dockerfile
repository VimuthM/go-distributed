# From jfbrandhorst/grpc-web-generators, with later versions.
FROM jfbrandhorst/grpc-web-generators
# FROM ubuntu:latest

# ENV DEBIAN_FRONTEND noninteractive

# RUN apt-get update && apt-get install -y \
#     automake \
#     build-essential \
#     git \
#     libtool \
#     make \
#     npm \
#     wget \
#     unzip \
#     libprotoc-dev \
#     python3-pip \
#     golang

# ## Install protoc

# # ENV PROTOBUF_VERSION 3.12.1
# # Later versions do not have proto-gen-js
# # https://stackoverflow.com/questions/72572040/protoc-gen-js-program-not-found-or-is-not-executable
# ENV PROTOBUF_VERSION 3.20.1

# RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-linux-x86_64.zip && \
#     unzip protoc-$PROTOBUF_VERSION-linux-x86_64.zip -d /usr/local/ && \
#     rm -rf protoc-$PROTOBUF_VERSION-linux-x86_64.zip

# ## Install protoc-gen-go

# ENV PROTOC_GEN_GO_VERSION v1.5.3

# RUN git clone https://github.com/golang/protobuf /root/go/src/github.com/golang/protobuf && \
#     cd /root/go/src/github.com/golang/protobuf && \
#     git fetch --all --tags --prune && \
#     git checkout tags/$PROTOC_GEN_GO_VERSION && \
#     go install ./protoc-gen-go && \
#     ln -s /root/go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go && \
#     rm -rf /root/go/src

# ## Install protoc-gen-grpc-web

# ENV PROTOC_GEN_GRPC_WEB_VERSION 1.4.2

# RUN git clone https://github.com/grpc/grpc-web /github/grpc-web && \
#     cd /github/grpc-web && \
#     git fetch --all --tags --prune && \
#     git checkout tags/$PROTOC_GEN_GRPC_WEB_VERSION && \
#     make install-plugin && \
#     rm -rf /github

# ## Install protoc-gen-ts

# ENV PROTOC_GEN_TS_VERSION 0.12.0

# RUN npm install ts-protoc-gen@$PROTOC_GEN_TS_VERSION google-protobuf@$PROTOBUF_VERSION && \
#     ln -s /node_modules/.bin/protoc-gen-ts /usr/local/bin/protoc-gen-ts


WORKDIR /app

COPY backend/api /app/api
COPY frontend/src/jsclient /app/jsclient

CMD ["protoc", "-I", "/app/api/", \
    "--go_out=plugins=grpc,paths=source_relative:/app/api/new", \
    # "--go-grpc_out=plugins=grpc,paths=source_relative:/app/api/new", \
    "--js_out=import_style=commonjs,binary:/app/jsclient", \
    # "--ts_out=import_style=commonjs,service=grpc-web:/app/jsclient", \
    "--grpc-web_out=import_style=commonjs,mode=grpcwebtext:/app/jsclient", \
    "/app/api/runner/runner.proto"]
