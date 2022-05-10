# log-cleaning

+ 通过配置清理日志规则，配合cron,实现磁盘达到阈值自动清理
+ 支持飞书通知
+ 支持规则自定义
+ 支持全局规则设定
+ 禁止根目录和一级目录
+ 支持debug，该模式下，只显示满足条件的文件，但不执行删除或清空

```
-
  dir : /workspace/logs/webapps
  suffix : .log
  beforeTime : 1300
  mode : rm
  retain  : 0
  minSize : 0
  debug : true

# dir 目录 (必填)，目录最少2级，限制根目录和一级目录
# suffix 后缀 (必填：二选1：前缀或后缀,后缀优先),允许的后缀：".log",".json",".gz",".tgz",".tar",".zip",".rar",".mp4",".out",".bz2",".txt"
# prefix 前缀 (必填：二选1：前缀或后缀,后缀优先)，前缀必须包含 . 且文件的后缀长度大于8 ，避免勿删，例如：log.inc.java ，后缀为java，会跳过文件
# beforeTime 查找多少分钟前的文件 (选填，默认0，单位分钟)
# mode 执行方式，rm=删除，null=清空-写成空文件,waring=告警-只告警不处理 (选填，默认 null)
# retain 保留文件个数 (选填,默认0)
# minSize 最小文件大小,大于这个值才会执行 (选填,默认 0，单位 Mb)
# debug 测试模式，开启=true,关闭=false 或 不设置 （选填，默认：false）只显示要执行的结果，不实际执行
```