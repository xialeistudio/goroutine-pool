# golang通用协程池
简化协程管理

## 实现功能

+ 限制协程数量
+ 阻塞等待指定数量任务执行（sync.WaitGroup）

## 使用方法

```go
pool := New(2, 1) // 初始化一个大小为2，等待长度为1的协程池
pool.Submit(func() {    
	time.Sleep(time.Second) // 提交任务
})
// 等待所有协程执行完毕
pool.Wait()
```

## 个人博客

[每天进步一点点](https://www.ddhigh.com)