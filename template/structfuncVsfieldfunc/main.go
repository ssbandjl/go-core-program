package main

import (
	"log"
	"os"
	"text/template"
)

type T struct {
	Add func(int) int
}

func (t *T) Sub(i int) int {
	log.Println("执行方法: get argument i:", i)
	return i - 1
}

func main() {
	ts := &T{
		Add: func(i int) int {
			return i + 1
		},
	}
	tpl := `
// 只能使用 call 调用
call field func Add 返回值: {{ call .ts.Add .y }}
// 直接传入 .y 调用
call method func Sub: {{ .ts.Sub .y }}
	`
	t, _ := template.New("test").Parse(tpl)
	t.Execute(os.Stdout, map[string]interface{}{
		"y":  3,
		"ts": ts,
	})
}
