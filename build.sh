go build -o dist/main.dll -buildmode=c-shared cmd/client/main.go
go build -o dist/main.so -buildmode=c-shared cmd/client/main.go
md5 dist/main.dll dist/main.so