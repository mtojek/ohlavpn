build: clean-code install

clean-code:
	go get golang.org/x/tools/cmd/goimports && goimports -w .
	gofmt -s -w .
	go get golang.org/x/lint/golint && golint -set_exit_status ./...

install:
	go get -v github.com/mtojek/ohlavpn

