# packio-golang

Step by step
```
> go mod init github.com/nuttipong/packio-golang
> go build .
> go run --race . adhoc_share_data_20.csv
```

Get dependencies
```
> go get -u github.com/gorilla/mux
> go list -m all
> go list -m -versions github.com/gorilla/mux
> go mod verify
```

Clean up no longer use dependencies
```
> go mod tidy
```

Setup Pact
```
> go get gopkg.in/pact-foundation/pact-go.v1
> curl -LO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v1.88.81/pact-1.88.81-osx.tar.gz
tar xzf pact-1.88.81-osx.tar.gz
cd pact/bin
./pact-mock-service --help start
./pact-provider-verifier --help verify
```

Add Path
```
export PACT_PATH=/Users/nuttipongtaechasanguanwong/pact/bin
export PATH=$PACT_PATH:$PATH
```

Test Consumer
```
> go test -v -run TestConsumer .
```

Test Provider
```
> go test -v -run TestConsumer .
```