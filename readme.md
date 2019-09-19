# Uniqueid 用于生成不重复ID的库

提供统一接口，利用不同的算法生成字符串格式的唯一ID

## 可用驱动
* simpleid 利用时间戳，自增ID和后缀生成唯一ID
* UUID 利用uuid生成唯一ID
* Snowflake 利用snowflake生成唯一ID