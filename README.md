# gin_benchmark
* Go 1.5.2
* Gin 1.0rc2

```
$ go run test.go

$ wrk -t4 -c100 -d30S --timeout 2000 "http://127.0.0.1:8080/welcome"
```
