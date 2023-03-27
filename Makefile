proto:
	protoc --go_out=. --go_opt=paths=source_relative v1/employee.proto
tag:
	protoc --go_out=. --go_opt=paths=source_relative --go_opt=tags="bson,json" v1/employee.proto


clean:
	rm -f v1/employee.pb.go
