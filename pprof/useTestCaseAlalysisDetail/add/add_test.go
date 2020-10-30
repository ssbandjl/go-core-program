package add

//执行测试和采集命令:go test -bench=. -cpuprofile=cpu.profile
//利用pprof工具分析:go tool pprof -http=:6001 cpu.profile

//内存采集:go test -bench=. -memprofile=mem.profile
//分析:go tool pprof -http=:6001 mem.profile

import "testing"

func TestAdd(t *testing.T) {
	_ = Add("go-programming-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("go-programming-tour-book")
	}
}
