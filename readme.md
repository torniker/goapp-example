# Setup

`go get github.com/torniker/wrap`

`go run main.go`




`openssl genrsa -out server.key 2048`
`openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
