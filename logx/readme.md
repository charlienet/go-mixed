# 日志记录组件

日志分割及备份

日志可按照日期或大小进行分割，保留的历史日志文件数量由备份数量决定。

1. 按天拆分，每天生成新的日志文件名称。格式为file.yyyy-mm-dd.log， 其中file和log为配置的日志文件名称。
2. 按大小拆分，使用lumberjack组件对日志文件进行分割。
3. 按时间间隔拆分，日志文件按照指定的间隔拆分，

日志输出流
支持控制台和文件输出，可扩展输出组件

``` golang
logx.NewLogger(
   WithLevel("debug"),
   WithFormatter(),
   WithConsole(),
   WithRoateBySize(FileRoateSize{
      MaxSize 
      MaxAge
      MaxBackups
   }),
   WithRoateByDate("filename", MaxAge, MaxBackups),
   WithFile("filename"))
```
