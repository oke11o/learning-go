# Maps

- [syncmap_bench.go](00-syncmap-bench/syncmap_bench.go)


## benchstat

```shell
go install golang.org/x/perf/cmd/benchstat@latest
benchstat --help
```

You can have problem: `command not found: benchstat`. Fix it:

```shell
go env GOPATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zprofile
```