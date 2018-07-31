# Go REST API Example

## Setup

### Go Install

```bash
$ brew install golang
$ mkdir ~/go
$ export GOPATH="$HOME/go"
$ export GOBIN="$GOPATH/bin"
$ export PATH="$GOBIN:$PATH"
```

### Project Setup

```bash
$ mkdir -p ~/go/src/github.com/spilth/
$ cd ~/go/src/github.com/spilth/
$ git clone git@github.com:spilth/golang-rest-api.git
$ cd golang-rest-api
$ go build
$ ./golang-rest-api
```

Then visit <http://localhost:8000/api/v1/message>
