install_benchstat:
	go install golang.org/x/perf/cmd/benchstat@latest

bench_r1000_w1:
	go test -bench=BenchmarkRead1000Write1 -count=10 -run=^ > bench_results_r1000_w1.txt
	benchstat bench_results_r1000_w1.txt

bench_r1_w1:
	go test -bench=BenchmarkReadWriteOne -count=10 -run=^ > bench_results_r1_w1.txt
	benchstat bench_results_r1_w1.txt

bench_r1_w1000:
	go test -bench=BenchmarkRead1Write1000 -count=10 -run=^ > bench_results_r1_w1000.txt
	benchstat bench_results_r1_w1000.txt

bench_all:
	go test -bench=. -count=10 > bench_all.txt
	benchstat bench_all.txt