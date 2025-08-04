# GoFrame Template For sqlbug
该库旨在复现当数据库字段名仅有下划线区别时，利用goframe的orm模块无法正常对应字段的问题。
## 配置

```sql
CREATE TABLE `eventlog` (
  `event_time` DATETIME NULL,
  `adid` VARCHAR(200) NULL,
  `ad_id` VARCHAR(200) NULL
)
ENGINE=OLAP
UNIQUE KEY(`event_time`, `adid`)
DISTRIBUTED BY HASH(`event_time`) BUCKETS 1
PROPERTIES (
  "replication_allocation" = "tag.location.default: 1",
  "enable_unique_key_merge_on_write" = "true"
);
```
创建如上的Doris数据库，利用`gf gen dao`生成自动生成对应的`dao/do/entity`文件。

## 复现方式

```shell
curl --location --request GET 'http://127.0.0.1:8000/test' \
--form 'Adid="123"' \
--form 'AdID="abc"'
```
向test接口发送请求，查看日志输出。
能观察到参数输入一致情况下，会出现生成的sql语句不同的结果