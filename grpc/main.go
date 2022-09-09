package main

import "fmt"

// multiple options of run proto buf

// 1 run relative path protoc --go_out=./src --go_opt=paths=source_relative  proto/employee.proto
// this creates a directory in src but with tree src/proto

// 2 run I protoc -I ./ --go_out=./src proto/employee.proto
// this creates a directory in src but don't create dir proto

// 3 run file js protoc --js_out=import_style=commonjs,binary:. proto/employee.proto

func main() {
	fmt.Println("hello word")
}
