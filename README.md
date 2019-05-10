Public course material for the ELTE Go language course in 2019.

### 2019-02-14
Introduction, environmental setup: [slides](env-slides.pdf)

Install go on OS linux 64-bit use the install_go.sh script in the script/ folder

```shell
$ chmod +x install_go.sh
$ ./script/install_go.sh
```

it will install go version 1.11.4 into your /tmp folder and set all the necessary environment variables

### 2019-02-21, 2019-02-28, 2019-03-07
Introduction to the language, with [exercises](intro/README.md):

  - https://tour.golang.org/
  - https://golang.org/doc/
  - https://golang.org/pkg/
  - https://golang.org/doc/effective_go.html
  - https://code.google.com/p/go-wiki/wiki/SliceTricks

### 2019-03-14
HTTP sever programming with exercises in the `server` directory.

  - https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
  - https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355

Robust servers with Context:

  - https://blog.golang.org/context
  - https://peter.bourgon.org/go-for-industrial-programming/

### 2019-03-21
Unit testing with [slides](testing-slides.pdf) and exercises in the `testing` directory.

  - https://code.google.com/p/go-wiki/wiki/TableDrivenTests

### 2019-03-28
gRPC testing with [slides](grpc-testing-slides.pdf) and exercises in the `grpc-testing` directory.

```shell
$ go get -u google.golang.org/grpc
$ go get -u google.golang.org/genproto
$ go get -u github.com/golang/protobuf/protoc-gen-go

$ mkdir -p protoc/3.7.0-rc3
$ cd protoc/3.7.0-rc3
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0-rc.3/protoc-3.7.0-rc-3-linux-x86_64.zip
$ unzip protoc-3.7.0-rc-3-linux-x86_64.zip
$ export PROTOC_ROOT=$PWD
$ export PATH=$PATH:$PWD/bin
```

### Recommended Book

[The Go Programming Language](https://www.amazon.de/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)

