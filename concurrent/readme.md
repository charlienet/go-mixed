延迟队列

包含以下实现模型

1. 内存模式。
2. Redis模式
3. MQ模式


内存模式，在队列中放入。定时检查过期时间，过期时间到达时取出并使用协程执行。

Redis模式，使用ZSET存储任务队列。定时取出规则内的任务。取出后使用ZREM删除，并放入执行队列LPUSH。放入成功后在任务执行通道发送消息，执行通道使用LPOP取出并执行。


```
queue := delayqueue.New()


```

