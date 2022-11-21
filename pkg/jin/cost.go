package jin

import (
	"log"
	"time"
)

//执行耗时
func Cost() HandleFunc {
	// 使用函数回调
	return func(c *Context) {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.request.URL, cost.Seconds())
	}
}
