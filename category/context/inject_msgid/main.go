/*
withValue() 的完整工作示例。 在下面的示例中，我们为每个传入请求注入一个 msgId。 如果您在下面的程序中注意到

inejctMsgID 是一个网络 HTTP 中间件函数，它在上下文中填充“msgID”字段
HelloWorld 是 api "localhost:8080/welcome" 的处理函数，它从上下文中获取这个 msgID 并将其作为响应标头发送回来
在


测试:

curl -v http://localhost:8080/welcome

*/

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", inejctMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

//HelloWorld hellow world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID) // head中注入Msgid: 8b03ecdd-922c-4960-8431-aab30dbda1bf
	w.Write([]byte("Hello, world"))
}

func inejctMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		fmt.Println("msgID:", msgID)
		ctx := context.WithValue(r.Context(), "msgId", msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)

	})
}
