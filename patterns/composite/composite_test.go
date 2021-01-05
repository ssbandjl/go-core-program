package composite

//参考文档: https://golang.org/pkg/testing/

import (
	"go_code/patterns/composite/concurrency"
	"go_code/patterns/composite/normal"
	"log"
	"runtime"
	"testing"
)

// func TestComposite(t *testing.T) {
// 	normal.Demo()
// }

// func TestCompositeConcurrency(t *testing.T) {
// 	concurrency.Demo()
// }

// go test -benchmem -run=^$ go_code/go/patterns/composite -bench . -v -count=1 --benchtime 20s
// go test -benchmem -run=^$ . -bench . -v -count=1 --benchtime 20s
// -benchmem, 打印内存分配
// -run 只运行正则匹配的测试用例
// -bench 只运行正则匹配的基准测试
// -v 显示详情
// -count=1, 每个测试运行1次
// --benchtime 对每个基准测试, 迭代运行20秒

// 基准测试方法名匹配规则:func BenchmarkXxx(*testing.B)
//If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function;
//such benchmarks are intended to be used with the go test -cpu flag
// 如果需要在并行设置中做性能测试,可以使用RunParallel方法, 这类方法会与go test -cpu参数一起使用
func Benchmark_Normal(b *testing.B) {
	log.Printf("CPU Num:%d", runtime.NumCPU())
	b.SetParallelism(runtime.NumCPU()) //设置并发使用的CPU个数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			normal.Demo()
		}
	})
}

func Benchmark_Concurrency(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concurrency.Demo()
		}
	})
}
