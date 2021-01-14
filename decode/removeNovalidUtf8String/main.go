package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/golang/glog"
)

// removeNotValidUtf8InString 移除无效的UTF8编码
// golang中的rune类型 rune是int32类型，占4个字节，golang的默认编码方式为utf-8，每个utf-8字符（占1~4个字节）都可转化为rune类型
func removeNotValidUtf8InString(s string) string {
	ret := s
	if !utf8.ValidString(s) {
		glog.V(4).Infof("删除不规范utf-8编码前：%q", ret)
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				// r是尝试读取的一个utf-8字符，并不是一个字节
				// if the encoding is invalid, it returns (RuneError, 1)
				// DecodeRuneInString 取出字符串中第一个合法UTF-8字符rune类型
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		ret = string(v)
		glog.V(4).Infof("删除不规范utf-8编码后: %q", ret)
	}
	return ret
}

func main() {

	s := "截取中文"
	//试试这样能不能截取?
	res := []rune(s)                  //rune是int32的别名, 这里将中文字符串转为int32, byte 表示一个字节，rune 表示四个字节, 中文字符串每个占三个字节, 利用 [] rune 转换成 unicode 码点， 再利用 string 转化回去
	fmt.Println("s转为int32字节码为:", res) //[25130 21462 20013 25991]
	log.Printf(string(res[:2]))

	old := "sdajglkwjkelgjkjglkajwljegl"
	new := removeNotValidUtf8InString(old)
	log.Printf("移除非UTF8编码前:%s", old)
	log.Printf("移除非UTF8编码后:%s", new)
}
