https://github.com/szamcsi/golang-elte-2019-public/

```shell
 $ git clone https://github.com/szamcsi/golang-elte-2019-public.git
 $ cd golang-elte-2019-public/testing/ex1
 $ go run lc.go lc.go
```

 - http://golang.org/pkg/testing#T
 - https://golang.org/pkg/path/filepath/#Join
 - https://golang.org/wiki/TableDrivenTests

https://blog.golang.org/cover

```shell
$ go test -coverprofile=coverage.out lc.go lc_test.go 
$ go tool cover -func=coverage.out
```

https://godoc.org/github.com/google/go-cmp/cmp

```shell
$ go get github.com/google/go-cmp/cmp
```

