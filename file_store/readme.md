# 文件存储组件

1. 支持磁盘，OSS，FTP等
2. 支持自定义存储
3. 支持同时存储到不同的存储目标

创建文件存储

```GO

New().WithStore(LocalDisk("D:\abc"))


```

删除文件

```GO

filestore.Delete("filename")

```

文件移动

```GO

filestore.Move("old", "new")

```

文件重命名

```GO

filestore.Rename("old", "new")

```
