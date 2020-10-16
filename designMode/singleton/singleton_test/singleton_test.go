//基准测试,对比懒汉式双重检查与不带双重检查的性能,参考命令:go test -bench=. -benchmem
package singleton_test

import (
	"go_code/designMode/singleton/lazy"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lazy.New()
	}
}

func BenchmarkNew2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lazy.New2()
	}
}
