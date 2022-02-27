
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./person/hello_grpc.proto
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./person/person.proto
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./home/home.proto
protoc -I ./ --go_out=. --go_opt=paths=source_relative  --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./personreq/person.proto