package grifts

import (
	"errors"
	"fmt"
	"os"
	"strings"

	. "github.com/markbates/grift/grift"
)

// 加个骗子。 如果已经有一个具有给定名称的 grift，则两个 grift 将捆绑在一起。
var _ = Add("boom", func(c *Context) error {
	return errors.New("boom!!!")
})

var _ = Add("hello", func(c *Context) error {
	fmt.Println("Hello World!")
	return nil
})

var _ = Add("hello", func(c *Context) error {
	fmt.Println("Hello World! Again")
	err := Run("db:migrate", c)
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("### dir -> %+v\n", dir)
	return nil
})

var _ = Add("env:print", func(c *Context) error {
	if len(c.Args) >= 1 {
		for _, e := range c.Args {
			fmt.Printf("%s=%s\n", e, os.Getenv(e))
		}
	} else {
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			fmt.Printf("%s=%s\n", pair[0], os.Getenv(pair[0]))
		}
	}

	return nil
})

// 命名空间会将所有任务放在给定的前缀内。如这里的db
var _ = Namespace("db", func() {
    Desc("migrate", "Migrates the databases")
		// 设置一个骗子。 这类似于“添加”，但它会覆盖现有的同名 grift。
    Set("migrate", func(c *Context) error {
            fmt.Println("db:migrate")
            fmt.Printf("### args -> %+v\n", c.Args)
            return nil
    })
}