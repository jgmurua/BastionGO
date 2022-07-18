# BastionGO
# Testea tus endpoint privados desde una funcion lambda en go

## build
```sh
go mod init BastionGO
go mod tidy
##go build main.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip main.zip main
```

## test data
```sh
{
  "Url": "https://httpbin.org/post",
  "Method": "POST",
  "Payload": "`{ 'hola': 'hola'}`"
}
```
