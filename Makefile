.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:. buff/*.proto

run:
	go run main.go server

echosvc-example:
	go run example/client/echosvcclient/main.go