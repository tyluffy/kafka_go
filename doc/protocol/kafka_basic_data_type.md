## 参考文献
- https://kafka.apache.org/protocol#protocol_types
## 需知
- 大端表示字节法，也叫网络字节序。数据的高字节存放在内存的低地址中，和阅读习惯一致。
## 数据类型
### BOOLEAN
### INT8
在-2^7 ~ 2^7-1范围内的整数。
### INT16
在-2^15 ~ 2^15-1范围内的整数，使用大端表示法。
### INT32
在-2^31 ~ 2^31-1范围内的整数，使用大端表示法。
### INT64
在-2^63 ~ 2^63-1范围内的整数，使用大端表示法。
### UINT32
在0 ~ 2^32-1范围内的整数，使用大端表示法。
### VARINT
在-2^31 ~ 2^31-1范围内的整数，采用protobuf编码方式。
### VARLONG
在-2^63 ~ 2^63-1范围内的整数，采用protobuf编码方式。
### UUID
type4 UUID，使用大端表示法。
### FLOAT64
IEEE 754模式的小数，使用大端表示法。
### STRING
连续的字符。首先是以INT16表示的长度N。接下来的N个字节是字符串的UTF8编码
### COMPACT_STRING
连续的字符。首先是以UNSIGNED_VARINT表示的长度N+1。接下来的N个字节是字符串的UTF8编码
### COMPACT_NULLABLE_STRING
连续的字符。首先是以UNSIGNED_VARINT表示的长度N+1。接下来的N个字节是字符串的UTF8编码。

null使用长度0表示
### BYTES
代表字节的原始序列。首先以INT32表示字节长度N。接下来是N个字节。
### COMPACT_BYTES
代表字节的原始序列。首先以UNSIGNED_VARINT表示长度N+1。接下来是N个字节。
### NULLABLE_BYTES
代表字节的原始序列。

对于非空数据，首先以INT32表示字节长度N。接下来是N个字节。

对于空值，以-1表示长度。接下来没有数据。
### COMPACT_NULLABLE_BYTES
代表字节的原始序列。

对于非空数据，首先以UNSIGNED_VARINT表示字节长度N+1。接下来是N个字节。

对于空值，以-1表示长度。接下来没有数据。
### ARRAY
代表连续的T类型数据。T可以是基础类型，也可以是结构体。首先，以INT32表示字节长度N。然后是N个Type T。

空数组用长度-1表示。协议中当用`[T]`表示T数组。
### COMPACT_ARRAY
代表连续的T类型数据。T可以是基础类型，也可以是结构体。首先，以UNSIGNED_VARINT表示字节长度N+1。然后是N个Type T。

空数组用长度0表示。协议中当用`[T]`表示T数组。

## Messages
### 头部
```
Request Header v0 => request_api_key request_api_version correlation_id 
  request_api_key => INT16
  request_api_version => INT16
  correlation_id => INT32
```