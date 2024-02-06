proto-gen:
	protoc calculator/calculatorpb/calculator.proto --go-grpc_out=.
	protoc calculator/calculatorpb/calculator.proto --go_out=.
run-client:
	go run client/client.go
run-server:
	go run server/server.go