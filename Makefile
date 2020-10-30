generate:
	#protoc -I=./darkbase --go_out=./protogen ./darkbase/darkbase.darkbase
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./darkbase/darkbase.proto
