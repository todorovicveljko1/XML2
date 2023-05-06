gen: 	
	@protoc --go_out=./services/$(service) --go-grpc_out=./services/$(service) ./proto/${service}.proto
	@protoc --go_out=./api_gateway  --go-grpc_out=./api_gateway ./proto/${service}.proto

	@echo "generate ${service} done"