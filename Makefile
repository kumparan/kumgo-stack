.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:. buff/*.proto

run:
	go run main.go server

story-client-example:
	go run example/client/storysvcclient/main.go

echo-example-proto:
	protoc --go_out=plugins=grpc:. example/server/echosvcserver/buff/*.proto

echo-server-example:
	go run example/server/echosvcserver/main.go

echo-client-example:
	go run example/client/echosvcclient/main.go