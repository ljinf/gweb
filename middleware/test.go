package middleware

import (
	"fmt"
	"github.com/ljinf/gweb"
)

func Test1() gweb.ControllerHandler {
	// 使用函数回调
	return func(c *gweb.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() gweb.ControllerHandler {
	// 使用函数回调
	return func(c *gweb.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() gweb.ControllerHandler {
	// 使用函数回调
	return func(c *gweb.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
