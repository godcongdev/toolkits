package toolkits

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./errors/rpcerr/*.proto
