package main

import "github.com/taosdata/driver-go/v3/bench/standard/executor"

func main() {
	test := executor.NewTDTest(executor.DefaultNativeDriverName, executor.DefaultNativeDSN)

	test.Clean()
	tableName := test.PrepareRead(1000, 100)
	test.BenchmarkRead("select * from " + tableName)
}
