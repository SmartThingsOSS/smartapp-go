# smartapp-go

A simple framework for creating SmartThings SmartApp's in Golang.

__This is an experimental library, and should be treated as such.__

## Dependencies
* Install Go (https://github.com/moovweb/gvm)
* Install Glide (https://github.com/Masterminds/glide)
* `make get-deps` will install dependencies via Glide.
* symlink the `smartapp-go.yml` into `~/.smartapp-go/`. i.e.
```
mkdir ~/.smartapp-go
ln -s $GOPATH/src/github.com/SmartThingsOSS/smartapp-go/smartapp-go.yml ~/.smartapp-go
```

## Running Examples
After fetching dependencies examples can be run by issueing the following commands from project root directory:
* `go run examples/basic/main.go`
* `go run examples/c2c-connector/main.go`
