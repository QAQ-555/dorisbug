-- Active: 1753418398908@@127.0.0.1@9030
CREATE TABLE `event_log_test` (
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
