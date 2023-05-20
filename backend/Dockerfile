# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
ADD . /src
WORKDIR /src
RUN go get ./...&& go build -o distributed

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/distributed /app/
ENTRYPOINT ./distributed