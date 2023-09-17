gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ./proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go
