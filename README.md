Lint install:

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.6

go get -u golang.org/x/tools/cmd/cover

go get github.com/Shodocan/InstanceInventoryApi

make build -> realiza as checagens de cobertura, gera o relatorio de coverage e gera um compilado